# gRPC .NET client with Unity

This example folder implements a
[.NET gRPC](https://github.com/grpc/grpc-dotnet) client with Unity.

## Usage

In a terminal, go to `gnonative/goserver` and run the gRPC server:

```bash
go run . tcp -addr 127.0.0.1:7042
```

In Unity Hab, import this folder.

In Unity, select the `SampleScene` in `Assets/Scenes`. Play the project, click
on the button and look at the result in the text object.

## Structure

`Assets/Scenes/SampleScene.unity` is the Unity scene we use for the demo. We
implemented a button to start an action, and a text object to print the result.

`Assets/GnoClient.cs` is the script where lives the gRPC client. The `Start()`
method initializes the client on a fixed local port (`7042`). When you click on
the button object on the scene, the `OnHello` method calls the gRPC `Hello` call
and print the result on the text object.

`Assets/Gno` contains a copy of the generated GnoNative API files and gRPC C#
stubs (`api/gen/csharp`).

## Nuget package management

The gRPC .NET client is avalaible through the Nuget ecosystem. Unity doesn't
support it natively, so we had to add a module
([NuGetForUnity](https://github.com/GlitchEnzo/NuGetForUnity)) to be able to
install it easily. We installed the `Grpc.Net.Client` family packages.
