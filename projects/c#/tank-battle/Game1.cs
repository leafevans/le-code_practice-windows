using System;
using System.Collections.Generic;
using System.IO;
using Microsoft.Xna.Framework;
using Microsoft.Xna.Framework.Graphics;
using Microsoft.Xna.Framework.Input;

namespace TankBattle;

public class Game1 : Game {
    private readonly GraphicsDeviceManager _graphics;
    private SpriteBatch _spriteBatch;
    private SpriteFont _font;

    // 游戏状态
    private GameState _state = GameState.Title;
    private GameState _previousState;
    private int _score = 0;
    private int _highScore = 0;
    private int _lives = GameConstants.PlayerLives;
    private int _level = 1;
    private int _enemiesRemaining;
    private int _enemiesOnScreen;
    private float _enemySpawnTimer;
    private float _levelCompleteTimer;
    private float _gameOverTimer;
    private float _respawnTimer = 0f;
    private float _respawnTimerP2 = 0f;

    // 游戏对象
    private GameMap _map;
    private Tank _player;
    private Tank _player2;
    private readonly List<EnemyTank> _enemies = [];
    private readonly List<Bullet> _bullets = [];
    private readonly List<Explosion> _explosions = [];
    private readonly List<PowerUp> _powerUps = [];

    // 玩家无敌状态
    private float _invincibleTimer = 0f;
    private float _invincibleTimerP2 = 0f;

    // 道具效果计时器
    private float _freezeTimer = 0f;
    private float _shovelTimer = 0f;

    // Star 道具状态
    private int _starLevel = 0;
    private float _starTimer = 0f;
    private int _starLevelP2 = 0;
    private float _starTimerP2 = 0f;

    // 关卡过渡动画
    private LevelTransition _transition;

    // 音效
    private SoundManager _sound;

    // 双人模式
    private bool _isTwoPlayer = false;

    // 高分文件路径
    private string _highScorePath;

    // 敌方出生点
    private readonly Vector2[] _enemySpawnPoints = [
        new Vector2(0, 0),
        new Vector2(GameConstants.GameAreaWidth - GameConstants.TankSize, 0),
        new Vector2(GameConstants.GameAreaWidth / 2 - GameConstants.TankSize / 2, 0)
    ];
    private int _nextSpawnPoint = 0;

    // 玩家出生点
    private readonly Vector2 _playerSpawnPos = new(
        8 * GameConstants.TileSize,
        24 * GameConstants.TileSize
    );

    private readonly Vector2 _player2SpawnPos = new(
        16 * GameConstants.TileSize,
        24 * GameConstants.TileSize
    );

    // 按键状态
    private KeyboardState _prevKeyState;

    // 闪烁效果
    private float _blinkTimer = 0f;
    private bool _blinkVisible = true;

    // 侧边栏
    private Texture2D _pixelTexture;
    private Texture2D _enemyIconTexture;
    private Texture2D _lifeIconTexture;

    public Game1() {
        _graphics = new GraphicsDeviceManager(this);
        Content.RootDirectory = "Content";
        IsMouseVisible = true;

        _graphics.PreferredBackBufferWidth = GameConstants.WindowWidth;
        _graphics.PreferredBackBufferHeight = GameConstants.WindowHeight;
    }

    protected override void Initialize() => base.Initialize();

    protected override void LoadContent() {
        _spriteBatch = new SpriteBatch(GraphicsDevice);

        _pixelTexture = new Texture2D(GraphicsDevice, 1, 1);
        _pixelTexture.SetData(new[] { Color.White });

        CreateSidePanelTextures();

        _font = Content.Load<SpriteFont>("DefaultFont");

        _sound = new SoundManager();
        _sound.Initialize();

        _highScorePath = Path.Combine(
            Environment.GetFolderPath(Environment.SpecialFolder.LocalApplicationData),
            "TankBattle", "highscore.txt");
        LoadHighScore();

        StartNewGame();
    }

    private void CreateSidePanelTextures() {
        _enemyIconTexture = new Texture2D(GraphicsDevice, 8, 8);
        Color[] enemyIcon = new Color[64];
        for (int i = 0; i < 64; i++) {
            enemyIcon[i] = Color.Red;
        }
        _enemyIconTexture.SetData(enemyIcon);

        _lifeIconTexture = new Texture2D(GraphicsDevice, 8, 8);
        Color[] lifeIcon = new Color[64];
        for (int i = 0; i < 64; i++) {
            lifeIcon[i] = Color.Green;
        }
        _lifeIconTexture.SetData(lifeIcon);
    }

    private void LoadHighScore() {
        try {
            if (File.Exists(_highScorePath)) {
                string text = File.ReadAllText(_highScorePath).Trim();
                int.TryParse(text, out _highScore);
            }
        } catch { }
    }

    private void SaveHighScore() {
        try {
            string dir = Path.GetDirectoryName(_highScorePath)!;
            if (!Directory.Exists(dir)) {
                Directory.CreateDirectory(dir);
            }
            File.WriteAllText(_highScorePath, _highScore.ToString());
        } catch { }
    }

