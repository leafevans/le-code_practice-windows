using System;
using Microsoft.Xna.Framework;
using Microsoft.Xna.Framework.Graphics;

namespace TankBattle;

public class GameMap {
    private readonly TileType[,] _tiles;
    private Texture2D _brickTexture;
    private Texture2D _steelTexture;
    private Texture2D _waterTexture;
    private Texture2D _treesTexture;
    private Texture2D _baseTexture;
    private Texture2D _baseDestroyedTexture;

    public static int Width => GameConstants.MapWidth;
    public static int Height => GameConstants.MapHeight;

    public GameMap(GraphicsDevice graphicsDevice) {
        _tiles = new TileType[GameConstants.MapWidth, GameConstants.MapHeight];
        CreateTextures(graphicsDevice);
    }

    private void CreateTextures(GraphicsDevice graphicsDevice) {
        int ts = GameConstants.TileSize;

        _brickTexture = new Texture2D(graphicsDevice, ts, ts);
        var brickData = new Color[ts * ts];
        for (int i = 0; i < brickData.Length; i++) {
            int x = i % ts;
            int y = i / ts;
            bool isMortar = (x % 8 == 0) || (y % 4 == 0);
            brickData[i] = isMortar ? new Color(139, 90, 43) : new Color(180, 100, 50);
        }
        _brickTexture.SetData(brickData);

        _steelTexture = new Texture2D(graphicsDevice, ts, ts);
        var steelData = new Color[ts * ts];
        for (int i = 0; i < steelData.Length; i++) {
            int x = i % ts;
            int y = i / ts;
            bool isEdge = x == 0 || y == 0;
            steelData[i] = isEdge ? new Color(200, 200, 200) : new Color(160, 160, 160);
        }
        _steelTexture.SetData(steelData);

        _waterTexture = new Texture2D(graphicsDevice, ts, ts);
        var waterData = new Color[ts * ts];
        for (int i = 0; i < waterData.Length; i++) {
            int x = i % ts;
            int y = i / ts;
            bool wave = (x + y) % 4 < 2;
            waterData[i] = wave ? new Color(0, 100, 200) : new Color(0, 80, 180);
        }
        _waterTexture.SetData(waterData);

        _treesTexture = new Texture2D(graphicsDevice, ts, ts);
        var treesData = new Color[ts * ts];
        for (int i = 0; i < treesData.Length; i++) {
            int x = i % ts;
            int y = i / ts;
            bool leaf = (x + y * 3) % 5 != 0;
            treesData[i] = leaf ? new Color(0, 140, 0) : new Color(0, 100, 0);
        }
        _treesTexture.SetData(treesData);

        _baseTexture = new Texture2D(graphicsDevice, ts * 2, ts * 2);
        var baseData = new Color[ts * 2 * ts * 2];
        for (int i = 0; i < baseData.Length; i++) {
            int x = i % (ts * 2);
            int y = i / (ts * 2);
            int cx = ts;
            int cy = ts;
            float dist = MathF.Sqrt((x - cx) * (x - cx) + (y - cy) * (y - cy));
            baseData[i] = dist < ts ? new Color(200, 200, 0) : new Color(100, 100, 0);
        }
        _baseTexture.SetData(baseData);

        _baseDestroyedTexture = new Texture2D(graphicsDevice, ts * 2, ts * 2);
        var destroyedData = new Color[ts * 2 * ts * 2];
        for (int i = 0; i < destroyedData.Length; i++) {
            destroyedData[i] = new Color(80, 40, 0);
        }
        _baseDestroyedTexture.SetData(destroyedData);
    }

    public void LoadLevel(int levelNumber) {
        for (int x = 0; x < Width; x++) {
            for (int y = 0; y < Height; y++) {
                _tiles[x, y] = TileType.Empty;
            }
        }

        _tiles[12, 24] = TileType.Base;
        _tiles[13, 24] = TileType.Base;
        _tiles[12, 25] = TileType.Base;
        _tiles[13, 25] = TileType.Base;

        switch ((levelNumber - 1) % 5) {
            case 0:
                LoadLevel1();
                break;
            case 1:
                LoadLevel2();
                break;
            case 2:
                LoadLevel3();
                break;
            case 3:
                LoadLevel4();
                break;
            case 4:
                LoadLevel5();
                break;
        }
    }

    private void PlaceBaseProtection(TileType wallType = TileType.Brick) {
        for (int x = 11; x <= 14; x++) {
            _tiles[x, 23] = wallType;
        }
        _tiles[11, 24] = wallType;
        _tiles[11, 25] = wallType;
        _tiles[14, 24] = wallType;
        _tiles[14, 25] = wallType;
    }

    public void FortifyBase() => PlaceBaseProtection(TileType.Steel);

    public void UnfortifyBase() => PlaceBaseProtection(TileType.Brick);

