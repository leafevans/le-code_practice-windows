namespace TankBattle;

public enum Direction {
    Up,
    Down,
    Left,
    Right
}

public enum TileType {
    Empty = 0,
    Brick = 1,
    Steel = 2,
    Water = 3,
    Trees = 4,
    Base = 5,
    BaseDestroyed = 6
}

public enum EnemyType {
    Basic,
    Fast,
    Power,
    Armor
}

public enum GameState {
    Title,
    Playing,
    GameOver,
    Victory,
    LevelComplete,
    Paused,
    Transition
}

public enum PowerUpType {
    Star,
    Helmet,
    Bomb,
    Shovel,
    Tank,
    Timer
}
