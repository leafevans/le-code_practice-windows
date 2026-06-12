using Microsoft.Xna.Framework;
using Microsoft.Xna.Framework.Graphics;

namespace TankBattle;

public class Bullet {
    public Vector2 Position;
    public Direction Direction;
    public float Speed;
    public bool IsPlayerBullet;
    public bool IsActive = true;

    private readonly Texture2D _texture;
    private readonly int _size = GameConstants.BulletSize;

    public Bullet(GraphicsDevice graphicsDevice, Vector2 position, Direction direction,
                  float speed, bool isPlayerBullet) {
        Position = position;
        Direction = direction;
        Speed = speed;
        IsPlayerBullet = isPlayerBullet;

        _texture = new Texture2D(graphicsDevice, _size, _size);
        var data = new Color[_size * _size];
        var bulletColor = isPlayerBullet ? Color.White : Color.Red;
        for (int i = 0; i < data.Length; i++) {
            data[i] = bulletColor;
        }
        _texture.SetData(data);
    }

    public Rectangle Bounds => new(
        (int)Position.X - _size / 2,
        (int)Position.Y - _size / 2,
        _size, _size);

    public void Update(GameTime gameTime, GameMap map) {
        if (!IsActive) {
            return;
        }

        float dt = (float)gameTime.ElapsedGameTime.TotalSeconds;
        var velocity = Direction switch {
            Direction.Up => new Vector2(0, -Speed * dt),
            Direction.Down => new Vector2(0, Speed * dt),
            Direction.Left => new Vector2(-Speed * dt, 0),
            Direction.Right => new Vector2(Speed * dt, 0),
            _ => Vector2.Zero
        };

        Position += velocity;

        if (Position.X < 0 || Position.X > GameConstants.GameAreaWidth ||
            Position.Y < 0 || Position.Y > GameConstants.GameAreaHeight) {
            IsActive = false;
            return;
        }

        int tileX = (int)Position.X / GameConstants.TileSize;
        int tileY = (int)Position.Y / GameConstants.TileSize;
        if (map.HitTile(tileX, tileY)) {
            IsActive = false;
        }
    }

    public void Draw(SpriteBatch spriteBatch) {
        if (!IsActive) {
            return;
        }
        spriteBatch.Draw(_texture, Bounds, Color.White);
    }
}