    private void StartNewGame() {
        _score = 0;
        _lives = GameConstants.PlayerLives;
        _level = 1;
        _state = GameState.Playing;
        _starLevel = 0;
        _starTimer = 0f;
        _starLevelP2 = 0;
        _starTimerP2 = 0f;
        StartLevel();
    }

    private void StartLevel() {
        _map = new GameMap(GraphicsDevice);
        _map.LoadLevel(_level);
        _enemies.Clear();
        _bullets.Clear();
        _explosions.Clear();
        _powerUps.Clear();
        _enemiesRemaining = GameConstants.GetEnemiesForLevel(_level);
        _enemiesOnScreen = 0;
        _enemySpawnTimer = 1f;
        _nextSpawnPoint = 0;
        _freezeTimer = 0f;
        _shovelTimer = 0f;

        SpawnPlayer(ref _player, _playerSpawnPos, ref _invincibleTimer, _starLevel);
        if (_isTwoPlayer) {
            SpawnPlayer(ref _player2, _player2SpawnPos, ref _invincibleTimerP2, _starLevelP2);
        }
    }

    private void SpawnPlayer(ref Tank player, Vector2 spawnPos, ref float invincibleTimer, int starLevel) {
        player = new Tank(GraphicsDevice, Color.Yellow, true) {
            Position = spawnPos,
            Speed = GameConstants.PlayerSpeed,
            ShootCooldownMax = 0.3f
        };
        // 保留 Star 升级效果
        for (int i = 0; i < starLevel; i++) {
            player.Speed *= GameConstants.StarSpeedMultiplier;
            player.ShootCooldownMax *= GameConstants.StarCooldownMultiplier;
        }
        invincibleTimer = GameConstants.InvincibleDuration;
    }

    private void SpawnEnemy() {
        if (_enemiesRemaining <= 0 || _enemiesOnScreen >= GameConstants.MaxEnemiesOnScreen) {
            return;
        }

        var spawnPos = _enemySpawnPoints[_nextSpawnPoint % _enemySpawnPoints.Length];
        _nextSpawnPoint++;

        var type = _level switch {
            1 => _enemiesRemaining <= 3 ? EnemyType.Fast : EnemyType.Basic,
            2 => (_enemiesRemaining % 3) switch {
                0 => EnemyType.Fast,
                1 => EnemyType.Power,
                _ => EnemyType.Basic
            },
            3 => (_enemiesRemaining % 4) switch {
                0 => EnemyType.Armor,
                1 => EnemyType.Fast,
                2 => EnemyType.Power,
                _ => EnemyType.Basic
            },
            4 => (_enemiesRemaining % 5) switch {
                0 => EnemyType.Armor,
                1 => EnemyType.Fast,
                2 => EnemyType.Power,
                3 => EnemyType.Fast,
                _ => EnemyType.Basic
            },
            _ => (_enemiesRemaining % 4) switch {
                0 => EnemyType.Armor,
                1 => EnemyType.Fast,
                2 => EnemyType.Power,
                _ => EnemyType.Basic
            }
        };

        var enemy = new EnemyTank(GraphicsDevice, type, spawnPos);

        // 标记特殊敌方（第4个和第11个生成的敌方掉落道具）
        int totalEnemies = GameConstants.GetEnemiesForLevel(_level);
        int spawnIndex = totalEnemies - _enemiesRemaining;
        if (spawnIndex == 3 || spawnIndex == 10) {
            enemy.IsSpecial = true;
        }

        // 检查出生点是否被占用
        bool blocked = false;
        foreach (var e in _enemies) {
            if (e.IsActive && e.Bounds.Intersects(enemy.Bounds)) {
                blocked = true;
                break;
            }
        }
        if (_player != null && _player.IsActive && _player.Bounds.Intersects(enemy.Bounds)) {
            blocked = true;
        }
        if (_player2 != null && _player2.IsActive && _player2.Bounds.Intersects(enemy.Bounds)) {
            blocked = true;
        }

        if (!blocked) {
            _enemies.Add(enemy);
            _enemiesOnScreen++;
            _enemiesRemaining--;
        }
    }

    protected override void Update(GameTime gameTime) {
        var keyState = Keyboard.GetState();

        switch (_state) {
            case GameState.Title:
                UpdateTitle(keyState);
                break;
            case GameState.Playing:
                UpdatePlaying(gameTime, keyState);
                break;
            case GameState.GameOver:
                UpdateGameOver(gameTime, keyState);
                break;
            case GameState.Victory:
                UpdateGameOver(gameTime, keyState);
                break;
            case GameState.LevelComplete:
                UpdateLevelComplete(gameTime);
                break;
            case GameState.Paused:
                UpdatePaused(keyState);
                break;
            case GameState.Transition:
                UpdateTransition(gameTime);
                break;
        }

        _prevKeyState = keyState;
        base.Update(gameTime);
    }

