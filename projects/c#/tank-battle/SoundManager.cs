using System;
using Microsoft.Xna.Framework.Audio;

namespace TankBattle;

public class SoundManager {
    private SoundEffect _shootSound;
    private SoundEffect _enemyShootSound;
    private SoundEffect _explosionSound;
    private SoundEffect _bigExplosionSound;
    private SoundEffect _pickupSound;
    private SoundEffect _gameOverSound;
    private SoundEffect _levelCompleteSound;
    private SoundEffect _bulletHitSound;
    private SoundEffect _powerUpSpawnSound;

    public void Initialize() {
        _shootSound = GenerateSquare(0.08f, 800, 400, 0.4f);
        _enemyShootSound = GenerateSquare(0.06f, 600, 300, 0.25f);
        _explosionSound = GenerateNoise(0.15f, 0.5f);
        _bigExplosionSound = GenerateNoise(0.3f, 0.7f);
        _pickupSound = GenerateRisingTone(0.15f, 500, 1200, 0.4f);
        _gameOverSound = GenerateFallingTone(0.6f, 600, 150, 0.5f);
        _levelCompleteSound = GenerateArpeggio(0.5f, new[] { 523, 659, 784, 1047 }, 0.4f);
        _bulletHitSound = GenerateNoise(0.05f, 0.3f);
        _powerUpSpawnSound = GenerateRisingTone(0.2f, 300, 900, 0.3f);
    }

    public void PlayShoot(bool isPlayer) {
        try { (isPlayer ? _shootSound : _enemyShootSound)?.Play(0.3f, 0, 0); } catch { }
    }

    public void PlayExplosion(bool isBig) {
        try { (isBig ? _bigExplosionSound : _explosionSound)?.Play(isBig ? 0.5f : 0.3f, 0, 0); } catch { }
    }

    public void PlayPickup() {
        try { _pickupSound?.Play(0.4f, 0, 0); } catch { }
    }

    public void PlayGameOver() {
        try { _gameOverSound?.Play(0.5f, 0, 0); } catch { }
    }

    public void PlayLevelComplete() {
        try { _levelCompleteSound?.Play(0.4f, 0, 0); } catch { }
    }

    public void PlayBulletHit() {
        try { _bulletHitSound?.Play(0.2f, 0, 0); } catch { }
    }

    public void PlayPowerUpSpawn() {
        try { _powerUpSpawnSound?.Play(0.3f, 0, 0); } catch { }
    }

    private static SoundEffect GenerateSquare(float duration, float freqStart, float freqEnd, float volume) {
        int sampleRate = 22050;
        int sampleCount = (int)(sampleRate * duration);
        var buffer = new byte[sampleCount * 2];

        for (int i = 0; i < sampleCount; i++) {
            float t = (float)i / sampleRate;
            float progress = (float)i / sampleCount;
            float freq = freqStart + (freqEnd - freqStart) * progress;
            float envelope = 1f - progress;
            float sample = (MathF.Sin(t * freq * MathF.PI * 2f) > 0 ? 1f : -1f) * envelope * volume;
            short s = (short)(sample * short.MaxValue);
            buffer[i * 2] = (byte)(s & 0xFF);
            buffer[i * 2 + 1] = (byte)((s >> 8) & 0xFF);
        }

        return new SoundEffect(buffer, sampleRate, AudioChannels.Mono);
    }

    private static SoundEffect GenerateNoise(float duration, float volume) {
        var rng = Random.Shared;
        int sampleRate = 22050;
        int sampleCount = (int)(sampleRate * duration);
        var buffer = new byte[sampleCount * 2];

        for (int i = 0; i < sampleCount; i++) {
            float progress = (float)i / sampleCount;
            float envelope = (1f - progress) * (1f - progress);
            float sample = (rng.NextSingle() * 2f - 1f) * envelope * volume;
            short s = (short)(sample * short.MaxValue);
            buffer[i * 2] = (byte)(s & 0xFF);
            buffer[i * 2 + 1] = (byte)((s >> 8) & 0xFF);
        }

        return new SoundEffect(buffer, sampleRate, AudioChannels.Mono);
    }

    private static SoundEffect GenerateRisingTone(float duration, float freqStart, float freqEnd, float volume) {
        int sampleRate = 22050;
        int sampleCount = (int)(sampleRate * duration);
        var buffer = new byte[sampleCount * 2];

        for (int i = 0; i < sampleCount; i++) {
            float t = (float)i / sampleRate;
            float progress = (float)i / sampleCount;
            float freq = freqStart + (freqEnd - freqStart) * progress;
            float envelope = 1f - progress * 0.5f;
            float sample = MathF.Sin(t * freq * MathF.PI * 2f) * envelope * volume;
            short s = (short)(sample * short.MaxValue);
            buffer[i * 2] = (byte)(s & 0xFF);
            buffer[i * 2 + 1] = (byte)((s >> 8) & 0xFF);
        }

        return new SoundEffect(buffer, sampleRate, AudioChannels.Mono);
    }

    private static SoundEffect GenerateFallingTone(float duration, float freqStart, float freqEnd, float volume) {
        int sampleRate = 22050;
        int sampleCount = (int)(sampleRate * duration);
        var buffer = new byte[sampleCount * 2];

        for (int i = 0; i < sampleCount; i++) {
            float t = (float)i / sampleRate;
            float progress = (float)i / sampleCount;
            float freq = freqStart + (freqEnd - freqStart) * progress;
            float envelope = 1f - progress;
            float sample = MathF.Sin(t * freq * MathF.PI * 2f) * envelope * volume;
            short s = (short)(sample * short.MaxValue);
            buffer[i * 2] = (byte)(s & 0xFF);
            buffer[i * 2 + 1] = (byte)((s >> 8) & 0xFF);
        }

        return new SoundEffect(buffer, sampleRate, AudioChannels.Mono);
    }

    private static SoundEffect GenerateArpeggio(float duration, int[] freqs, float volume) {
        int sampleRate = 22050;
        int sampleCount = (int)(sampleRate * duration);
        var buffer = new byte[sampleCount * 2];
        float noteLength = (float)sampleCount / freqs.Length;

        for (int i = 0; i < sampleCount; i++) {
            int noteIdx = Math.Min((int)(i / noteLength), freqs.Length - 1);
            float freq = freqs[noteIdx];
            float t = (float)i / sampleRate;
            float noteProgress = (i % (int)noteLength) / noteLength;
            float envelope = (1f - noteProgress) * 0.8f;
            float sample = MathF.Sin(t * freq * MathF.PI * 2f) * envelope * volume;
            short s = (short)(sample * short.MaxValue);
            buffer[i * 2] = (byte)(s & 0xFF);
            buffer[i * 2 + 1] = (byte)((s >> 8) & 0xFF);
        }

        return new SoundEffect(buffer, sampleRate, AudioChannels.Mono);
    }
}
