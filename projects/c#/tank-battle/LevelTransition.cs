using Microsoft.Xna.Framework;
using Microsoft.Xna.Framework.Graphics;

namespace TankBattle;

public class LevelTransition {
    public bool IsActive = true;
    public int StageNumber;

    private float _timer;
    private const float WipeInDuration = 0.8f;
    private const float ShowTextDuration = 0.9f;
    private const float WipeOutDuration = 0.8f;
    private const float TotalDuration = WipeInDuration + ShowTextDuration + WipeOutDuration;

    private enum Phase { WipeIn, ShowText, WipeOut }
    private Phase _phase = Phase.WipeIn;
    private readonly Texture2D _overlayTexture;

    public LevelTransition(GraphicsDevice graphicsDevice, int stageNumber) {
        StageNumber = stageNumber;
        _timer = 0f;
        _overlayTexture = new Texture2D(graphicsDevice, 1, 1);
        _overlayTexture.SetData([new Color(80, 80, 80)]);
    }

    public void Update(GameTime gameTime) {
        float dt = (float)gameTime.ElapsedGameTime.TotalSeconds;
        _timer += dt;

        if (_timer >= TotalDuration) {
            IsActive = false;
            return;
        }

        if (_timer < WipeInDuration) {
            _phase = Phase.WipeIn;
        } else if (_timer < WipeInDuration + ShowTextDuration) {
            _phase = Phase.ShowText;
        } else {
            _phase = Phase.WipeOut;
        }
    }

    public void Draw(SpriteBatch spriteBatch, SpriteFont font, int gameAreaWidth, int gameAreaHeight) {
        float centerY = gameAreaHeight / 2f;
        float barHeight;
        var overlay = new Color(80, 80, 80);

        switch (_phase) {
            case Phase.WipeIn:
                barHeight = centerY * (_timer / WipeInDuration);
                spriteBatch.Draw(_overlayTexture, new Rectangle(0, 0, gameAreaWidth, (int)barHeight), overlay);
                spriteBatch.Draw(_overlayTexture,
                    new Rectangle(0, gameAreaHeight - (int)barHeight, gameAreaWidth, (int)barHeight), overlay);
                break;

            case Phase.ShowText:
                spriteBatch.Draw(_overlayTexture, new Rectangle(0, 0, gameAreaWidth, (int)centerY), overlay);
                spriteBatch.Draw(_overlayTexture,
                    new Rectangle(0, (int)centerY, gameAreaWidth, (int)centerY), overlay);
                string text = $"STAGE  {StageNumber}";
                var textSize = font.MeasureString(text);
                float textX = gameAreaWidth / 2f - textSize.X / 2;
                float textY = gameAreaHeight / 2f - textSize.Y / 2;
                spriteBatch.DrawString(font, text, new Vector2(textX, textY), Color.White);
                break;

            case Phase.WipeOut:
                float wipeOutTime = _timer - WipeInDuration - ShowTextDuration;
                barHeight = centerY * (1f - wipeOutTime / WipeOutDuration);
                spriteBatch.Draw(_overlayTexture, new Rectangle(0, 0, gameAreaWidth, (int)barHeight), overlay);
                spriteBatch.Draw(_overlayTexture,
                    new Rectangle(0, gameAreaHeight - (int)barHeight, gameAreaWidth, (int)barHeight), overlay);
                break;
        }
    }

    public void Dispose() => _overlayTexture?.Dispose();
}