    private void LoadLevel1() {
        PlaceBaseProtection();

        PlaceWall(2, 2, 2, 4, TileType.Brick);
        PlaceWall(4, 2, 4, 4, TileType.Brick);
        PlaceWall(6, 2, 8, 4, TileType.Brick);
        PlaceWall(10, 2, 12, 4, TileType.Brick);
        PlaceWall(14, 2, 16, 4, TileType.Brick);
        PlaceWall(18, 2, 20, 4, TileType.Brick);
        PlaceWall(22, 2, 22, 4, TileType.Brick);
        PlaceWall(24, 2, 24, 4, TileType.Brick);

        PlaceWall(2, 6, 2, 8, TileType.Brick);
        PlaceWall(4, 6, 4, 8, TileType.Brick);
        PlaceWall(6, 6, 8, 8, TileType.Steel);
        PlaceWall(10, 6, 10, 8, TileType.Brick);
        PlaceWall(12, 6, 14, 8, TileType.Brick);
        PlaceWall(16, 6, 16, 8, TileType.Brick);
        PlaceWall(18, 6, 20, 8, TileType.Steel);
        PlaceWall(22, 6, 22, 8, TileType.Brick);
        PlaceWall(24, 6, 24, 8, TileType.Brick);

        PlaceWall(6, 10, 8, 12, TileType.Brick);
        PlaceWall(10, 10, 10, 12, TileType.Brick);
        PlaceWall(12, 10, 14, 12, TileType.Water);
        PlaceWall(16, 10, 16, 12, TileType.Brick);
        PlaceWall(18, 10, 20, 12, TileType.Brick);

        PlaceWall(2, 14, 2, 16, TileType.Brick);
        PlaceWall(4, 14, 4, 16, TileType.Brick);
        PlaceWall(6, 14, 6, 16, TileType.Brick);
        PlaceWall(8, 14, 8, 16, TileType.Brick);
        PlaceWall(10, 14, 12, 16, TileType.Trees);
        PlaceWall(14, 14, 14, 16, TileType.Brick);
        PlaceWall(16, 14, 16, 16, TileType.Brick);
        PlaceWall(18, 14, 18, 16, TileType.Brick);
        PlaceWall(20, 14, 20, 16, TileType.Brick);
        PlaceWall(22, 14, 22, 16, TileType.Brick);
        PlaceWall(24, 14, 24, 16, TileType.Brick);

        PlaceWall(2, 18, 4, 20, TileType.Brick);
        PlaceWall(6, 18, 8, 20, TileType.Brick);
        PlaceWall(10, 18, 12, 20, TileType.Brick);
        PlaceWall(14, 18, 16, 20, TileType.Brick);
        PlaceWall(18, 18, 20, 20, TileType.Brick);
        PlaceWall(22, 18, 24, 20, TileType.Brick);
    }

    private void LoadLevel2() {
        for (int x = 11; x <= 14; x++) {
            _tiles[x, 23] = TileType.Brick;
        }

        PlaceWall(0, 6, 5, 6, TileType.Water);
        PlaceWall(8, 6, 17, 6, TileType.Water);
        PlaceWall(20, 6, 25, 6, TileType.Water);
        PlaceWall(0, 18, 5, 18, TileType.Water);
        PlaceWall(8, 18, 17, 18, TileType.Water);
        PlaceWall(20, 18, 25, 18, TileType.Water);

        PlaceWall(2, 2, 2, 4, TileType.Brick);
        PlaceWall(4, 2, 4, 4, TileType.Brick);
        PlaceWall(2, 8, 2, 10, TileType.Brick);
        PlaceWall(4, 8, 4, 10, TileType.Brick);

        PlaceWall(10, 9, 15, 9, TileType.Steel);
        PlaceWall(10, 12, 15, 12, TileType.Steel);
        PlaceWall(10, 9, 10, 12, TileType.Brick);
        PlaceWall(15, 9, 15, 12, TileType.Brick);
        PlaceWall(12, 10, 13, 11, TileType.Brick);

        PlaceWall(22, 2, 22, 4, TileType.Brick);
        PlaceWall(24, 2, 24, 4, TileType.Brick);
        PlaceWall(22, 8, 22, 10, TileType.Brick);
        PlaceWall(24, 8, 24, 10, TileType.Brick);

        PlaceWall(6, 14, 8, 16, TileType.Brick);
        PlaceWall(10, 14, 12, 16, TileType.Trees);
        PlaceWall(14, 14, 16, 16, TileType.Brick);
        PlaceWall(18, 14, 20, 16, TileType.Trees);

        PlaceWall(2, 20, 4, 22, TileType.Brick);
        PlaceWall(6, 20, 8, 22, TileType.Steel);
        PlaceWall(10, 20, 12, 22, TileType.Brick);
        PlaceWall(14, 20, 16, 22, TileType.Brick);
        PlaceWall(18, 20, 20, 22, TileType.Steel);
        PlaceWall(22, 20, 24, 22, TileType.Brick);
    }