    private void UpdateTitle(KeyboardState keyState) {
        // 单人/双人模式选择
        if (keyState.IsKeyDown(Keys.D1) && !_prevKeyState.IsKeyDown(Keys.D1)) {
            _isTwoPlayer = false;
            StartNewGame();
        } else if (keyState.IsKeyDown(Keys.D2) && !_prevKeyState.IsKeyDown(Keys.D2)) {
            _isTwoPlayer = true;
            StartNewGame();
        } else if (keyState.IsKeyDown(Keys.Enter) && !_prevKeyState.IsKeyDown(Keys.Enter)) {
            StartNewGame();
        }
    }

    private void UpdateGameOver(GameTime gameTime, KeyboardState keyState) {
        _gameOverTimer -= (float)gameTime.ElapsedGameTime.TotalSeconds;
        if (_gameOverTimer <= 0 && keyState.IsKeyDown(Keys.Enter) && !_prevKeyState.IsKeyDown(Keys.Enter)) {
            _state = GameState.Title;
        }
    }

    private void UpdateLevelComplete(GameTime gameTime) {
        _levelCompleteTimer -= (float)gameTime.ElapsedGameTime.TotalSeconds;
        if (_levelCompleteTimer <= 0) {
            if (_level >= GameConstants.TotalLevels) {
                if (_score > _highScore) {
                    _highScore = _score;
                    SaveHighScore();
                }
                _state = GameState.Victory;
                _gameOverTimer = 3f;
            } else {
                _level++;
                _transition = new LevelTransition(GraphicsDevice, _level);
                _state = GameState.Transition;
                StartLevel();
            }
        }
    }

    private void UpdatePaused(KeyboardState keyState) {
        if (keyState.IsKeyDown(Keys.P) && !_prevKeyState.IsKeyDown(Keys.P)) {
            _state = _previousState;
        }
    }

    private void UpdateTransition(GameTime gameTime) {
        _transition.Update(gameTime);
        if (!_transition.IsActive) {
            _state = GameState.Playing;
            _transition.Dispose();
            _transition = null;
        }
    }

    private void UpdatePlaying(GameTime gameTime, KeyboardState keyState) {
        // P 键暂停
        if (keyState.IsKeyDown(Keys.P) && !_prevKeyState.IsKeyDown(Keys.P)) {
            _previousState = _state;
            _state = GameState.Paused;
            return;
        }

        float dt = (float)gameTime.ElapsedGameTime.TotalSeconds;

        // 闪烁效果
        _blinkTimer += dt;
        if (_blinkTimer >= 0.15f) {
            _blinkTimer = 0;
            _blinkVisible = !_blinkVisible;
        }

        // 无敌时间
        if (_invincibleTimer > 0) {
            _invincibleTimer -= dt;
        }
        if (_invincibleTimerP2 > 0) {
            _invincibleTimerP2 -= dt;
        }

        // Star 道具计时
        if (_starTimer > 0) {
            _starTimer -= dt;
            if (_starTimer <= 0) {
                ResetStarEffect(ref _player, ref _starLevel);
            }
        }
        if (_starTimerP2 > 0) {
            _starTimerP2 -= dt;
            if (_starTimerP2 <= 0) {
                ResetStarEffect(ref _player2, ref _starLevelP2);
            }
        }

        // 道具效果计时
        if (_freezeTimer > 0) {
            _freezeTimer -= dt;
        }
        if (_shovelTimer > 0) {
            _shovelTimer -= dt;
            if (_shovelTimer <= 0) {
                _map.UnfortifyBase();
            }
        }

        // 玩家重生计时
        if (_respawnTimer > 0) {
            _respawnTimer -= dt;
            if (_respawnTimer <= 0 && _lives > 0) {
                SpawnPlayer(ref _player, _playerSpawnPos, ref _invincibleTimer, _starLevel);
            }
        }
        if (_respawnTimerP2 > 0) {
            _respawnTimerP2 -= dt;
            if (_respawnTimerP2 <= 0 && _lives > 0 && _isTwoPlayer) {
                SpawnPlayer(ref _player2, _player2SpawnPos, ref _invincibleTimerP2, _starLevelP2);
            }
        }

        // 生成敌方
        _enemySpawnTimer -= dt;
        if (_enemySpawnTimer <= 0) {
            SpawnEnemy();
            _enemySpawnTimer = GameConstants.GetSpawnIntervalForLevel(_level);
        }

        // 收集所有坦克碰撞体
        var allTankBounds = CollectTankBounds();

        // 更新玩家1
        UpdatePlayer(gameTime, keyState, _player, allTankBounds, true);

        // 更新玩家2
        if (_isTwoPlayer) {
            UpdatePlayer2(gameTime, keyState, allTankBounds);
        }

        // 更新敌方（冻结时跳过）
        if (_freezeTimer <= 0) {
            UpdateEnemies(gameTime);
        }

        // 更新子弹
        UpdateBullets(gameTime);

        // 更新爆炸
        UpdateExplosions(gameTime);

        // 更新道具
        UpdatePowerUps(gameTime);

        // 检查碰撞
        CheckCollisions();

        // 检查游戏状态
        CheckGameState();
    }

