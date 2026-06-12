// Animal Well-style lighting shader
// Creates atmospheric darkness with point light sources

#if OPENGL
#define VS_SHADERMODEL vs_3_0
#define PS_SHADERMODEL ps_3_0
#else
#define VS_SHADERMODEL vs_4_0_level_9_1
#define PS_SHADERMODEL ps_4_0_level_9_1
#endif

// ---- Textures ----

texture ScreenTexture;
sampler ScreenSampler = sampler_state
{
    Texture = <ScreenTexture>;
    MinFilter = Linear;
    MagFilter = Linear;
    MipFilter = Linear;
    AddressU = Clamp;
    AddressV = Clamp;
};

texture LightMap;
sampler LightSampler = sampler_state
{
    Texture = <LightMap>;
    MinFilter = Linear;
    MagFilter = Linear;
    MipFilter = Linear;
    AddressU = Clamp;
    AddressV = Clamp;
};

// ---- Parameters ----

float AmbientDarkness;
float LightIntensityMul;
float3 FogColor;
float VignetteStrength;
float Time;

// ---- Structures ----

struct VS_INPUT
{
    float4 Position : POSITION0;
    float4 Color    : COLOR0;
    float2 TexCoord : TEXCOORD0;
};

struct VS_OUTPUT
{
    float4 Position : POSITION0;
    float4 Color    : COLOR0;
    float2 TexCoord : TEXCOORD0;
};

// ---- Vertex Shader (passthrough for SpriteBatch) ----

VS_OUTPUT SpriteVS(VS_INPUT input)
{
    VS_OUTPUT output;
    output.Position = input.Position;
    output.Color    = input.Color;
    output.TexCoord = input.TexCoord;
    return output;
}

// ---- Noise functions for atmospheric fog ----

float hash(float2 p)
{
    return frac(sin(dot(p, float2(127.1, 311.7))) * 43758.5453);
}

float noise(float2 p)
{
    float2 i = floor(p);
    float2 f = frac(p);
    float2 u = f * f * (3.0 - 2.0 * f);

    return lerp(
        lerp(hash(i),                hash(i + float2(1, 0)), u.x),
        lerp(hash(i + float2(0, 1)), hash(i + float2(1, 1)), u.x),
        u.y
    );
}

float fbm(float2 p)
{
    float v = 0.0;
    float a = 0.5;
    for (int i = 0; i < 4; i++)
    {
        v += a * noise(p);
        p  *= 2.0;
        a  *= 0.5;
    }
    return v;
}

// ---- Main lighting pixel shader ----

float4 LightingPS(VS_OUTPUT input) : COLOR0
{
    float4 scene = tex2D(ScreenSampler, input.TexCoord);
    float4 light = tex2D(LightSampler, input.TexCoord);

    // Base light contribution from light map
    float luma = light.r * LightIntensityMul;

    // Atmospheric noise (subtle moving fog)
    float2 nc = input.TexCoord * 6.0 + Time * 0.03;
    float fog = fbm(nc) * 0.12;

    float finalLight = saturate(luma + fog * 0.25);

    // Darkness = ambient * (1 - light)
    float dark = AmbientDarkness * (1.0 - finalLight);

    // Vignette
    float2 vc = input.TexCoord - 0.5;
    float vig = saturate(1.0 - dot(vc, vc) * VignetteStrength);

    // Fog tint in dark areas
    float3 fogged = lerp(scene.rgb, FogColor, dark * 0.55);

    // Final composite
    float3 color = fogged * (1.0 - dark * 0.75) * vig;

    // Light bleed glow
    color += light.rgb * 0.25 * saturate(1.0 - dark);

    return float4(color, scene.a);
}

// ---- Light map passthrough ----

float4 LightMapPS(VS_OUTPUT input) : COLOR0
{
    return tex2D(ScreenSampler, input.TexCoord);
}

// ---- Techniques ----

technique Lighting
{
    pass P0
    {
        VertexShader = compile VS_SHADERMODEL SpriteVS();
        PixelShader  = compile PS_SHADERMODEL LightingPS();
    }
}

technique LightMap
{
    pass P0
    {
        VertexShader = compile VS_SHADERMODEL SpriteVS();
        PixelShader  = compile PS_SHADERMODEL LightMapPS();
    }
}
