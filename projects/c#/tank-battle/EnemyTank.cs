using System;
using System.Collections.Generic;
using Microsoft.Xna.Framework;
using Microsoft.Xna.Framework.Graphics;

namespace TankBattle;

public class EnemyTank : Tank {
    public EnemyType Type;
    public bool IsSpecial;  // 特殊敌方 - 击败后掉落道具

    private float _directionChangeTimer;
    private readonly float _directionChangeInterval = 2f;
    private float _shootTimer;  // 射击计时器（帧率无关）
    private readonly Random _random = new();

    // 基地位置（用于寻路目标）
    private static readonly Vector2 BasePosition = new(12 * GameConstants.TileSize, 24 * GameConstants.TileSize);

    public EnemyTank(GraphicsDevice graphicsDevice, EnemyType type, Vector2 spawnPosition)
        : base(graphicsDevice, GetColor(type), false) {
        Type = type;
        Position = spawnPosition;
        Facing = Direction.Down;

        switch (type) {
            case EnemyType.Basic:
                Speed = GameConstants.EnemyBasicSpeed;
                Health = 1;
                break;
            case EnemyType.Fast:
                Speed = GameConstants.EnemyFastSpeed;
                Health = 1;
                break;
            case EnemyType.Power:
                Speed = GameConstants.EnemyPowerSpeed;
                Health = 1;
                break;
            case EnemyType.Armor:
                Speed = GameConstants.EnemyArmorSpeed;
                Health = 3;
                break;
        }

        ShootCooldownMax = GameConstants.EnemySmartShootCooldown;
        _directionChangeTimer = _directionChangeInterval;
        _shootTimer = ShootCooldownMax * (0.5f + (float)_random.NextDouble() * 0.5f);
    }

    private static Color GetColor(EnemyType type) {
        return type switch {
            EnemyType.Basic => new Color(200, 200, 200),   // 灰白色
            EnemyType.Fast => new Color(255, 100, 100),    // 红色
            EnemyType.Power => new Color(100, 255, 100),   // 绿色
            EnemyType.Armor => new Color(100, 100, 255),   // 蓝色
            _ => Color.White
        };
    }

    public void Update(GameTime gameTime, GameMap map, List<Rectangle> playerBounds, List<Bullet> playerBullets) {
        if (!IsActive) {
            return;
        }

        float dt = (float)gameTime.ElapsedGameTime.TotalSeconds;

        // 射击冷却
        if (ShootCooldown > 0) {
            ShootCooldown -= dt;
        }

        // 定时改变方向
        _directionChangeTimer -= dt;
        if (_directionChangeTimer <= 0) {
            ChangeDirection(playerBounds.Count > 0 ? playerBounds[0] : Rectangle.Empty);
            _directionChangeTimer = _directionChangeInterval * (0.5f + (float)_random.NextDouble());
        }

        // 闪避玩家子弹（装甲型特有行为）
        if (Type == EnemyType.Armor) {
            TryDodgeBullet(playerBullets);
        }

        // 尝试移动，如果撞墙则换方向
        if (!TryMove(Facing, dt, map, playerBounds)) {
            ChangeDirection(playerBounds.Count > 0 ? playerBounds[0] : Rectangle.Empty);
        }
    }

    /// <summary>
    /// 根据敌方类型选择不同的AI行为
    /// </summary>
    private void ChangeDirection(Rectangle playerBounds) {
        Direction[] dirs = [Direction.Up, Direction.Down, Direction.Left, Direction.Right];

        switch (Type) {
            case EnemyType.Basic:
                // 基础型：随机 + 轻微向下偏移
                if (_random.Next(100) < 40) {
                    Facing = Direction.Down;
                } else {
                    Facing = dirs[_random.Next(dirs.Length)];
                }
                break;

            case EnemyType.Fast:
                // 快速型：主动追踪玩家
                if (_random.Next(100) < 60) {
                    Facing = GetDirectionToTarget(new Vector2(playerBounds.X, playerBounds.Y));
                } else {
                    Facing = dirs[_random.Next(dirs.Length)];
                }
                break;

            case EnemyType.Power:
                // 强力型：偏好向基地方向移动
                if (_random.Next(100) < 60) {
                    Facing = GetDirectionToTarget(BasePosition);
                } else {
                    Facing = dirs[_random.Next(dirs.Length)];
                }
                break;

            case EnemyType.Armor:
                // 装甲型：随机移动，但会闪避子弹
                Facing = dirs[_random.Next(dirs.Length)];
                break;
        }
    }

