package vlc

import "errors"

// Generic errors.
var (
	ErrInvalid = errors.New("the provided value is not valid")
)

// Module errors.
var (
	ErrModuleInitialize     = errors.New("could not initialize module")
	ErrModuleNotInitialized = errors.New("module not initialized")
	ErrUserInterfaceStart   = errors.New("could not start user interface")
)

// Player errors.
var (
	ErrPlayerCreate         = errors.New("could not create player")
	ErrPlayerNotInitialized = errors.New("player is not initialized")
	ErrPlayerPlay           = errors.New("cannot play the requested media")
	ErrPlayerSetVolume      = errors.New("could not set player volume")
	ErrPlayerSetEqualizer   = errors.New("could not set player equalizer")
)

// List player errors.
var (
	ErrListPlayerCreate         = errors.New("could not create list player")
	ErrListPlayerNotInitialized = errors.New("list player is not initialized")
)

// Media errors.
var (
	ErrMediaCreate             = errors.New("could not create media")
	ErrMediaNotFound           = errors.New("could not find media")
	ErrMediaNotInitialized     = errors.New("media is not initialized")
	ErrMediaListCreate         = errors.New("could not create media list")
	ErrMediaListNotFound       = errors.New("could not find media list")
	ErrMediaListNotInitialized = errors.New("media list is not initialized")
	ErrMediaListReadOnly       = errors.New("media list is read-only")
	ErrMediaListActionFailed   = errors.New("could not perform media list action")
	ErrMissingMediaStats       = errors.New("could not get media statistics")
	ErrInvalidMediaStats       = errors.New("invalid media statistics")
	ErrMissingMediaLocation    = errors.New("could not get media location")
	ErrMissingMediaDimensions  = errors.New("could not get media dimensions")
	ErrMediaMetaSave           = errors.New("could not save media metadata")
	ErrMediaNotParsed          = errors.New("media is not parsed")
)

// Media track errors.
var (
	ErrMediaTrackNotInitialized = errors.New("media track is not initialized")
	ErrMediaTrackNotFound       = errors.New("could not find media track")
	ErrInvalidMediaTrack        = errors.New("invalid media track")
)

// Event manager errors.
var (
	ErrMissingEventManager  = errors.New("could not get event manager instance")
	ErrInvalidEventCallback = errors.New("invalid event callback")
)

// Audio/Video errors.
var (
	ErrAudioOutputListMissing       = errors.New("could not get audio output list")
	ErrAudioOutputSet               = errors.New("could not set audio output")
	ErrAudioOutputDeviceListMissing = errors.New("could not get audio output device list")
	ErrFilterListMissing            = errors.New("could not get filter list")
	ErrStereoModeSet                = errors.New("could not set stereo mode")
	ErrVideoSnapshot                = errors.New("could not take video snapshot")
	ErrCursorPositionMissing        = errors.New("could not get cursor position")
)

// Equalizer errors.
var (
	ErrEqualizerCreate         = errors.New("could not create equalizer")
	ErrEqualizerNotInitialized = errors.New("equalizer is not initialized")
	ErrEqualizerAmpValueSet    = errors.New("could not set equalizer amplification value")
	ErrEventAttach             = errors.New("could not attach event")
)
