Write.as
========
[![Build Status](https://travis-ci.org/writeas/cmd.svg)](https://travis-ci.org/writeas/cmd)

This is a simple HTTP-based interface for publishing text. Users POST whatever they want to publish using the `w` parameter. When the request finishes, a URL is returned — this is the publicly-accessible address to the post on the web.

## Try it
```
echo "Hello world, by $USER" | curl -F 'w=<-' http://cmd.write.as
```

## Run it yourself
```
Usage:
  cmd [options]

Options:
  --debug
       Enables garrulous debug logging.
  -o   
       Directory where text files will be stored.
  -s
       Directory where required static files exist (like the banner).
  -p
       Port to listen on.
```

The default configuration (without any flags) is essentially:

```
cmd -o /var/write -s ./static -p 8080
```

## How it works
The user's input is simply written to a flat file in a given directory. To provide web access, a web server (sold separately) serves all files in this directory as `plain/text`. That's it!

## License
This project is licensed under the MIT open source license.
