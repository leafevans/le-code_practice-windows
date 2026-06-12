using System;
using System.Collections.Generic;
using Microsoft.Xna.Framework;
using Microsoft.Xna.Framework.Graphics;

namespace AnimalWellLight.Core.Effects
{
    /// <summary>
    /// Represents a point light source in the game world.
    /// </summary>
    public class PointLight
    {
        /// <summary>
        /// Position of the light in world space.
        /// </summary>
        public Vector2 Position { get; set; }

        /// <summary>
        /// Color of the light.
        /// </summary>
        public Color Color { get; set; }

        /// <summary>
        /// Radius of the light effect.
        /// </summary>
        public float Radius { get; set; }

        /// <summary>
        /// Intensity of the light (0.0 to 1.0).
        /// </summary>
        public float Intensity { get; set; }

        /// <summary>
        /// Whether the light is currently active.
        /// </summary>
        public bool IsActive { get; set; } = true;

        /// <summary>
        /// Creates a new point light.
        /// </summary>
        /// <param name="position">World position.</param>
        /// <param name="color">Light color.</param>
        /// <param name="radius">Effect radius.</param>
        /// <param name="intensity">Light intensity.</param>
        public PointLight(Vector2 position, Color color, float radius, float intensity)
        {
            Position = position;
            Color = color;
            Radius = radius;
            Intensity = intensity;
        }
    }

    /// <summary>
    /// Manages Animal Well-style lighting effects including darkness, point lights, and atmospheric fog.
    /// </summary>
    public class LightingManager : IDisposable
    {
        private GraphicsDevice graphicsDevice;
        private SpriteBatch spriteBatch;
        private Effect lightingEffect;
        private Texture2D lightTexture;
        private Texture2D radialGradientTexture;

        // Render targets for multi-pass rendering
        private RenderTarget2D sceneTarget;
        private RenderTarget2D lightMapTarget;

        // Lighting parameters
        private float ambientDarkness = 0.85f;
        private float lightIntensity = 1.2f;
        private Color fogColor = new Color(15, 10, 35); // Dark purple/blue
        private float vignetteStrength = 1.5f;
        private float time;

        // Light sources
        private List<PointLight> lights = new List<PointLight>();

        // Screen dimensions
        private int screenWidth;
        private int screenHeight;

        /// <summary>
        /// Gets or sets the ambient darkness level (0.0 = fully lit, 1.0 = fully dark).
        /// </summary>
        public float AmbientDarkness
        {
            get => ambientDarkness;
            set => ambientDarkness = MathHelper.Clamp(value, 0f, 1f);
        }

        /// <summary>
        /// Gets or sets the light intensity multiplier.
        /// </summary>
        public float LightIntensityMultiplier
        {
            get => lightIntensity;
            set => lightIntensity = MathHelper.Clamp(value, 0f, 3f);
        }

        /// <summary>
        /// Gets or sets the atmospheric fog color.
        /// </summary>
        public Color FogColor
        {
            get => fogColor;
            set => fogColor = value;
        }

        /// <summary>
        /// Gets or sets the vignette effect strength.
        /// </summary>
        public float VignetteStrength
        {
            get => vignetteStrength;
            set => vignetteStrength = MathHelper.Clamp(value, 0f, 3f);
        }

        /// <summary>
        /// Gets the list of active lights.
        /// </summary>
        public List<PointLight> Lights => lights;

        /// <summary>
        /// Creates a new LightingManager instance.
        /// </summary>
        /// <param name="graphicsDevice">The graphics device for rendering.</param>
        /// <param name="content">Content manager for loading shaders.</param>
        /// <param name="width">Screen width.</param>
        /// <param name="height">Screen height.</param>
        public LightingManager(GraphicsDevice graphicsDevice, Microsoft.Xna.Framework.Content.ContentManager content, int width, int height)
        {
            this.graphicsDevice = graphicsDevice;
            this.screenWidth = width;
            this.screenHeight = height;

            spriteBatch = new SpriteBatch(graphicsDevice);

            // Load the lighting effect
            lightingEffect = content.Load<Effect>("Effects/Lighting");

            // Create render targets
            sceneTarget = new RenderTarget2D(graphicsDevice, width, height, false, SurfaceFormat.Color, DepthFormat.None);
            lightMapTarget = new RenderTarget2D(graphicsDevice, width, height, false, SurfaceFormat.Color, DepthFormat.None);

            // Create light textures
            CreateLightTextures();
        }

        /// <summary>
        /// Creates textures used for rendering lights.
        /// </summary>
        private void CreateLightTextures()
        {
            // Create a simple radial gradient texture for point lights
            int size = 256;
            radialGradientTexture = new Texture2D(graphicsDevice, size, size);
            Color[] data = new Color[size * size];

            Vector2 center = new Vector2(size / 2, size / 2);
            float radius = size / 2;

            for (int y = 0; y < size; y++)
            {
                for (int x = 0; x < size; x++)
                {
                    Vector2 pos = new Vector2(x, y);
                    float distance = Vector2.Distance(pos, center);

                    // Smooth falloff
                    float falloff = 1.0f - MathHelper.Clamp(distance / radius, 0f, 1f);
                    falloff = falloff * falloff; // Quadratic falloff for smoothness

                    data[y * size + x] = new Color(falloff, falloff, falloff, 1f);
                }
            }

            radialGradientTexture.SetData(data);

            // Create a 1x1 white texture for other uses
            lightTexture = new Texture2D(graphicsDevice, 1, 1);
            lightTexture.SetData(new[] { Color.White });
        }

