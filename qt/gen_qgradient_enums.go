package qt

type QGradient__Type int64

const (
	QGradient__LinearGradient QGradient__Type = 0
	QGradient__RadialGradient QGradient__Type = 1
	QGradient__ConicalGradient QGradient__Type = 2
	QGradient__NoGradient QGradient__Type = 3
)

type QGradient__Spread int64

const (
	QGradient__PadSpread QGradient__Spread = 0
	QGradient__ReflectSpread QGradient__Spread = 1
	QGradient__RepeatSpread QGradient__Spread = 2
)

type QGradient__CoordinateMode int64

const (
	QGradient__LogicalMode QGradient__CoordinateMode = 0
	QGradient__StretchToDeviceMode QGradient__CoordinateMode = 1
	QGradient__ObjectBoundingMode QGradient__CoordinateMode = 2
	QGradient__ObjectMode QGradient__CoordinateMode = 3
)

type QGradient__InterpolationMode int64

const (
	QGradient__ColorInterpolation QGradient__InterpolationMode = 0
	QGradient__ComponentInterpolation QGradient__InterpolationMode = 1
)

type QGradient__Preset int64

const (
	QGradient__WarmFlame QGradient__Preset = 1
	QGradient__NightFade QGradient__Preset = 2
	QGradient__SpringWarmth QGradient__Preset = 3
	QGradient__JuicyPeach QGradient__Preset = 4
	QGradient__YoungPassion QGradient__Preset = 5
	QGradient__LadyLips QGradient__Preset = 6
	QGradient__SunnyMorning QGradient__Preset = 7
	QGradient__RainyAshville QGradient__Preset = 8
	QGradient__FrozenDreams QGradient__Preset = 9
	QGradient__WinterNeva QGradient__Preset = 10
	QGradient__DustyGrass QGradient__Preset = 11
	QGradient__TemptingAzure QGradient__Preset = 12
	QGradient__HeavyRain QGradient__Preset = 13
	QGradient__AmyCrisp QGradient__Preset = 14
	QGradient__MeanFruit QGradient__Preset = 15
	QGradient__DeepBlue QGradient__Preset = 16
	QGradient__RipeMalinka QGradient__Preset = 17
	QGradient__CloudyKnoxville QGradient__Preset = 18
	QGradient__MalibuBeach QGradient__Preset = 19
	QGradient__NewLife QGradient__Preset = 20
	QGradient__TrueSunset QGradient__Preset = 21
	QGradient__MorpheusDen QGradient__Preset = 22
	QGradient__RareWind QGradient__Preset = 23
	QGradient__NearMoon QGradient__Preset = 24
	QGradient__WildApple QGradient__Preset = 25
	QGradient__SaintPetersburg QGradient__Preset = 26
	QGradient__PlumPlate QGradient__Preset = 28
	QGradient__EverlastingSky QGradient__Preset = 29
	QGradient__HappyFisher QGradient__Preset = 30
	QGradient__Blessing QGradient__Preset = 31
	QGradient__SharpeyeEagle QGradient__Preset = 32
	QGradient__LadogaBottom QGradient__Preset = 33
	QGradient__LemonGate QGradient__Preset = 34
	QGradient__ItmeoBranding QGradient__Preset = 35
	QGradient__ZeusMiracle QGradient__Preset = 36
	QGradient__OldHat QGradient__Preset = 37
	QGradient__StarWine QGradient__Preset = 38
	QGradient__HappyAcid QGradient__Preset = 41
	QGradient__AwesomePine QGradient__Preset = 42
	QGradient__NewYork QGradient__Preset = 43
	QGradient__ShyRainbow QGradient__Preset = 44
	QGradient__MixedHopes QGradient__Preset = 46
	QGradient__FlyHigh QGradient__Preset = 47
	QGradient__StrongBliss QGradient__Preset = 48
	QGradient__FreshMilk QGradient__Preset = 49
	QGradient__SnowAgain QGradient__Preset = 50
	QGradient__FebruaryInk QGradient__Preset = 51
	QGradient__KindSteel QGradient__Preset = 52
	QGradient__SoftGrass QGradient__Preset = 53
	QGradient__GrownEarly QGradient__Preset = 54
	QGradient__SharpBlues QGradient__Preset = 55
	QGradient__ShadyWater QGradient__Preset = 56
	QGradient__DirtyBeauty QGradient__Preset = 57
	QGradient__GreatWhale QGradient__Preset = 58
	QGradient__TeenNotebook QGradient__Preset = 59
	QGradient__PoliteRumors QGradient__Preset = 60
	QGradient__SweetPeriod QGradient__Preset = 61
	QGradient__WideMatrix QGradient__Preset = 62
	QGradient__SoftCherish QGradient__Preset = 63
	QGradient__RedSalvation QGradient__Preset = 64
	QGradient__BurningSpring QGradient__Preset = 65
	QGradient__NightParty QGradient__Preset = 66
	QGradient__SkyGlider QGradient__Preset = 67
	QGradient__HeavenPeach QGradient__Preset = 68
	QGradient__PurpleDivision QGradient__Preset = 69
	QGradient__AquaSplash QGradient__Preset = 70
	QGradient__SpikyNaga QGradient__Preset = 72
	QGradient__LoveKiss QGradient__Preset = 73
	QGradient__CleanMirror QGradient__Preset = 75
	QGradient__PremiumDark QGradient__Preset = 76
	QGradient__ColdEvening QGradient__Preset = 77
	QGradient__CochitiLake QGradient__Preset = 78
	QGradient__SummerGames QGradient__Preset = 79
	QGradient__PassionateBed QGradient__Preset = 80
	QGradient__MountainRock QGradient__Preset = 81
	QGradient__DesertHump QGradient__Preset = 82
	QGradient__JungleDay QGradient__Preset = 83
	QGradient__PhoenixStart QGradient__Preset = 84
	QGradient__OctoberSilence QGradient__Preset = 85
	QGradient__FarawayRiver QGradient__Preset = 86
	QGradient__AlchemistLab QGradient__Preset = 87
	QGradient__OverSun QGradient__Preset = 88
	QGradient__PremiumWhite QGradient__Preset = 89
	QGradient__MarsParty QGradient__Preset = 90
	QGradient__EternalConstance QGradient__Preset = 91
	QGradient__JapanBlush QGradient__Preset = 92
	QGradient__SmilingRain QGradient__Preset = 93
	QGradient__CloudyApple QGradient__Preset = 94
	QGradient__BigMango QGradient__Preset = 95
	QGradient__HealthyWater QGradient__Preset = 96
	QGradient__AmourAmour QGradient__Preset = 97
	QGradient__RiskyConcrete QGradient__Preset = 98
	QGradient__StrongStick QGradient__Preset = 99
	QGradient__ViciousStance QGradient__Preset = 100
	QGradient__PaloAlto QGradient__Preset = 101
	QGradient__HappyMemories QGradient__Preset = 102
	QGradient__MidnightBloom QGradient__Preset = 103
	QGradient__Crystalline QGradient__Preset = 104
	QGradient__PartyBliss QGradient__Preset = 106
	QGradient__ConfidentCloud QGradient__Preset = 107
	QGradient__LeCocktail QGradient__Preset = 108
	QGradient__RiverCity QGradient__Preset = 109
	QGradient__FrozenBerry QGradient__Preset = 110
	QGradient__ChildCare QGradient__Preset = 112
	QGradient__FlyingLemon QGradient__Preset = 113
	QGradient__NewRetrowave QGradient__Preset = 114
	QGradient__HiddenJaguar QGradient__Preset = 115
	QGradient__AboveTheSky QGradient__Preset = 116
	QGradient__Nega QGradient__Preset = 117
	QGradient__DenseWater QGradient__Preset = 118
	QGradient__Seashore QGradient__Preset = 120
	QGradient__MarbleWall QGradient__Preset = 121
	QGradient__CheerfulCaramel QGradient__Preset = 122
	QGradient__NightSky QGradient__Preset = 123
	QGradient__MagicLake QGradient__Preset = 124
	QGradient__YoungGrass QGradient__Preset = 125
	QGradient__ColorfulPeach QGradient__Preset = 126
	QGradient__GentleCare QGradient__Preset = 127
	QGradient__PlumBath QGradient__Preset = 128
	QGradient__HappyUnicorn QGradient__Preset = 129
	QGradient__AfricanField QGradient__Preset = 131
	QGradient__SolidStone QGradient__Preset = 132
	QGradient__OrangeJuice QGradient__Preset = 133
	QGradient__GlassWater QGradient__Preset = 134
	QGradient__NorthMiracle QGradient__Preset = 136
	QGradient__FruitBlend QGradient__Preset = 137
	QGradient__MillenniumPine QGradient__Preset = 138
	QGradient__HighFlight QGradient__Preset = 139
	QGradient__MoleHall QGradient__Preset = 140
	QGradient__SpaceShift QGradient__Preset = 142
	QGradient__ForestInei QGradient__Preset = 143
	QGradient__RoyalGarden QGradient__Preset = 144
	QGradient__RichMetal QGradient__Preset = 145
	QGradient__JuicyCake QGradient__Preset = 146
	QGradient__SmartIndigo QGradient__Preset = 147
	QGradient__SandStrike QGradient__Preset = 148
	QGradient__NorseBeauty QGradient__Preset = 149
	QGradient__AquaGuidance QGradient__Preset = 150
	QGradient__SunVeggie QGradient__Preset = 151
	QGradient__SeaLord QGradient__Preset = 152
	QGradient__BlackSea QGradient__Preset = 153
	QGradient__GrassShampoo QGradient__Preset = 154
	QGradient__LandingAircraft QGradient__Preset = 155
	QGradient__WitchDance QGradient__Preset = 156
	QGradient__SleeplessNight QGradient__Preset = 157
	QGradient__AngelCare QGradient__Preset = 158
	QGradient__CrystalRiver QGradient__Preset = 159
	QGradient__SoftLipstick QGradient__Preset = 160
	QGradient__SaltMountain QGradient__Preset = 161
	QGradient__PerfectWhite QGradient__Preset = 162
	QGradient__FreshOasis QGradient__Preset = 163
	QGradient__StrictNovember QGradient__Preset = 164
	QGradient__MorningSalad QGradient__Preset = 165
	QGradient__DeepRelief QGradient__Preset = 166
	QGradient__SeaStrike QGradient__Preset = 167
	QGradient__NightCall QGradient__Preset = 168
	QGradient__SupremeSky QGradient__Preset = 169
	QGradient__LightBlue QGradient__Preset = 170
	QGradient__MindCrawl QGradient__Preset = 171
	QGradient__LilyMeadow QGradient__Preset = 172
	QGradient__SugarLollipop QGradient__Preset = 173
	QGradient__SweetDessert QGradient__Preset = 174
	QGradient__MagicRay QGradient__Preset = 175
	QGradient__TeenParty QGradient__Preset = 176
	QGradient__FrozenHeat QGradient__Preset = 177
	QGradient__GagarinView QGradient__Preset = 178
	QGradient__FabledSunset QGradient__Preset = 179
	QGradient__PerfectBlue QGradient__Preset = 180
	QGradient__NumPresets QGradient__Preset = 181
)

