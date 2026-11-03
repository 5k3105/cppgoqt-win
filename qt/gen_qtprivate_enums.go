package qt

type QtPrivate__Ordering int64

const (
	QtPrivate__Ordering__Equal QtPrivate__Ordering = 0
	QtPrivate__Ordering__Equivalent QtPrivate__Ordering = 0
	QtPrivate__Ordering__Less QtPrivate__Ordering = -1
	QtPrivate__Ordering__Greater QtPrivate__Ordering = 1
)

type QtPrivate__Uncomparable int64

const (
	QtPrivate__Uncomparable__Unordered QtPrivate__Uncomparable = 2
)

type QtPrivate__LegacyUncomparable int64

const (
	QtPrivate__LegacyUncomparable__Unordered QtPrivate__LegacyUncomparable = -127
)

type QtPrivate__SynthesizedAccessFunction int64

const (
	QtPrivate__SynthesizedAccessFunction__IterableSize QtPrivate__SynthesizedAccessFunction = 0
	QtPrivate__SynthesizedAccessFunction__SequenceAt QtPrivate__SynthesizedAccessFunction = 1
)

