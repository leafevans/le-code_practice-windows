# Animal Well Light

A MonoGame 2D side-scrolling platformer with Animal Well-style lighting effects.

## Features

- **Animal Well-style Lighting System**
  - Dynamic point lights with soft falloff
  - Ambient darkness with atmospheric fog
  - Player light that pulses and follows the character
  - Gem lights with cyan/blue glow
  - Exit light with green glow
  - Vignette effect for cinematic feel
  - Atmospheric noise/fog for depth

- **2D Platformer Gameplay**
  - Side-scrolling level design
  - Player movement and jumping
  - Gem collection system
  - Enemy AI and collision
  - Multiple levels with increasing difficulty
  - Score tracking and leaderboard

## Lighting System

The lighting system uses multi-pass rendering:

1. **Scene Pass**: Renders all game objects to an off-screen target
2. **Light Map Pass**: Renders point lights with additive blending
3. **Composition Pass**: Combines scene and light map with atmospheric effects

### Key Components

- `LightingManager`: Manages render targets, shaders, and light sources
- `PointLight`: Represents individual light sources with position, color, radius, and intensity
- `Lighting.fx`: HLSL shader for atmospheric lighting effects

### Light Configuration

```csharp
// Ambient darkness (0.0 = fully lit, 1.0 = fully dark)
lightingManager.AmbientDarkness = 0.88f;

// Light intensity multiplier
lightingManager.LightIntensityMultiplier = 1.3f;

// Fog color (dark purple/blue like Animal Well)
lightingManager.FogColor = new Color(12, 8, 28);

// Vignette strength
lightingManager.VignetteStrength = 1.8f;
```

### Adding Lights

```csharp
// Add a point light
var light = lightingManager.AddLight(
    position,           // World position
    Color.White,        // Light color
    150f,              // Radius
    1.0f               // Intensity
);

// Update light position
light.Position = newPosition;

// Remove light
lightingManager.RemoveLight(light);
```

## Project Structure

```
animal-well-light/
├── AnimalWellLight.Core/           # Core game logic
│   ├── Content/                    # Game assets
│   │   ├── Effects/               # Shaders
│   │   │   └── Lighting.fx       # Animal Well lighting shader
│   │   ├── Backgrounds/          # Parallax background layers
│   │   ├── Sprites/              # Character and object sprites
│   │   ├── Tiles/                # Level tile textures
│   │   ├── Sounds/               # Audio files
│   │   └── Levels/               # Level data files
│   ├── Effects/                   # Effect managers
│   │   ├── LightingManager.cs    # Main lighting system
│   │   ├── Particle.cs           # Particle system
│   │   └── ParticleManager.cs    # Particle management
│   ├── Game/                      # Game entities
│   │   ├── Level.cs              # Level management (with lighting)
│   │   ├── Player.cs             # Player character
│   │   ├── Enemy.cs              # Enemy AI
│   │   └── Gem.cs                # Collectible gems
│   ├── Screens/                   # Game screens
│   │   ├── GameplayScreen.cs     # Main gameplay
│   │   ├── MainMenuScreen.cs     # Title screen
│   │   └── PauseScreen.cs        # Pause menu
│   └── AnimalWellLightGame.cs    # Main game class
├── AnimalWellLight.DesktopGL/     # Desktop platform project
├── AnimalWellLight.Android/       # Android platform project
├── AnimalWellLight.iOS/           # iOS platform project
└── AnimalWellLight.sln            # Solution file
```

## Building and Running

### Prerequisites

- .NET 9.0 SDK
- MonoGame templates (`dotnet new install MonoGame.Templates.CSharp`)

### Build

```bash
cd animal-well-light
dotnet build AnimalWellLight.DesktopGL/AnimalWellLight.DesktopGL.csproj
```

### Run

```bash
dotnet run --project AnimalWellLight.DesktopGL/AnimalWellLight.DesktopGL.csproj
```

## Controls

- **Arrow Keys / WASD**: Move left/right, jump
- **Space**: Jump (alternative)
- **Escape / P**: Pause game
- **Enter**: Confirm/continue

## Customization

### Adjusting Lighting

Edit the `InitializeLighting()` method in `Level.cs` to change:

- Ambient darkness level
- Light colors and intensities
- Fog color
- Vignette strength
- Number and placement of atmospheric lights

### Adding New Light Sources

1. Create a `PointLight` instance in `Level.cs`
2. Update its position in `UpdateLighting()`
3. The lighting system automatically handles rendering

## Technical Details

### Shader Features

- **Radial gradient falloff**: Smooth light attenuation
- **Atmospheric noise**: Subtle fog/mist effect using FBM noise
- **Vignette**: Cinematic darkening at screen edges
- **Light bleeding**: Soft glow effect around lit areas
- **Color mixing**: Lights blend naturally with scene colors

### Performance Considerations

- Lights are culled based on screen visibility
- Render targets are optimized for the target resolution
- Shader uses efficient noise functions for atmospheric effects

## Credits

- MonoGame framework
- Animal Well for lighting inspiration
- 2D StartKit template base

## License

This project is for educational purposes. Feel free to use and modify.