    /// <summary>
    /// 收集所有活跃坦克的碰撞体（用于移动碰撞检测）
    /// </summary>
    private List<Rectangle> CollectTankBounds() {
        var bounds = new List<Rectangle>();
        if (_player != null && _player.IsActive) {
            bounds.Add(_player.Bounds);
        }
        if (_player2 != null && _player2.IsActive) {
            bounds.Add(_player2.Bounds);
        }
        foreach (var enemy in _enemies) {
            if (enemy.IsActive) {
                bounds.Add(enemy.Bounds);
            }
        }
        return bounds;
    }

    private void UpdatePlayer(GameTime gameTime, KeyboardState keyState, Tank player,
                              List<Rectangle> allTankBounds, bool isP1) {
        if (player == null || !player.IsActive) {
            return;
        }

        float dt = (float)gameTime.ElapsedGameTime.TotalSeconds;

        // 射击冷却
        if (player.ShootCooldown > 0) {
            player.ShootCooldown -= dt;
        }

        // 移动（过滤掉自己）
        var moveBounds = new List<Rectangle>();
        foreach (var b in allTankBounds) {
            if (b != player.Bounds) {
                moveBounds.Add(b);
            }
        }

        if (isP1) {
            if (keyState.IsKeyDown(Keys.W) || keyState.IsKeyDown(Keys.Up)) {
                player.TryMove(Direction.Up, dt, _map, moveBounds);
            } else if (keyState.IsKeyDown(Keys.S) || keyState.IsKeyDown(Keys.Down)) {
                player.TryMove(Direction.Down, dt, _map, moveBounds);
            } else if (keyState.IsKeyDown(Keys.A) || keyState.IsKeyDown(Keys.Left)) {
                player.TryMove(Direction.Left, dt, _map, moveBounds);
            } else if (keyState.IsKeyDown(Keys.D) || keyState.IsKeyDown(Keys.Right)) {
                player.TryMove(Direction.Right, dt, _map, moveBounds);
            }

            // 射击
            if (keyState.IsKeyDown(Keys.Space) && !_prevKeyState.IsKeyDown(Keys.Space)) {
                if (player.ShootCooldown <= 0) {
                    FireBullet(player, true);
                    player.ShootCooldown = player.ShootCooldownMax;
                    _sound.PlayShoot(true);
                }
            }
        }
    }

    private void UpdatePlayer2(GameTime gameTime, KeyboardState keyState, List<Rectangle> allTankBounds) {
        if (_player2 == null || !_player2.IsActive) {
            return;
        }

        float dt = (float)gameTime.ElapsedGameTime.TotalSeconds;

        if (_player2.ShootCooldown > 0) {
            _player2.ShootCooldown -= dt;
        }

        var moveBounds = new List<Rectangle>();
        foreach (var b in allTankBounds) {
            if (b != _player2.Bounds) {
                moveBounds.Add(b);
            }
        }

        if (keyState.IsKeyDown(Keys.Up)) {
            _player2.TryMove(Direction.Up, dt, _map, moveBounds);
        } else if (keyState.IsKeyDown(Keys.Down)) {
            _player2.TryMove(Direction.Down, dt, _map, moveBounds);
        } else if (keyState.IsKeyDown(Keys.Left)) {
            _player2.TryMove(Direction.Left, dt, _map, moveBounds);
        } else if (keyState.IsKeyDown(Keys.Right)) {
            _player2.TryMove(Direction.Right, dt, _map, moveBounds);
        }

        // P2 射击用 RightShift
        if (keyState.IsKeyDown(Keys.RightShift) && !_prevKeyState.IsKeyDown(Keys.RightShift)) {
            if (_player2.ShootCooldown <= 0) {
                FireBullet(_player2, true);
                _player2.ShootCooldown = _player2.ShootCooldownMax;
                _sound.PlayShoot(true);
            }
        }
    }

    private void UpdateEnemies(GameTime gameTime) {
        // 收集活跃的玩家子弹（用于敌方AI闪避判断）
        List<Bullet> playerBullets = [];
        foreach (var b in _bullets) {
            if (b.IsActive && b.IsPlayerBullet) {
                playerBullets.Add(b);
            }
        }

        // 收集玩家碰撞体
        var playerBounds = new List<Rectangle>();
        if (_player != null && _player.IsActive) {
            playerBounds.Add(_player.Bounds);
        }
        if (_player2 != null && _player2.IsActive) {
            playerBounds.Add(_player2.Bounds);
        }

        for (int i = _enemies.Count - 1; i >= 0; i--) {
            var enemy = _enemies[i];
            if (!enemy.IsActive) {
                _enemies.RemoveAt(i);
                _enemiesOnScreen--;
                continue;
            }

            enemy.Update(gameTime, _map, playerBounds, playerBullets);

            // 智能射击：对最近的玩家判断
            var nearestPlayer = GetNearestPlayer(enemy.Center);
            if (enemy.WantsToShoot(nearestPlayer?.Bounds ?? Rectangle.Empty)) {
                FireBullet(enemy, false);
                _sound.PlayShoot(false);
            }
        }
    }

    private Tank GetNearestPlayer(Vector2 from) {
        float dist1 = _player != null && _player.IsActive
            ? Vector2.Distance(from, _player.Center) : float.MaxValue;
        float dist2 = _player2 != null && _player2.IsActive
            ? Vector2.Distance(from, _player2.Center) : float.MaxValue;
        return dist1 <= dist2 ? _player : _player2;
    }