    private void LoadLevel3() {
        PlaceWall(2, 2, 2, 6, TileType.Steel);
        PlaceWall(4, 2, 8, 2, TileType.Steel);
        PlaceWall(8, 2, 8, 6, TileType.Brick);
        PlaceWall(4, 4, 6, 4, TileType.Brick);

        PlaceWall(18, 2, 18, 6, TileType.Steel);
        PlaceWall(18, 2, 22, 2, TileType.Steel);
        PlaceWall(22, 2, 22, 6, TileType.Brick);
        PlaceWall(20, 4, 20, 4, TileType.Brick);

        PlaceWall(4, 8, 10, 8, TileType.Brick);
        PlaceWall(16, 8, 22, 8, TileType.Brick);
        PlaceWall(12, 8, 13, 10, TileType.Water);

        PlaceWall(12, 10, 13, 10, TileType.Steel);
        PlaceWall(12, 12, 13, 12, TileType.Steel);
        PlaceWall(10, 10, 10, 12, TileType.Steel);
        PlaceWall(15, 10, 15, 12, TileType.Steel);

        PlaceWall(2, 10, 4, 12, TileType.Trees);
        PlaceWall(22, 10, 24, 12, TileType.Trees);

        PlaceWall(2, 14, 6, 14, TileType.Brick);
        PlaceWall(8, 14, 8, 18, TileType.Brick);
        PlaceWall(10, 16, 16, 16, TileType.Water);
        PlaceWall(18, 14, 18, 18, TileType.Brick);
        PlaceWall(20, 14, 24, 14, TileType.Brick);

        PlaceWall(4, 16, 6, 18, TileType.Steel);
        PlaceWall(20, 16, 22, 18, TileType.Steel);

        PlaceWall(2, 20, 2, 22, TileType.Brick);
        PlaceWall(4, 20, 6, 20, TileType.Brick);
        PlaceWall(8, 20, 8, 22, TileType.Steel);
        PlaceWall(10, 20, 12, 22, TileType.Brick);
        PlaceWall(14, 20, 16, 22, TileType.Brick);
        PlaceWall(18, 20, 18, 22, TileType.Steel);
        PlaceWall(20, 20, 22, 20, TileType.Brick);
        PlaceWall(24, 20, 24, 22, TileType.Brick);
    }

    private void LoadLevel4() {
        PlaceBaseProtection(TileType.Steel);

        // 中央十字钢墙
        PlaceWall(12, 4, 13, 4, TileType.Steel);
        PlaceWall(12, 8, 13, 8, TileType.Steel);
        PlaceWall(12, 12, 13, 12, TileType.Steel);
        PlaceWall(12, 16, 13, 16, TileType.Steel);
        PlaceWall(8, 10, 8, 15, TileType.Steel);
        PlaceWall(17, 10, 17, 15, TileType.Steel);

        // 左侧迷宫
        PlaceWall(2, 2, 4, 2, TileType.Brick);
        PlaceWall(2, 2, 2, 6, TileType.Brick);
        PlaceWall(2, 6, 6, 6, TileType.Brick);
        PlaceWall(6, 2, 6, 6, TileType.Water);
        PlaceWall(4, 4, 4, 4, TileType.Trees);

        // 右侧迷宫
        PlaceWall(20, 2, 22, 2, TileType.Brick);
        PlaceWall(22, 2, 22, 6, TileType.Brick);
        PlaceWall(18, 6, 22, 6, TileType.Brick);
        PlaceWall(18, 2, 18, 6, TileType.Water);
        PlaceWall(20, 4, 20, 4, TileType.Trees);

        // 中层横向屏障
        PlaceWall(0, 10, 5, 10, TileType.Brick);
        PlaceWall(20, 10, 25, 10, TileType.Brick);
        PlaceWall(3, 12, 5, 14, TileType.Trees);
        PlaceWall(20, 12, 22, 14, TileType.Trees);

        // 下层防御
        PlaceWall(2, 18, 6, 18, TileType.Brick);
        PlaceWall(2, 18, 2, 22, TileType.Brick);
        PlaceWall(6, 18, 6, 22, TileType.Steel);
        PlaceWall(18, 18, 22, 18, TileType.Brick);
        PlaceWall(22, 18, 22, 22, TileType.Brick);
        PlaceWall(18, 18, 18, 22, TileType.Steel);

        // 基地前方水障碍
        PlaceWall(10, 22, 10, 22, TileType.Water);
        PlaceWall(15, 22, 15, 22, TileType.Water);
    }

