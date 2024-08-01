package interfaces

import "fmt"

// interface is a type that is used to specify a set of one or more method signatures.
// An interface is similar to a class in OOP.
// An interface is a reference type.
// An interface is a collection of abstract methods.
// An interface is a contract that specifies the behavior or functionality of an object.

// Syntax:
/*
type interfaceName interface {
	methodName() returnType
	methodName() returnType
	methodName() returnType
	....
}
*/

type AudioPlayer interface {
	Play()
	Pause()
	Stop()
}

type LocalFilePlayer struct {
	filePath string
}

func (localFilePlayer LocalFilePlayer) Play() {
	fmt.Println("Playing audio file from local file system")
}

func (LocalFilePlayer LocalFilePlayer) Pause() {
	fmt.Println("Pausing audio file from local file system")
}

func (LocalFilePlayer LocalFilePlayer) Stop() {
	fmt.Println("Stopping audio file from local file system")
}

type RemoteStreamPlayer struct {
	streamURL string
}

func (remotePlayer RemoteStreamPlayer) Play() {
	fmt.Println("Playing audio file from local file system")
}

func (remotePlayer RemoteStreamPlayer) Pause() {
	fmt.Println("Pausing audio file from local file system")
}

func (remotePlayer RemoteStreamPlayer) Stop() {
	fmt.Println("Stopping audio file from local file system")
}

func testInterface() {
	fmt.Println("Learning about interfaces.")

	// Create a local file player
	localPlayer := &LocalFilePlayer{
		filePath: "/path/to/local/file.mp3",
	}

	// Create a remote stream player
	remotePlayer := &RemoteStreamPlayer{
		streamURL: "http://example.com/stream.mp3",
	}

	// Play audio using the local file player
	playAudio(localPlayer)

	// Pause audio using the remote stream player
	pauseAudio(remotePlayer)

	// Stop audio using the local file player
	stopAudio(localPlayer)

}

func playAudio(audioPlayer AudioPlayer) {
	audioPlayer.Play()
}

func pauseAudio(audioPlayer AudioPlayer) {
	audioPlayer.Pause()
}

func stopAudio(audioPlayer AudioPlayer) {
	audioPlayer.Stop()
}

// type AuthMethod interface {
// 	login() bool
// }

// type EmailPasswordLoginMethod struct {
// 	email    string
// 	password string
// }

// type GoogleLoginMethod struct {
// 	userId        int
// 	client_id     string
// 	client_secret string
// }

// func (emailPasswordLogin EmailPasswordLoginMethod) login() bool {
// 	fmt.Println("Authenticating user with email and password")
// 	return true
// }

// func (googleLogin GoogleLoginMethod) login() bool {
// 	fmt.Println("Authenticating user with google")
// 	return true
// }

// func login(authMethod AuthMethod) bool {
// 	return authMethod.login()
// }