    private void UpdateBullets(GameTime gameTime) {
        for (int i = _bullets.Count - 1; i >= 0; i--) {
            _bullets[i].Update(gameTime, _map);
            if (!_bullets[i].IsActive) {
                _bullets.RemoveAt(i);
            }
        }
    }

    private void UpdateExplosions(GameTime gameTime) {
        for (int i = _explosions.Count - 1; i >= 0; i--) {
            _explosions[i].Update(gameTime);
            if (!_explosions[i].IsActive) {
                _explosions.RemoveAt(i);
            }
        }
    }

    private void UpdatePowerUps(GameTime gameTime) {
        for (int i = _powerUps.Count - 1; i >= 0; i--) {
            _powerUps[i].Update(gameTime);
            if (!_powerUps[i].IsActive) {
                _powerUps.RemoveAt(i);
                continue;
            }

            // 检测玩家拾取
            bool picked = false;
            if (_player != null && _player.IsActive && _powerUps[i].Bounds.Intersects(_player.Bounds)) {
                ApplyPowerUp(_powerUps[i], true);
                picked = true;
            } else if (_player2 != null && _player2.IsActive && _powerUps[i].Bounds.Intersects(_player2.Bounds)) {
                ApplyPowerUp(_powerUps[i], false);
                picked = true;
            }

            if (picked) {
                _sound.PlayPickup();
                _powerUps.RemoveAt(i);
            }
        }
    }

    /// <summary>
    /// 应用道具效果
    /// </summary>
    private void ApplyPowerUp(PowerUp powerUp, bool isP1) {
        switch (powerUp.Type) {
            case PowerUpType.Star:
                if (isP1) {
                    if (_starLevel < GameConstants.MaxStarLevel) {
                        _player.Speed *= GameConstants.StarSpeedMultiplier;
                        _player.ShootCooldownMax *= GameConstants.StarCooldownMultiplier;
                        _starLevel++;
                    }
                    _starTimer = GameConstants.StarDuration;
                } else {
                    if (_starLevelP2 < GameConstants.MaxStarLevel) {
                        _player2.Speed *= GameConstants.StarSpeedMultiplier;
                        _player2.ShootCooldownMax *= GameConstants.StarCooldownMultiplier;
                        _starLevelP2++;
                    }
                    _starTimerP2 = GameConstants.StarDuration;
                }
                break;

            case PowerUpType.Helmet:
                if (isP1) {
                    _invincibleTimer = GameConstants.HelmetPowerUpDuration;
                } else {
                    _invincibleTimerP2 = GameConstants.HelmetPowerUpDuration;
                }
                break;

            case PowerUpType.Bomb:
                foreach (var enemy in _enemies) {
                    if (enemy.IsActive) {
                        _explosions.Add(new Explosion(GraphicsDevice, enemy.Center, true));
                        _score += GetEnemyScore(enemy.Type);
                        enemy.IsActive = false;
                        _sound.PlayExplosion(true);
                    }
                }
                break;

            case PowerUpType.Shovel:
                _map.FortifyBase();
                _shovelTimer = GameConstants.ShovelFortifyDuration;
                break;

            case PowerUpType.Tank:
                _lives++;
                break;

            case PowerUpType.Timer:
                _freezeTimer = GameConstants.FreezeDuration;
                break;
        }
    }

    private void ResetStarEffect(ref Tank player, ref int starLevel) {
        if (player != null && starLevel > 0) {
            float speedRatio = MathF.Pow(GameConstants.StarSpeedMultiplier, starLevel);
            float cdRatio = MathF.Pow(GameConstants.StarCooldownMultiplier, starLevel);
            player.Speed = GameConstants.PlayerSpeed;
            player.ShootCooldownMax = 0.3f;
        }
        starLevel = 0;
    }

    private void FireBullet(Tank tank, bool isPlayerBullet) {
        var muzzle = tank.GetMuzzlePosition();
        float speed = isPlayerBullet ? GameConstants.PlayerBulletSpeed : GameConstants.BulletSpeed;

        if (!isPlayerBullet && tank is EnemyTank et && et.Type == EnemyType.Power) {
            speed = GameConstants.PowerBulletSpeed;
        }

        var bullet = new Bullet(GraphicsDevice, muzzle, tank.Facing, speed, isPlayerBullet);
        _bullets.Add(bullet);
    }

    private static int GetEnemyScore(EnemyType type) => type switch {
        EnemyType.Basic => 100,
        EnemyType.Fast => 200,
        EnemyType.Power => 300,
        EnemyType.Armor => 400,
        _ => 100
    };