        /// <summary>
        /// Adds a light source at the specified position.
        /// </summary>
        /// <param name="position">World position.</param>
        /// <param name="color">Light color.</param>
        /// <param name="radius">Effect radius.</param>
        /// <param name="intensity">Light intensity.</param>
        /// <returns>The created light.</returns>
        public PointLight AddLight(Vector2 position, Color color, float radius, float intensity)
        {
            var light = new PointLight(position, color, radius, intensity);
            lights.Add(light);
            return light;
        }

        /// <summary>
        /// Removes a light source.
        /// </summary>
        /// <param name="light">The light to remove.</param>
        public void RemoveLight(PointLight light)
        {
            lights.Remove(light);
        }

        /// <summary>
        /// Begins the scene rendering pass. Call this before drawing game objects.
        /// </summary>
        public void BeginScene()
        {
            graphicsDevice.SetRenderTarget(sceneTarget);
            graphicsDevice.Clear(Color.Black);
        }

        /// <summary>
        /// Ends the scene rendering pass.
        /// </summary>
        public void EndScene()
        {
            graphicsDevice.SetRenderTarget(null);
        }

        /// <summary>
        /// Renders the light map with all active lights.
        /// </summary>
        /// <param name="cameraPosition">Current camera position for world-to-screen conversion.</param>
        /// <param name="transformScale">Global transformation scale.</param>
        public void RenderLightMap(Vector2 cameraPosition, float transformScale)
        {
            // Set render target for light map
            graphicsDevice.SetRenderTarget(lightMapTarget);
            graphicsDevice.Clear(Color.Black); // Start with full darkness

            // Enable additive blending for lights
            spriteBatch.Begin(SpriteSortMode.Deferred, BlendState.Additive);

            foreach (var light in lights)
            {
                if (!light.IsActive)
                    continue;

                // Convert world position to screen position
                Vector2 screenPos = (light.Position - cameraPosition) / transformScale;

                // Calculate draw position (centered on light position)
                float scaledRadius = light.Radius / transformScale;
                Vector2 drawPos = screenPos - new Vector2(scaledRadius);

                // Calculate source rectangle for the gradient texture
                Rectangle sourceRect = new Rectangle(0, 0, radialGradientTexture.Width, radialGradientTexture.Height);

                // Draw the light with color and intensity
                Color lightDrawColor = new Color(
                    light.Color.R / 255f * light.Intensity,
                    light.Color.G / 255f * light.Intensity,
                    light.Color.B / 255f * light.Intensity,
                    1f
                );

                spriteBatch.Draw(
                    radialGradientTexture,
                    new Rectangle(
                        (int)drawPos.X,
                        (int)drawPos.Y,
                        (int)(scaledRadius * 2),
                        (int)(scaledRadius * 2)
                    ),
                    sourceRect,
                    lightDrawColor
                );
            }

            spriteBatch.End();

            graphicsDevice.SetRenderTarget(null);
        }

        /// <summary>
        /// Applies the lighting effect to the scene and renders the final output.
        /// </summary>
        /// <param name="gameTime">Game time for animated effects.</param>
        public void ApplyLighting(GameTime gameTime)
        {
            time += (float)gameTime.ElapsedGameTime.TotalSeconds;

            // Set shader parameters (must match Lighting.fx names exactly)
            lightingEffect.Parameters["ScreenTexture"]?.SetValue(sceneTarget);
            lightingEffect.Parameters["LightMap"]?.SetValue(lightMapTarget);
            lightingEffect.Parameters["AmbientDarkness"]?.SetValue(ambientDarkness);
            lightingEffect.Parameters["LightIntensityMul"]?.SetValue(lightIntensity);
            lightingEffect.Parameters["FogColor"]?.SetValue(fogColor.ToVector3());
            lightingEffect.Parameters["VignetteStrength"]?.SetValue(vignetteStrength);
            lightingEffect.Parameters["Time"]?.SetValue(time);

            // Select the Lighting technique and apply
            lightingEffect.CurrentTechnique = lightingEffect.Techniques["Lighting"];

            spriteBatch.Begin(SpriteSortMode.Deferred, BlendState.Opaque, null, null, null, lightingEffect);
            spriteBatch.Draw(sceneTarget, Vector2.Zero, Color.White);
            spriteBatch.End();
        }

        /// <summary>
        /// Disposes of resources used by the LightingManager.
        /// </summary>
        public void Dispose()
        {
            sceneTarget?.Dispose();
            lightMapTarget?.Dispose();
            radialGradientTexture?.Dispose();
            lightTexture?.Dispose();
            spriteBatch?.Dispose();
            lightingEffect?.Dispose();
        }
    }
}
