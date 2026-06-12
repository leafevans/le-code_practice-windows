using System.Collections.Generic;
using Microsoft.Xna.Framework;
using Microsoft.Xna.Framework.Graphics;

namespace TankBattle;

public class Tank {
    public Vector2 Position;
    public Direction Facing = Direction.Up;
    public float Speed;
    public int Health = 1;
    public bool IsActive = true;
    public bool IsPlayer;
    public Color TankColor;
    public float ShootCooldown = 0f;
    public float ShootCooldownMax = 0.5f;

    protected Texture2D _texture;
    protected Texture2D _trackTexture;
    protected int _size = GameConstants.TankSize;
    protected float _trackAnimTimer = 0f;

    public Tank(GraphicsDevice graphicsDevice, Color color, bool isPlayer) {
        TankColor = color;
        IsPlayer = isPlayer;
        CreateTextures(graphicsDevice);
    }

    private void CreateTextures(GraphicsDevice graphicsDevice) {
        int size = _size;
        _texture = new Texture2D(graphicsDevice, size, size);
        var data = new Color[size * size];

        for (int y = 0; y < size; y++) {
            for (int x = 0; x < size; x++) {
                int idx = y * size + x;
                bool isLeftTrack = x >= 2 && x <= size / 2 - 3;
                bool isRightTrack = x >= size / 2 + 2 && x <= size - 3;
                bool isTrack = (isLeftTrack || isRightTrack) && y >= 2 && y <= size - 3;
                bool isBody = x >= size / 2 - 2 && x <= size / 2 + 1 && y >= 4 && y <= size - 5;
                bool isTurret = x >= size / 2 - 1 && x <= size / 2 && y >= 2 && y <= size - 5;
                bool isBarrel = x >= size / 2 - 1 && x <= size / 2 && y >= 0 && y <= 4;

                if (isBarrel) {
                    data[idx] = Color.White;
                } else if (isTurret) {
                    data[idx] = TankColor * 0.8f;
                } else if (isBody) {
                    data[idx] = TankColor;
                } else if (isTrack) {
                    data[idx] = TankColor * 0.6f;
                } else {
                    data[idx] = Color.Transparent;
                }
            }
        }
        _texture.SetData(data);
    }

    public Rectangle Bounds => new(
        (int)Position.X,
        (int)Position.Y,
        _size, _size);

    public Vector2 Center => new(
        Position.X + _size / 2,
        Position.Y + _size / 2);

    public Vector2 GetMuzzlePosition() => Facing switch {
        Direction.Up => new Vector2(Center.X, Position.Y),
        Direction.Down => new Vector2(Center.X, Position.Y + _size),
        Direction.Left => new Vector2(Position.X, Center.Y),
        Direction.Right => new Vector2(Position.X + _size, Center.Y),
        _ => Center
    };

    public bool TryMove(Direction direction, float dt, GameMap map, List<Rectangle> otherTankBounds) {
        Facing = direction;
        var newPos = Position;
        float moveAmount = Speed * dt;

        switch (direction) {
            case Direction.Up:
                newPos.Y -= moveAmount;
                break;
            case Direction.Down:
                newPos.Y += moveAmount;
                break;
            case Direction.Left:
                newPos.X -= moveAmount;
                break;
            case Direction.Right:
                newPos.X += moveAmount;
                break;
        }

        newPos.X = MathHelper.Clamp(newPos.X, 0, GameConstants.GameAreaWidth - _size);
        newPos.Y = MathHelper.Clamp(newPos.Y, 0, GameConstants.GameAreaHeight - _size);

        if (CheckTileCollision(newPos, map)) {
            return false;
        }

        Rectangle newBounds = new((int)newPos.X, (int)newPos.Y, _size, _size);
        if (otherTankBounds != null) {
            foreach (var bounds in otherTankBounds) {
                if (newBounds.Intersects(bounds)) {
                    return false;
                }
            }
        }

        Position = newPos;
        _trackAnimTimer += dt;
        return true;
    }

    private bool CheckTileCollision(Vector2 newPos, GameMap map) {
        int ts = GameConstants.TileSize;
        int leftTile = (int)newPos.X / ts;
        int rightTile = (int)(newPos.X + _size - 1) / ts;
        int topTile = (int)newPos.Y / ts;
        int bottomTile = (int)(newPos.Y + _size - 1) / ts;

        for (int tx = leftTile; tx <= rightTile; tx++) {
            for (int ty = topTile; ty <= bottomTile; ty++) {
                if (!map.IsPassable(tx, ty)) {
                    return true;
                }
            }
        }
        return false;
    }

    public void TakeDamage() {
        Health--;
        if (Health <= 0) {
            IsActive = false;
        }
    }

    public void Draw(SpriteBatch spriteBatch) {
        if (!IsActive) {
            return;
        }

        float rotation = Facing switch {
            Direction.Up => 0f,
            Direction.Right => MathHelper.PiOver2,
            Direction.Down => MathHelper.Pi,
            Direction.Left => -MathHelper.PiOver2,
            _ => 0f
        };

        Vector2 origin = new(_size / 2f, _size / 2f);
        Vector2 drawPos = new(Position.X + _size / 2, Position.Y + _size / 2);
        spriteBatch.Draw(_texture, drawPos, null, Color.White, rotation, origin, 1f,
            SpriteEffects.None, 0f);
    }
}