    private void CheckCollisions() {
        for (int i = _bullets.Count - 1; i >= 0; i--) {
            var bullet = _bullets[i];
            if (!bullet.IsActive) {
                continue;
            }

            if (bullet.IsPlayerBullet) {
                // 玩家子弹 vs 敌方坦克
                foreach (var enemy in _enemies) {
                    if (!enemy.IsActive) {
                        continue;
                    }
                    if (bullet.Bounds.Intersects(enemy.Bounds)) {
                        enemy.TakeDamage();
                        bullet.IsActive = false;

                        if (!enemy.IsActive) {
                            _explosions.Add(new Explosion(GraphicsDevice, enemy.Center, true));
                            _score += GetEnemyScore(enemy.Type);
                            _sound.PlayExplosion(true);

                            if (enemy.IsSpecial) {
                                SpawnPowerUp(enemy.Position);
                                _sound.PlayPowerUpSpawn();
                            }
                        } else {
                            _explosions.Add(new Explosion(GraphicsDevice, bullet.Position, false));
                            _sound.PlayBulletHit();
                        }
                        break;
                    }
                }
            } else {
                // 敌方子弹 vs 玩家1
                if (_player != null && _player.IsActive && _invincibleTimer <= 0) {
                    if (bullet.Bounds.Intersects(_player.Bounds)) {
                        bullet.IsActive = false;
                        PlayerHit(ref _player, ref _invincibleTimer, ref _respawnTimer);
                        break;
                    }
                }
                // 敌方子弹 vs 玩家2
                if (_player2 != null && _player2.IsActive && _invincibleTimerP2 <= 0) {
                    if (bullet.Bounds.Intersects(_player2.Bounds)) {
                        bullet.IsActive = false;
                        PlayerHit(ref _player2, ref _invincibleTimerP2, ref _respawnTimerP2);
                        break;
                    }
                }
            }

            // 子弹 vs 子弹
            for (int j = i - 1; j >= 0; j--) {
                var other = _bullets[j];
                if (!other.IsActive || bullet.IsPlayerBullet == other.IsPlayerBullet) {
                    continue;
                }
                if (bullet.Bounds.Intersects(other.Bounds)) {
                    bullet.IsActive = false;
                    other.IsActive = false;
                    _explosions.Add(new Explosion(GraphicsDevice, bullet.Position, false));
                    _sound.PlayBulletHit();
                    break;
                }
            }
        }
    }

    private void PlayerHit(ref Tank player, ref float invincibleTimer, ref float respawnTimer) {
        player.IsActive = false;
        _explosions.Add(new Explosion(GraphicsDevice, player.Center, true));
        _lives--;
        _sound.PlayExplosion(true);

        if (_lives > 0) {
            respawnTimer = 1.0f;
        }
    }

    private void SpawnPowerUp(Vector2 position) {
        var types = Enum.GetValues<PowerUpType>();
        var type = types[Random.Shared.Next(types.Length)];
        _powerUps.Add(new PowerUp(GraphicsDevice, type, position));
    }

    private void CheckGameState() {
        // 基地被摧毁
        if (_map.IsBaseDestroyed()) {
            if (_score > _highScore) {
                _highScore = _score;
                SaveHighScore();
            }
            _state = GameState.GameOver;
            _gameOverTimer = 3f;
            _sound.PlayGameOver();
            return;
        }

        // 所有玩家生命耗尽
        bool p1Dead = _player == null || !_player.IsActive;
        bool p2Dead = !_isTwoPlayer || _player2 == null || !_player2.IsActive;
        if (_lives <= 0 && p1Dead && p2Dead) {
            if (_score > _highScore) {
                _highScore = _score;
                SaveHighScore();
            }
            _state = GameState.GameOver;
            _gameOverTimer = 3f;
            _sound.PlayGameOver();
            return;
        }

        // 所有敌方被消灭
        if (_enemiesRemaining <= 0 && _enemiesOnScreen <= 0) {
            _state = GameState.LevelComplete;
            _levelCompleteTimer = 2f;
            _sound.PlayLevelComplete();
        }
    }

    protected override void Draw(GameTime gameTime) {
        GraphicsDevice.Clear(Color.Black);

        _spriteBatch.Begin(SpriteSortMode.Deferred, BlendState.AlphaBlend, SamplerState.PointClamp);

        switch (_state) {
            case GameState.Title:
                DrawTitle();
                break;
            case GameState.Playing:
            case GameState.LevelComplete:
                DrawGame();
                break;
            case GameState.GameOver:
                DrawGame();
                DrawGameOver();
                break;
            case GameState.Victory:
                DrawGame();
                DrawVictory();
                break;
            case GameState.Paused:
                DrawGame();
                DrawPaused();
                break;
            case GameState.Transition:
                DrawGame();
                _transition?.Draw(_spriteBatch, _font, GameConstants.GameAreaWidth, GameConstants.GameAreaHeight);
                break;
        }

        _spriteBatch.End();

        base.Draw(gameTime);
    }

