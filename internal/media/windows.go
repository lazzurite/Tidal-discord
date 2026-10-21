//go:build windows

package media

import (
	"fmt"
	"time"

	"github.com/go-ole/go-ole"
	"github.com/saltosystems/winrt-go"
	"github.com/saltosystems/winrt-go/windows/foundation"
	"github.com/saltosystems/winrt-go/windows/media/control"
)

func getManager() (*control.GlobalSystemMediaTransportControlsSessionManager, error) {
	op, err := control.GlobalSystemMediaTransportControlsSessionManagerRequestAsync()
	if err != nil {
		return nil, err
	}

	done := make(chan foundation.AsyncStatus, 1)

	handler := foundation.NewAsyncOperationCompletedHandler(
		ole.NewGUID(winrt.ParameterizedInstanceGUID(
			foundation.GUIDAsyncOperationCompletedHandler,
			"rc(Windows.Media.Control.GlobalSystemMediaTransportControlsSessionManager;{cace8eac-e86e-504a-ab31-5ff8ff1bce49})",
		)),
		func(
			_ *foundation.AsyncOperationCompletedHandler,
			_ *foundation.IAsyncOperation,
			status foundation.AsyncStatus,
		) {
			done <- status
		},
	)

	defer handler.Release()

	if err := op.SetCompleted(handler); err != nil {
		return nil, err
	}

	status := <-done

	if status != foundation.AsyncStatusCompleted {
		return nil, fmt.Errorf("RequestAsync failed with status %d", status)
	}

	result, err := op.GetResults()
	if err != nil {
		return nil, err
	}

	return (*control.GlobalSystemMediaTransportControlsSessionManager)(result), nil
}

type WindowsProvider struct{}

func (WindowsProvider) CurrentTrack() (*Track, error) {
	manager, err := getManager()
	if err != nil {
		return nil, err
	}

	session, err := manager.GetCurrentSession()
	if err != nil {
		return nil, err
	}

	playbackInfo, err := session.GetPlaybackInfo()
	if err != nil {
		return nil, err
	}

	playbackStatus, err := playbackInfo.GetPlaybackStatus()
	if err != nil {
		return nil, err
	}
	timeline, err := session.GetTimelineProperties()
	if err != nil {
		return nil, err
	}

	position, err := timeline.GetPosition()
	if err != nil {
		return nil, err
	}

	endTime, err := timeline.GetEndTime()
	if err != nil {
		return nil, err
	}
	duration := time.Duration(endTime.Duration) * 100 * time.Nanosecond
	positionDuration := time.Duration(position.Duration) * 100 * time.Nanosecond

	playing := playbackStatus == control.GlobalSystemMediaTransportControlsSessionPlaybackStatusPlaying
	propertiesOp, err := session.TryGetMediaPropertiesAsync()
	if err != nil {
		return nil, err
	}

	done := make(chan foundation.AsyncStatus, 1)
	propertiesHandler := foundation.NewAsyncOperationCompletedHandler(
		ole.NewGUID(winrt.ParameterizedInstanceGUID(
			foundation.GUIDAsyncOperationCompletedHandler,
			"rc(Windows.Media.Control.GlobalSystemMediaTransportControlsSessionMediaProperties;{68856cf6-adb4-54b2-ac16-05837907acb6})",
		)),
		func(
			_ *foundation.AsyncOperationCompletedHandler,
			_ *foundation.IAsyncOperation,
			status foundation.AsyncStatus,
		) {
			done <- status
		},
	)
	defer propertiesHandler.Release()

	if err := propertiesOp.SetCompleted(propertiesHandler); err != nil {
		return nil, err
	}

	status := <-done
	if status != foundation.AsyncStatusCompleted {
		return nil, fmt.Errorf("media properties operation failed with status %d", status)
	}

	result, err := propertiesOp.GetResults()
	if err != nil {
		return nil, err
	}

	properties := (*control.GlobalSystemMediaTransportControlsSessionMediaProperties)(result)

	title, err := properties.GetTitle()
	if err != nil {
		return nil, err
	}

	artist, err := properties.GetArtist()
	if err != nil {
		return nil, err
	}

	album, err := properties.GetAlbumTitle()
	if err != nil {
		return nil, err
	}

	return &Track{
		Title:    title,
		Artist:   artist,
		Album:    album,
		Duration: duration,
		Position: positionDuration,
		Playing:  playing,
	}, nil
}
