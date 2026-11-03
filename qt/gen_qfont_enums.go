package qt

type QFont__StyleHint int64

const (
	QFont__Helvetica QFont__StyleHint = 0
	QFont__SansSerif QFont__StyleHint = 0
	QFont__Times QFont__StyleHint = 1
	QFont__Serif QFont__StyleHint = 1
	QFont__Courier QFont__StyleHint = 2
	QFont__TypeWriter QFont__StyleHint = 2
	QFont__OldEnglish QFont__StyleHint = 3
	QFont__Decorative QFont__StyleHint = 3
	QFont__System QFont__StyleHint = 4
	QFont__AnyStyle QFont__StyleHint = 5
	QFont__Cursive QFont__StyleHint = 6
	QFont__Monospace QFont__StyleHint = 7
	QFont__Fantasy QFont__StyleHint = 8
)

type QFont__StyleStrategy int64

const (
	QFont__PreferDefault QFont__StyleStrategy = 1
	QFont__PreferBitmap QFont__StyleStrategy = 2
	QFont__PreferDevice QFont__StyleStrategy = 4
	QFont__PreferOutline QFont__StyleStrategy = 8
	QFont__ForceOutline QFont__StyleStrategy = 16
	QFont__PreferMatch QFont__StyleStrategy = 32
	QFont__PreferQuality QFont__StyleStrategy = 64
	QFont__PreferAntialias QFont__StyleStrategy = 128
	QFont__NoAntialias QFont__StyleStrategy = 256
	QFont__NoSubpixelAntialias QFont__StyleStrategy = 2048
	QFont__PreferNoShaping QFont__StyleStrategy = 4096
	QFont__ContextFontMerging QFont__StyleStrategy = 8192
	QFont__PreferTypoLineMetrics QFont__StyleStrategy = 16384
	QFont__NoFontMerging QFont__StyleStrategy = 32768
)

type QFont__HintingPreference int64

const (
	QFont__PreferDefaultHinting QFont__HintingPreference = 0
	QFont__PreferNoHinting QFont__HintingPreference = 1
	QFont__PreferVerticalHinting QFont__HintingPreference = 2
	QFont__PreferFullHinting QFont__HintingPreference = 3
)

type QFont__Weight int64

const (
	QFont__Thin QFont__Weight = 100
	QFont__ExtraLight QFont__Weight = 200
	QFont__Light QFont__Weight = 300
	QFont__Normal QFont__Weight = 400
	QFont__Medium QFont__Weight = 500
	QFont__DemiBold QFont__Weight = 600
	QFont__Bold QFont__Weight = 700
	QFont__ExtraBold QFont__Weight = 800
	QFont__Black QFont__Weight = 900
)

type QFont__Style int64

const (
	QFont__StyleNormal QFont__Style = 0
	QFont__StyleItalic QFont__Style = 1
	QFont__StyleOblique QFont__Style = 2
)

type QFont__Stretch int64

const (
	QFont__AnyStretch QFont__Stretch = 0
	QFont__UltraCondensed QFont__Stretch = 50
	QFont__ExtraCondensed QFont__Stretch = 62
	QFont__Condensed QFont__Stretch = 75
	QFont__SemiCondensed QFont__Stretch = 87
	QFont__Unstretched QFont__Stretch = 100
	QFont__SemiExpanded QFont__Stretch = 112
	QFont__Expanded QFont__Stretch = 125
	QFont__ExtraExpanded QFont__Stretch = 150
	QFont__UltraExpanded QFont__Stretch = 200
)

type QFont__Capitalization int64

const (
	QFont__MixedCase QFont__Capitalization = 0
	QFont__AllUppercase QFont__Capitalization = 1
	QFont__AllLowercase QFont__Capitalization = 2
	QFont__SmallCaps QFont__Capitalization = 3
	QFont__Capitalize QFont__Capitalization = 4
)

type QFont__SpacingType int64

const (
	QFont__PercentageSpacing QFont__SpacingType = 0
	QFont__AbsoluteSpacing QFont__SpacingType = 1
)

type QFont__ResolveProperties int64

const (
	QFont__NoPropertiesResolved QFont__ResolveProperties = 0
	QFont__FamilyResolved QFont__ResolveProperties = 1
	QFont__SizeResolved QFont__ResolveProperties = 2
	QFont__StyleHintResolved QFont__ResolveProperties = 4
	QFont__StyleStrategyResolved QFont__ResolveProperties = 8
	QFont__WeightResolved QFont__ResolveProperties = 16
	QFont__StyleResolved QFont__ResolveProperties = 32
	QFont__UnderlineResolved QFont__ResolveProperties = 64
	QFont__OverlineResolved QFont__ResolveProperties = 128
	QFont__StrikeOutResolved QFont__ResolveProperties = 256
	QFont__FixedPitchResolved QFont__ResolveProperties = 512
	QFont__StretchResolved QFont__ResolveProperties = 1024
	QFont__KerningResolved QFont__ResolveProperties = 2048
	QFont__CapitalizationResolved QFont__ResolveProperties = 4096
	QFont__LetterSpacingResolved QFont__ResolveProperties = 8192
	QFont__WordSpacingResolved QFont__ResolveProperties = 16384
	QFont__HintingPreferenceResolved QFont__ResolveProperties = 32768
	QFont__StyleNameResolved QFont__ResolveProperties = 65536
	QFont__FamiliesResolved QFont__ResolveProperties = 131072
	QFont__FeaturesResolved QFont__ResolveProperties = 262144
	QFont__VariableAxesResolved QFont__ResolveProperties = 524288
	QFont__AllPropertiesResolved QFont__ResolveProperties = 1048575
)