    private void DrawTitle() {
        string titleText = "TANK BATTLE";
        string modeText = "[1] SINGLE PLAYER    [2] TWO PLAYERS";
        string startText = "OR PRESS ENTER";
        string controlTextP1 = "P1: WASD + SPACE";
        string controlTextP2 = "P2: ARROWS + RSHIFT";
        string pauseText = "P: Pause";

        Vector2 titleSize = _font.MeasureString(titleText);
        Vector2 modeSize = _font.MeasureString(modeText);
        Vector2 startSize = _font.MeasureString(startText);
        Vector2 p1Size = _font.MeasureString(controlTextP1);
        Vector2 p2Size = _font.MeasureString(controlTextP2);
        Vector2 pauseSize = _font.MeasureString(pauseText);

        float centerX = GameConstants.WindowWidth / 2f;
        float centerY = GameConstants.WindowHeight / 2f;

        _spriteBatch.DrawString(_font, titleText,
            new Vector2(centerX - titleSize.X / 2, centerY - 100), Color.Yellow, 0, Vector2.Zero, 2f, SpriteEffects.None, 0);
        _spriteBatch.DrawString(_font, modeText,
            new Vector2(centerX - modeSize.X / 2, centerY - 20), Color.Cyan);
        _spriteBatch.DrawString(_font, startText,
            new Vector2(centerX - startSize.X / 2, centerY + 10), Color.White);
        _spriteBatch.DrawString(_font, controlTextP1,
            new Vector2(centerX - p1Size.X / 2, centerY + 45), Color.Gray);
        _spriteBatch.DrawString(_font, controlTextP2,
            new Vector2(centerX - p2Size.X / 2, centerY + 65), Color.Gray);
        _spriteBatch.DrawString(_font, pauseText,
            new Vector2(centerX - pauseSize.X / 2, centerY + 90), Color.Gray);

        if (_highScore > 0) {
            string hiText = $"HI-SCORE: {_highScore}";
            var hiSize = _font.MeasureString(hiText);
            _spriteBatch.DrawString(_font, hiText,
                new Vector2(centerX - hiSize.X / 2, centerY - 130), Color.Gold);
        }
    }

    private void DrawGame() {
        _map.Draw(_spriteBatch, false);

        // 绘制玩家1
        DrawPlayer(_player, _invincibleTimer);

        // 绘制玩家2
        if (_isTwoPlayer) {
            DrawPlayer(_player2, _invincibleTimerP2);
        }

        foreach (var enemy in _enemies) {
            enemy.DrawEnemy(_spriteBatch);
        }

        foreach (var bullet in _bullets) {
            bullet.Draw(_spriteBatch);
        }

        foreach (var powerUp in _powerUps) {
            powerUp.Draw(_spriteBatch);
        }

        _map.Draw(_spriteBatch, true);

        foreach (var explosion in _explosions) {
            explosion.Draw(_spriteBatch);
        }

        DrawSidePanel();
    }

    private void DrawPlayer(Tank player, float invincibleTimer) {
        if (player != null && player.IsActive) {
            if (invincibleTimer > 0 && !_blinkVisible) {
                // 无敌闪烁
            } else {
                player.Draw(_spriteBatch);
            }
        }
    }

    private void DrawSidePanel() {
        int panelX = GameConstants.GameAreaWidth + 4;
        int panelY = 10;

        _spriteBatch.Draw(_pixelTexture,
            new Rectangle(GameConstants.GameAreaWidth, 0, GameConstants.InfoPanelWidth, GameConstants.WindowHeight),
            new Color(40, 40, 0));

        _spriteBatch.DrawString(_font, "HI-SCORE", new Vector2(panelX, panelY), Color.Gold);
        panelY += 18;
        _spriteBatch.DrawString(_font, _highScore.ToString(), new Vector2(panelX, panelY), Color.White);
        panelY += 25;

        _spriteBatch.DrawString(_font, $"STAGE", new Vector2(panelX, panelY), Color.White);
        panelY += 18;
        _spriteBatch.DrawString(_font, _level.ToString(), new Vector2(panelX, panelY), Color.Cyan);
        panelY += 25;

        _spriteBatch.DrawString(_font, "ENEMY", new Vector2(panelX, panelY), Color.White);
        panelY += 20;

        int iconsPerRow = 3;
        for (int i = 0; i < _enemiesRemaining; i++) {
            int row = i / iconsPerRow;
            int col = i % iconsPerRow;
            _spriteBatch.Draw(_enemyIconTexture,
                new Vector2(panelX + col * 12, panelY + row * 12),
                Color.Red);
        }
        panelY += (_enemiesRemaining + iconsPerRow - 1) / iconsPerRow * 12 + 20;

        _spriteBatch.DrawString(_font, "LIFE", new Vector2(panelX, panelY), Color.White);
        panelY += 20;
        for (int i = 0; i < _lives; i++) {
            _spriteBatch.Draw(_lifeIconTexture,
                new Vector2(panelX + i * 12, panelY),
                Color.Green);
        }
        panelY += 30;

        _spriteBatch.DrawString(_font, "SCORE", new Vector2(panelX, panelY), Color.White);
        panelY += 18;
        _spriteBatch.DrawString(_font, _score.ToString(), new Vector2(panelX, panelY), Color.Yellow);

        // 双人模式指示
        panelY += 25;
        if (_isTwoPlayer) {
            _spriteBatch.DrawString(_font, "2PLAY", new Vector2(panelX, panelY), Color.Cyan);
        }

        // 道具效果指示
        panelY += 20;
        if (_freezeTimer > 0) {
            _spriteBatch.DrawString(_font, "FREEZE", new Vector2(panelX, panelY), Color.Cyan);
        }
        if (_shovelTimer > 0) {
            _spriteBatch.DrawString(_font, "FORTIFY", new Vector2(panelX, panelY + 16), Color.Orange);
        }
        if (_starTimer > 0) {
            _spriteBatch.DrawString(_font, $"STAR x{_starLevel}", new Vector2(panelX, panelY + 32), Color.Gold);
        }
    }