    private void LoadLevel5() {
        PlaceBaseProtection(TileType.Steel);

        // 对称钢墙阵列
        PlaceWall(2, 2, 4, 2, TileType.Steel);
        PlaceWall(20, 2, 22, 2, TileType.Steel);
        PlaceWall(2, 6, 2, 8, TileType.Steel);
        PlaceWall(22, 6, 22, 8, TileType.Steel);

        // 中央堡垒
        PlaceWall(10, 4, 15, 4, TileType.Brick);
        PlaceWall(10, 4, 10, 8, TileType.Brick);
        PlaceWall(15, 4, 15, 8, TileType.Brick);
        PlaceWall(10, 8, 15, 8, TileType.Steel);
        PlaceWall(12, 5, 13, 7, TileType.Water);

        // 水域分割带
        PlaceWall(0, 12, 8, 12, TileType.Water);
        PlaceWall(16, 12, 25, 12, TileType.Water);
        PlaceWall(12, 12, 13, 12, TileType.Brick); // 桥

        // 下半部复杂地形
        PlaceWall(2, 14, 4, 16, TileType.Trees);
        PlaceWall(6, 14, 8, 14, TileType.Brick);
        PlaceWall(6, 16, 8, 18, TileType.Steel);
        PlaceWall(16, 14, 18, 14, TileType.Brick);
        PlaceWall(16, 16, 18, 18, TileType.Steel);
        PlaceWall(20, 14, 22, 16, TileType.Trees);

        // 蛇形通道
        PlaceWall(4, 20, 10, 20, TileType.Brick);
        PlaceWall(14, 20, 20, 20, TileType.Brick);
        PlaceWall(8, 20, 8, 22, TileType.Brick);
        PlaceWall(16, 20, 16, 22, TileType.Brick);

        // 基地前方最后防线
        PlaceWall(10, 22, 10, 23, TileType.Steel);
        PlaceWall(15, 22, 15, 23, TileType.Steel);
    }

    private void PlaceWall(int x1, int y1, int x2, int y2, TileType type) {
        for (int x = x1; x <= x2 && x < Width; x++) {
            for (int y = y1; y <= y2 && y < Height; y++) {
                _tiles[x, y] = type;
            }
        }
    }

    public TileType GetTile(int x, int y) {
        if (x < 0 || x >= Width || y < 0 || y >= Height) {
            return TileType.Steel;
        }
        return _tiles[x, y];
    }

    public void SetTile(int x, int y, TileType type) {
        if (x >= 0 && x < Width && y >= 0 && y < Height) {
            _tiles[x, y] = type;
        }
    }

    public bool IsPassable(int tileX, int tileY) {
        var tile = GetTile(tileX, tileY);
        return tile == TileType.Empty || tile == TileType.Trees;
    }

    public bool HitTile(int tileX, int tileY) {
        var tile = GetTile(tileX, tileY);
        switch (tile) {
            case TileType.Brick:
                SetTile(tileX, tileY, TileType.Empty);
                return true;
            case TileType.Steel:
                return true;
            case TileType.Base:
                SetTile(tileX, tileY, TileType.BaseDestroyed);
                return true;
            default:
                return false;
        }
    }

    public bool IsBaseDestroyed() {
        for (int x = 12; x <= 13; x++) {
            for (int y = 24; y <= 25; y++) {
                if (_tiles[x, y] == TileType.Base) {
                    return false;
                }
            }
        }
        return true;
    }

    public void Draw(SpriteBatch spriteBatch, bool drawTreesLayer = false) {
        int ts = GameConstants.TileSize;

        for (int x = 0; x < Width; x++) {
            for (int y = 0; y < Height; y++) {
                var tile = _tiles[x, y];
                Texture2D tex;

                if (drawTreesLayer) {
                    if (tile == TileType.Trees) {
                        tex = _treesTexture;
                    } else {
                        continue;
                    }
                } else {
                    switch (tile) {
                        case TileType.Brick:
                            tex = _brickTexture;
                            break;
                        case TileType.Steel:
                            tex = _steelTexture;
                            break;
                        case TileType.Water:
                            tex = _waterTexture;
                            break;
                        case TileType.Base:
                            tex = _baseTexture;
                            break;
                        case TileType.BaseDestroyed:
                            tex = _baseDestroyedTexture;
                            break;
                        default:
                            continue;
                    }
                }

                if (tex != null) {
                    Rectangle destRect = new(x * ts, y * ts, ts, ts);
                    if (tile == TileType.Base || tile == TileType.BaseDestroyed) {
                        if (x == 12 && y == 24) {
                            destRect = new(x * ts, y * ts, ts * 2, ts * 2);
                            spriteBatch.Draw(tex, destRect, Color.White);
                        }
                    } else {
                        spriteBatch.Draw(tex, destRect, Color.White);
                    }
                }
            }
        }
    }
}
