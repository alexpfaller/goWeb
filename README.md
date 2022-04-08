
![Logo](https://i.postimg.cc/MGCc7g5W/go-Web-thumb.png)


# goWeb

Is an easy to use and simple webserver, built with go. The structure and CLI are improving the productivity, since it guides the user through everything with ease.





## Installation

Installation manually with shell script.

```bash
git clone https://github.com/vimscape/goWeb
cd goWeb
sudo ./goWeb.sh
```
After that you can simply run the application from everywhere, with:
```
goWeb
```

## Documentation

After the installation process is completed, you can launch the webserver for the first time. For that, run the app with `goWeb` and insert `2` in the upcoming menu. Now you should be able to see a green message, which tells you that your server is up and running. To see the default page in your browser, insert `localhost:8090` in your browser-adressbar.

The default page is customizable and can be found in the `default/` directory, where you can change your current working directory to, with the following command:
```
cd default/
```
If you want to stop the running webserver, press `ctrl + c` to abort its execution.

By default, the port is set to 8090, you can set a static port to whatever you want, please be aware of already used ports. In case you don't really care about it and just want a different port than 8090, we recommend using the `auto port` function, to assign a random unused port to your webserver.

[Documentation](https://linktodocumentation)


## Features

- Configure webserver
    - Set sub-pages
    - Set ports
        - Port scanning
        - Auto port selecting
        - Static port assigning
    - Create redirects
- Create website templates