    /// <summary>
    /// 计算朝向目标的方向
    /// </summary>
    private Direction GetDirectionToTarget(Vector2 target) {
        float dx = target.X - Position.X;
        float dy = target.Y - Position.Y;

        // 选择距离更远的轴
        if (MathF.Abs(dx) > MathF.Abs(dy)) {
            return dx > 0 ? Direction.Right : Direction.Left;
        } else {
            return dy > 0 ? Direction.Down : Direction.Up;
        }
    }

    /// <summary>
    /// 尝试闪避玩家子弹
    /// </summary>
    private void TryDodgeBullet(List<Bullet> playerBullets) {
        if (playerBullets == null) {
            return;
        }

        foreach (var bullet in playerBullets) {
            if (!bullet.IsActive || !bullet.IsPlayerBullet) {
                continue;
            }

            // 检测子弹是否在附近且朝向自己
            float dist = Vector2.Distance(bullet.Position, Center);
            if (dist > GameConstants.BulletDetectionRadius) {
                continue;
            }

            // 检查子弹是否正朝自己飞来
            bool incoming = bullet.Direction switch {
                Direction.Up => bullet.Position.Y > Center.Y && MathF.Abs(bullet.Position.X - Center.X) < GameConstants.TankSize,
                Direction.Down => bullet.Position.Y < Center.Y && MathF.Abs(bullet.Position.X - Center.X) < GameConstants.TankSize,
                Direction.Left => bullet.Position.X > Center.X && MathF.Abs(bullet.Position.Y - Center.Y) < GameConstants.TankSize,
                Direction.Right => bullet.Position.X < Center.X && MathF.Abs(bullet.Position.Y - Center.Y) < GameConstants.TankSize,
                _ => false
            };

            if (incoming) {
                // 垂直于子弹方向闪避
                if (bullet.Direction == Direction.Up || bullet.Direction == Direction.Down) {
                    Facing = _random.Next(2) == 0 ? Direction.Left : Direction.Right;
                } else {
                    Facing = _random.Next(2) == 0 ? Direction.Up : Direction.Down;
                }
                return;
            }
        }
    }

    /// <summary>
    /// 智能射击判断（帧率无关，基于冷却时间）
    /// </summary>
    public bool WantsToShoot(Rectangle playerBounds) {
        if (ShootCooldown > 0) {
            return false;
        }

        _shootTimer -= 0.016f; // 模拟约60fps的检查频率
        if (_shootTimer > 0) {
            return false;
        }

        // 重置射击计时器
        _shootTimer = ShootCooldownMax * (0.7f + (float)_random.NextDouble() * 0.6f);

        // 检查是否与玩家或基地方向对齐（增加射击概率）
        bool alignedWithPlayer = IsAlignedWith(playerBounds, 2 * GameConstants.TileSize);
        bool alignedWithBase = IsAlignedWith(
            new Rectangle((int)BasePosition.X, (int)BasePosition.Y, GameConstants.TankSize, GameConstants.TankSize),
            2 * GameConstants.TileSize);

        int shootChance = 10; // 基础10%概率
        if (alignedWithPlayer) {
            shootChance = 50;
        } else if (alignedWithBase) {
            shootChance = 30;
        }

        if (_random.Next(100) < shootChance) {
            ShootCooldown = ShootCooldownMax;
            return true;
        }

        return false;
    }

    /// <summary>
    /// 检查是否与目标在同一行/列（容差范围内）
    /// </summary>
    private bool IsAlignedWith(Rectangle target, float tolerance) {
        bool sameColumn = MathF.Abs(Center.X - (target.X + target.Width / 2f)) < tolerance;
        bool sameRow = MathF.Abs(Center.Y - (target.Y + target.Height / 2f)) < tolerance;
        return sameColumn || sameRow;
    }

    public void DrawEnemy(SpriteBatch spriteBatch) {
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

        var drawColor = TankColor;

        // 装甲型闪烁效果
        if (Type == EnemyType.Armor && Health > 1) {
            drawColor = Color.Lerp(TankColor, Color.White, 0.3f + 0.3f * MathF.Sin(Environment.TickCount * 0.01f));
        }

        // 特殊敌方闪烁红色
        if (IsSpecial) {
            float blink = MathF.Sin(Environment.TickCount * 0.008f);
            drawColor = blink > 0 ? Color.Red : Color.Lerp(TankColor, Color.Red, 0.5f);
        }

        spriteBatch.Draw(_texture, drawPos, null, drawColor, rotation, origin, 1f, SpriteEffects.None, 0f);
    }
}
