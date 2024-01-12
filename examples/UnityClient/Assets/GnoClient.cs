using System;
using System.Collections;
using System.Collections.Generic;
using System.Net.Http;
using UnityEngine;
using TMPro;
using Grpc.Core;
using Grpc.Net.Client;
using Grpc.Net.Client.Web;
using Land.Gno.Gnonative.V1;

public class GnoClient : MonoBehaviour
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

        channel = GrpcChannel.ForAddress("http://localhost:7042", options);
        client = new GnoNativeService.GnoNativeServiceClient(channel);
    }

    private void OnDestroy()
    {
        channel.Dispose();
    }

    // Update is called once per frame
    void Update()
    {
        
    }

    public void OnHello()
    {
        var reply = client.Hello(new HelloRequest { Name = "Gno" });
        Debug.Log(reply.Greeting.ToString());
        Text.text = $"[{DateTime.Now}] {reply.Greeting}";
    }
}
