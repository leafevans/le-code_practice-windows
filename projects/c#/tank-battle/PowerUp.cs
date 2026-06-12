using System;
using Microsoft.Xna.Framework;
using Microsoft.Xna.Framework.Graphics;

namespace TankBattle;

public class PowerUp {
    public PowerUpType Type;
    public Vector2 Position;
    public bool IsActive = true;

    private Texture2D _texture;
    private readonly int _size = GameConstants.TileSize;
    private float _blinkTimer = 0f;
    private float _lifetime;
    private bool _blinkVisible = true;

    public PowerUp(GraphicsDevice graphicsDevice, PowerUpType type, Vector2 position) {
        Type = type;
        Position = position;
        _lifetime = GameConstants.PowerUpLifetime;
        CreateTexture(graphicsDevice);
    }

    private void CreateTexture(GraphicsDevice graphicsDevice) {
        int size = _size;
        _texture = new Texture2D(graphicsDevice, size, size);
        var data = new Color[size * size];
        int cx = size / 2;
        int cy = size / 2;

        switch (Type) {
            case PowerUpType.Star:
                for (int y = 0; y < size; y++) {
                    for (int x = 0; x < size; x++) {
                        float dx = x - cx;
                        float dy = y - cy;
                        float dist = MathF.Sqrt(dx * dx + dy * dy);
                        float angle = MathF.Atan2(dy, dx);
                        float star = MathF.Cos(angle * 5f) * 0.4f + 0.6f;
                        data[y * size + x] = dist / (size / 2f) < star
                            ? Color.Gold
                            : Color.Transparent;
                    }
                }
                break;

            case PowerUpType.Helmet:
                for (int y = 0; y < size; y++) {
                    for (int x = 0; x < size; x++) {
                        float dx = x - cx;
                        float dy = y - cy;
                        float dist = MathF.Sqrt(dx * dx + dy * dy);
                        if (dist < size / 2f - 1) {
                            data[y * size + x] = dist < size / 4f
                                ? Color.White
                                : new Color(180, 180, 180);
                        } else {
                            data[y * size + x] = Color.Transparent;
                        }
                    }
                }
                break;

            case PowerUpType.Bomb:
                for (int y = 0; y < size; y++) {
                    for (int x = 0; x < size; x++) {
                        float dx = x - cx;
                        float dy = y - cy;
                        float dist = MathF.Sqrt(dx * dx + dy * dy);
                        if (dist < size / 2f - 1) {
                            data[y * size + x] = dist < size / 4f
                                ? Color.Yellow
                                : Color.Red;
                        } else {
                            data[y * size + x] = Color.Transparent;
                        }
                    }
                }
                break;

            case PowerUpType.Shovel:
                for (int y = 0; y < size; y++) {
                    for (int x = 0; x < size; x++) {
                        bool isHandle = x >= cx - 1 && x <= cx && y >= 2 && y <= size - 4;
                        bool isHead = x >= cx - 3 && x <= cx + 2 && y >= size - 6 && y <= size - 3;
                        if (isHandle) {
                            data[y * size + x] = new Color(139, 90, 43);
                        } else if (isHead) {
                            data[y * size + x] = new Color(160, 160, 160);
                        } else {
                            data[y * size + x] = Color.Transparent;
                        }
                    }
                }
                break;

            case PowerUpType.Tank:
                for (int y = 0; y < size; y++) {
                    for (int x = 0; x < size; x++) {
                        bool isBody = x >= cx - 3 && x <= cx + 2 && y >= 4 && y <= size - 4;
                        bool isTurret = x >= cx - 1 && x <= cx && y >= 2 && y <= size - 4;
                        bool isBarrel = x >= cx - 1 && x <= cx && y >= 0 && y <= 4;
                        if (isBarrel) {
                            data[y * size + x] = Color.White;
                        } else if (isTurret) {
                            data[y * size + x] = new Color(0, 180, 0);
                        } else if (isBody) {
                            data[y * size + x] = new Color(0, 140, 0);
                        } else {
                            data[y * size + x] = Color.Transparent;
                        }
                    }
                }
                break;

            case PowerUpType.Timer:
                for (int y = 0; y < size; y++) {
                    for (int x = 0; x < size; x++) {
                        float dx = x - cx;
                        float dy = y - cy;
                        float dist = MathF.Sqrt(dx * dx + dy * dy);
                        if (dist < size / 2f - 1) {
                            if (y == cy && x >= cx) {
                                data[y * size + x] = Color.White;
                            } else if (x == cx && y <= cy) {
                                data[y * size + x] = Color.White;
                            } else {
                                data[y * size + x] = new Color(50, 100, 200);
                            }
                        } else {
                            data[y * size + x] = Color.Transparent;
                        }
                    }
                }
                break;
        }

        _texture.SetData(data);
    }

    public Rectangle Bounds => new(
        (int)Position.X,
        (int)Position.Y,
        _size, _size);

    public void Update(GameTime gameTime) {
        float dt = (float)gameTime.ElapsedGameTime.TotalSeconds;

        _blinkTimer += dt;
        if (_blinkTimer >= 0.2f) {
            _blinkTimer = 0;
            _blinkVisible = !_blinkVisible;
        }

        _lifetime -= dt;
        if (_lifetime <= 0) {
            IsActive = false;
        }
    }

    public void Draw(SpriteBatch spriteBatch) {
        if (!IsActive || !_blinkVisible) {
            return;
        }
        spriteBatch.Draw(_texture, Bounds, Color.White);
    }
}