    private void DrawGameOver() {
        string text = "GAME OVER";
        var size = _font.MeasureString(text);
        float centerX = GameConstants.GameAreaWidth / 2f;
        float centerY = GameConstants.GameAreaHeight / 2f;

        _spriteBatch.Draw(_pixelTexture,
            new Rectangle(0, 0, GameConstants.GameAreaWidth, GameConstants.GameAreaHeight),
            Color.Black * 0.6f);

        _spriteBatch.DrawString(_font, text,
            new Vector2(centerX - size.X, centerY - size.Y - 20), Color.Red, 0, Vector2.Zero, 2f, SpriteEffects.None, 0);

        string scoreText = $"SCORE: {_score}";
        var scoreSize = _font.MeasureString(scoreText);
        _spriteBatch.DrawString(_font, scoreText,
            new Vector2(centerX - scoreSize.X / 2, centerY + 10), Color.Yellow);

        string hiText = $"HI-SCORE: {_highScore}";
        var hiSize = _font.MeasureString(hiText);
        _spriteBatch.DrawString(_font, hiText,
            new Vector2(centerX - hiSize.X / 2, centerY + 30), Color.White);

        if (_score >= _highScore && _score > 0) {
            string newText = "NEW HIGH SCORE!";
            var newSize = _font.MeasureString(newText);
            _spriteBatch.DrawString(_font, newText,
                new Vector2(centerX - newSize.X / 2, centerY + 55), Color.Gold);
        }

        if (_gameOverTimer <= 0) {
            string restartText = "PRESS ENTER";
            Vector2 restartSize = _font.MeasureString(restartText);
            _spriteBatch.DrawString(_font, restartText,
                new Vector2(centerX - restartSize.X / 2, centerY + 80), Color.White);
        }
    }

    private void DrawVictory() {
        string text = "VICTORY!";
        var size = _font.MeasureString(text);
        float centerX = GameConstants.GameAreaWidth / 2f;
        float centerY = GameConstants.GameAreaHeight / 2f;

        _spriteBatch.Draw(_pixelTexture,
            new Rectangle(0, 0, GameConstants.GameAreaWidth, GameConstants.GameAreaHeight),
            Color.Black * 0.6f);

        _spriteBatch.DrawString(_font, text,
            new Vector2(centerX - size.X, centerY - size.Y - 30), Color.Gold, 0, Vector2.Zero, 2f, SpriteEffects.None, 0);

        string congratsText = "ALL STAGES CLEARED!";
        var congratsSize = _font.MeasureString(congratsText);
        _spriteBatch.DrawString(_font, congratsText,
            new Vector2(centerX - congratsSize.X / 2, centerY + 10), Color.White);

        string scoreText = $"FINAL SCORE: {_score}";
        var scoreSize = _font.MeasureString(scoreText);
        _spriteBatch.DrawString(_font, scoreText,
            new Vector2(centerX - scoreSize.X / 2, centerY + 35), Color.Yellow);

        string hiText = $"HI-SCORE: {_highScore}";
        var hiSize = _font.MeasureString(hiText);
        _spriteBatch.DrawString(_font, hiText,
            new Vector2(centerX - hiSize.X / 2, centerY + 55), Color.Gold);

        if (_gameOverTimer <= 0) {
            string restartText = "PRESS ENTER";
            var restartSize = _font.MeasureString(restartText);
            _spriteBatch.DrawString(_font, restartText,
                new Vector2(centerX - restartSize.X / 2, centerY + 80), Color.White);
        }
    }

    private void DrawPaused() {
        float centerX = GameConstants.GameAreaWidth / 2f;
        float centerY = GameConstants.GameAreaHeight / 2f;

        _spriteBatch.Draw(_pixelTexture,
            new Rectangle(0, 0, GameConstants.GameAreaWidth, GameConstants.GameAreaHeight),
            Color.Black * 0.5f);

        string text = "PAUSED";
        var size = _font.MeasureString(text);
        _spriteBatch.DrawString(_font, text,
            new Vector2(centerX - size.X, centerY - size.Y), Color.White, 0, Vector2.Zero, 2f, SpriteEffects.None, 0);

        string hintText = "PRESS P TO RESUME";
        var hintSize = _font.MeasureString(hintText);
        _spriteBatch.DrawString(_font, hintText,
            new Vector2(centerX - hintSize.X / 2, centerY + 30), Color.Gray);
    }

    protected override void UnloadContent() {
        _pixelTexture?.Dispose();
        _enemyIconTexture?.Dispose();
        _lifeIconTexture?.Dispose();
        _transition?.Dispose();
        base.UnloadContent();
    }
}
