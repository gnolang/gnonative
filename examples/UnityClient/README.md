# gRPC .NET client with Unity

This README will guide you to set up a new Unity project and to configure a
[.NET gRPC](https://github.com/grpc/grpc-dotnet) client in C#.

## Create a new project

Open Unity Hub, click on `New Project`. In the new window, select 3D in the Core category. Set the project name and the wanted location in the right panel. Click on `Create Project`.

## Install dependencies

The gRPC .NET client is avalaible through the Nuget ecosystem. Unity doesn't support it natively, so we had to add a module
([NuGetForUnity](https://github.com/GlitchEnzo/NuGetForUnity)) to be able to install it easily. Please install NuGetForUnity by following the README on their [github page](https://github.com/GlitchEnzo/NuGetForUnity) and come back here.

Upon the installation complete, restart Unity to have the new NuGet menu. Select Nuget -> Manage NuGet Packages. Search and install these packages:
- Grpc.Net.Client
- Grpc.Net.Client.Web
- Google.Protobuf

## Copy API dependencies

We will copy the generated GnoNative API files and gRPC C# stubs in `Assets/Gno`.

In a terminal, write this command (remplace ${gnonative} by the location where is gnonative in your filesystem):
```bash
cp -r ${gnonative}/api/gen/csharp Assets/GnoNative
```

## Create a C# script

In the project panel, click on the `Assets` folder to select it. Then click on the `+` sign and click on `C# script`. Name the new file `Hello`.

Double click on it to open it in Visual Studio. Replace the content with:
```c#
using System;
using System.Net.Http;
using UnityEngine;
using TMPro;

using Grpc.Net.Client;
using Grpc.Net.Client.Web;
using Land.Gno.Gnonative.V1;

public class Hello : MonoBehaviour
{
    public TextMeshProUGUI Text;

    private GrpcChannel channel;
    private GnoNativeService.GnoNativeServiceClient client;

    // Start is called before the first frame update
    void Start()
    {
        var options = new GrpcChannelOptions();
        var handler = new GrpcWebHandler(GrpcWebMode.GrpcWeb, new HttpClientHandler());
        options.HttpHandler = handler;

        channel = GrpcChannel.ForAddress("http://localhost:26658", options);
        client = new GnoNativeService.GnoNativeServiceClient(channel);
    }

    private void OnDestroy()
    {
        channel.Dispose();
    }

    public void OnHello()
    {
        var reply = client.Hello(new HelloRequest { Name = "Gno" });
        Debug.Log(reply.Greeting.ToString());
        Text.text = $"[{DateTime.Now}] {reply.Greeting}";
    }
}

```
In the `Start()` function, we initialize the gRPC client with the gRPC Web protocol and use the default GnoNative port.
In the `OnHello()` function, we do the `Hello` gRPC call and set the result into the `Text` object. This function can be trigger by a button for example.

Save the modifications and go back to Unity.

We have to link the script with an object in the scene. So in the `Hierarchy` panel, click on the `+` sign and `Create Empty`. Name it `UI Script`. From the Project panel, drag and drop the Hello script to the UI Script object in the Hierarchy panel.

## Create UI objects

We want to create a button and a text objects on the scene to start the GnoNative call and print the result.

In the Hierarchy panel, click on the `+` sign and select `UI` and `Text - TextMeshPro`. In the `TMP Importer` windows, click on `Import TMP Essentials` and close it. Still in the `UI` section, add a `Button - TextMeshPro`. You can move the objects in the scene to make it clean.

Now we have to attach these two objects to the script we wrote.

In the `Hierarchy` panel, select the button. Scroll down on the `Inspector` (right panel) to find the `On Click` section and click on `+` to add a new action. Drag and drop the `UI Script` object to the `None (Object)` in this action, and replace `No Function` by `Hello` -> `OnHello ()`.

In the `Hierarchy` panel, click on `UI Script`. In the `Inspector` panel, you can see the `Hello (Script)` section with the empty `Text` field. Drag and drop the `Text` object (in the `Hierarchy`) to this `Text` field (in the `Inspector`).

## Play

In a terminal, go to `gnonative/goserver` and run the gRPC server:

```bash
go run . tcp
```
In Unity, click on the play button and click to the button on the scene. The text object should change.
