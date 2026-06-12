using System;
using Microsoft.Xna.Framework;
using Microsoft.Xna.Framework.Graphics;

namespace TankBattle;

public class Explosion {
    public Vector2 Position;
    public float Timer;
    public float Duration;
    public bool IsActive = true;
    public bool IsBig;

    private readonly Texture2D _texture;

    public Explosion(GraphicsDevice graphicsDevice, Vector2 position, bool isBig = false) {
        Position = position;
        Duration = GameConstants.ExplosionDuration;
        Timer = Duration;
        IsBig = isBig;

        int size = isBig ? GameConstants.TankSize : GameConstants.TileSize;
        _texture = new Texture2D(graphicsDevice, size, size);
        var data = new Color[size * size];

        int cx = size / 2;
        int cy = size / 2;
        for (int y = 0; y < size; y++) {
            for (int x = 0; x < size; x++) {
                float dist = MathF.Sqrt((x - cx) * (x - cx) + (y - cy) * (y - cy));
                float ratio = dist / (size / 2f);
                if (ratio < 0.5f) {
                    data[y * size + x] = Color.White;
                } else if (ratio < 0.7f) {
                    data[y * size + x] = Color.Yellow;
                } else if (ratio < 0.9f) {
                    data[y * size + x] = new Color(255, 140, 0);
                } else if (ratio < 1.0f) {
                    data[y * size + x] = Color.Red;
                } else {
                    data[y * size + x] = Color.Transparent;
                }
            }
        }
        _texture.SetData(data);
    }

    public void Update(GameTime gameTime) {
        Timer -= (float)gameTime.ElapsedGameTime.TotalSeconds;
        if (Timer <= 0) {
            IsActive = false;
        }
    }

    public void Draw(SpriteBatch spriteBatch) {
        if (!IsActive) {
            return;
        }

        float progress = 1f - (Timer / Duration);
        float scale = IsBig ? 0.5f + progress * 1.5f : 0.5f + progress;
        float alpha = 1f - progress;

        int size = IsBig ? GameConstants.TankSize : GameConstants.TileSize;
        var origin = new Vector2(size / 2f, size / 2f);
        spriteBatch.Draw(_texture, Position, null, Color.White * alpha, 0f, origin, scale,
            SpriteEffects.None, 0f);
    }
}
