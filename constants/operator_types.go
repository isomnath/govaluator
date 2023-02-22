package constants

/*
All operator types
*/
const (
	/*
		Time value comparison operator names
	*/
	Before = "BEFORE"
	After  = "AFTER"
	During = "DURING"
	/*
		Value comparison operator names
	*/
	Equals                   = "EQUALS"
	NotEquals                = "NOT_EQUALS"
	LessThan                 = "LESS_THAN"
	LessThanOrEqualTo        = "LESS_THAN_OR_EQUAL_TO"
	LessThanAny              = "LESS_THAN_ANY"
	LessThanOrEqualToAny     = "LESS_THAN_OR_EQUAL_TO_ANY"
	LessThanNone             = "LESS_THAN_NONE"
	LessThanOrEqualToNone    = "LESS_THAN_OR_EQUAL_TO_NONE"
	LessThanAll              = "LESS_THAN_ALL"
	LessThanOrEqualToAll     = "LESS_THAN_OR_EQUAL_TO_ALL"
	GreaterThan              = "GREATER_THAN"
	GreaterThanOrEqualTo     = "GREATER_THAN_OR_EQUAL_TO"
	GreaterThanAny           = "GREATER_THAN_ANY"
	GreaterThanOrEqualToAny  = "GREATER_THAN_OR_EQUAL_TO_ANY"
	GreaterThanNone          = "GREATER_THAN_NONE"
	GreaterThanOrEqualToNone = "GREATER_THAN_OR_EQUAL_TO_NONE"
	GreaterThanAll           = "GREATER_THAN_ALL"
	GreaterThanOrEqualToAll  = "GREATER_THAN_OR_EQUAL_TO_ALL"
	Between                  = "BETWEEN"
	NotBetween               = "NOT_BETWEEN"
	AnyOf                    = "ANY_OF"
	NoneOf                   = "NONE_OF"
	RegexMatch               = "REGEX_MATCH"
	RegexNotMatch            = "REGEX_NOT_MATCH"
	/*
		Slice comparison operator names
	*/
	UnorderedEquals          = "UNORDERED_EQUALS"
	UnorderedNotEquals       = "UNORDERED_NOT_EQUALS"
	OrderedEquals            = "ORDERED_EQUALS"
	OrderedNotEquals         = "ORDERED_NOT_EQUALS"
	AnyLessThan              = "ANY_LESS_THAN"
	AnyLessThanOrEqualTo     = "ANY_LESS_THAN_OR_EQUAL_TO"
	NoneLessThan             = "NONE_LESS_THAN"
	NoneLessThanOrEqualTo    = "NONE_LESS_THAN_OR_EQUAL_TO"
	AllLessThan              = "ALL_LESS_THAN"
	AllLessThanOrEqualTo     = "ALL_LESS_THAN_OR_EQUAL_TO"
	AnyGreaterThan           = "ANY_GREATER_THAN"
	AnyGreaterThanOrEqualTo  = "ANY_GREATER_THAN_OR_EQUAL_TO"
	NoneGreaterThan          = "NONE_GREATER_THAN"
	NoneGreaterThanOrEqualTo = "NONE_GREATER_THAN_OR_EQUAL_TO"
	AllGreaterThan           = "ALL_GREATER_THAN"
	AllGreaterThanOrEqualTo  = "ALL_GREATER_THAN_OR_EQUAL_TO"
	ReversedBetween          = "REVERSED_BETWEEN"
	ReversedNotBetween       = "REVERSED_NOT_BETWEEN"
	SupersetOf               = "SUPERSET_OF"
	NotSupersetOf            = "NOT_SUPERSET_OF"
	SubsetOf                 = "SUBSET_OF"
	NotSubsetOf              = "NOT_SUBSET_OF"
	Intersection             = "INTERSECTION"
	NotIntersection          = "NOT_INTERSECTION"
	ReversedAnyOf            = "REVERSED_ANY_OF"
	ReversedNoneOf           = "REVERSED_NONE_OF"
	SizeEquals               = "SIZE_EQUALS"
	SizeNotEquals            = "SIZE_NOT_EQUALS"
	SizeLessThan             = "SIZE_LESS_THAN"
	SizeLessThanOrEqualTo    = "SIZE_LESS_THAN_OR_EQUAL_TO"
	SizeGreaterThan          = "SIZE_GREATER_THAN"
	SizeGreaterThanOrEqualTo = "SIZE_GREATER_THAN_OR_EQUAL_TO"
	AnyRegexMatch            = "ANY_REGEX_MATCH"
	NoneRegexMatch           = "NONE_REGEX_MATCH"
	AllRegexMatch            = "ALL_REGEX_MATCH"
)
