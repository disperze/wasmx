# Setup 
Setting up WasmX is pretty straightforward. It requires three things to be done:
1. Install wasmx.
1. Initialize the configuration. 
2. Start the parser. 

## Installing WasmX
In order to install WasmX you are required to have [Go 1.16+](https://golang.org/dl/) installed on your machine. Once you have it, the first thing to do is to clone the GitHub repository. To do this you can run

```shell
$ git clone https://github.com/disperze/wasmx.git
```

Then, you need to install the binary. To do this, run 

```shell
$ make install
```

This will put the `wasmx` binary inside your `$GOPATH/bin` folder. You should now be able to run `wasmx` to make sure it's installed: 

```shell
$ wasmx version
```

## Initializing the configuration
In order to correctly parse and store the data based on your requirements, WasmX allows you to customize its behavior via a TOML file called `config.toml`. In order to create the first instance of the `config.toml` file you can run

```shell
$ wasmx init
```

This will create such file inside the `~/.wasmx` folder.  
Note that if you want to change the folder used by WasmX you can do this using the `--home` flag: 

```shell
$ wasmx init --home /path/to/my/folder
```

Once the file is created, you are required to edit it and change the different values. To do this you can run 

```shell
$ nano ~/.wasmx/config.toml
```

For a better understanding of what each section and field refers to, please read the [config reference](https://github.com/forbole/juno/blob/v2/cosmos-stargate/.docs/config.md). 

## Running WasmX 
Once the configuration file has been setup, you can run WasmX using the following command: 

```shell
$ .wasmx parse
```

If you are using a custom folder for the configuration file, please specify it using the `--home` flag: 


```shell
$ wasmx parse --home /path/to/my/config/folder
```

We highly suggest you running WasmX as a system service so that it can be restarted automatically in the case it stops. To do this you can run: 

```shell
$ sudo tee /etc/systemd/system/wasmx.service > /dev/null <<EOF
[Unit]
Description=WasmX parser
After=network-online.target

[Service]
User=$USER
ExecStart=$GOPATH/bin/wasmx parse
Restart=always
RestartSec=3
LimitNOFILE=4096

[Install]
WantedBy=multi-user.target
EOF
```

Then you need to enable and start the service:

```shell
$ sudo systemctl enable wasmx
$ sudo systemctl start wasmx
```