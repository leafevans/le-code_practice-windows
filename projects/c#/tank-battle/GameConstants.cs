using System;

namespace TankBattle;

public static class GameConstants {
    public const int MapWidth = 26;
    public const int MapHeight = 26;
    public const int TileSize = 16;

    public const int GameAreaWidth = MapWidth * TileSize;
    public const int GameAreaHeight = MapHeight * TileSize;

    public const int InfoPanelWidth = 80;
    public const int WindowWidth = GameAreaWidth + InfoPanelWidth;
    public const int WindowHeight = GameAreaHeight;

    public const int TankSize = TileSize * 2;
    public const float PlayerSpeed = 80f;
    public const float EnemyBasicSpeed = 40f;
    public const float EnemyFastSpeed = 80f;
    public const float EnemyPowerSpeed = 40f;
    public const float EnemyArmorSpeed = 30f;

    public const float BulletSpeed = 200f;
    public const float PlayerBulletSpeed = 200f;
    public const float PowerBulletSpeed = 300f;
    public const int BulletSize = 4;

    public const int PlayerLives = 3;
    public const int MaxEnemiesOnScreen = 4;
    public const float InvincibleDuration = 3f;
    public const float EnemyShootCooldown = 1.5f;

    public const float ExplosionDuration = 0.5f;

    public const float PowerUpLifetime = 15f;
    public const float FreezeDuration = 10f;
    public const float StarSpeedMultiplier = 1.5f;
    public const float StarCooldownMultiplier = 0.6f;
    public const int MaxStarLevel = 3;
    public const float StarDuration = 20f;
    public const float HelmetPowerUpDuration = 10f;
    public const float ShovelFortifyDuration = 15f;

    public const float EnemySmartShootCooldown = 1.0f;
    public const float BulletDetectionRadius = 64f;

    public const int TotalLevels = 5;

    public static int GetEnemiesForLevel(int level) => level switch {
        1 => 20,
        2 => 22,
        3 => 24,
        4 => 26,
        5 => 28,
        _ => 20 + level * 2
    };

    public static float GetSpawnIntervalForLevel(int level) => level switch {
        1 => 3.0f,
        2 => 2.5f,
        3 => 2.0f,
        4 => 1.8f,
        5 => 1.5f,
        _ => MathF.Max(1.0f, 3.0f - level * 0.3f)
    };
}
