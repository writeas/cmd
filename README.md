Write.as
========
[![Build Status](https://travis-ci.org/writeas/cmd.svg)](https://travis-ci.org/writeas/cmd) [![#writeas on freenode](https://img.shields.io/badge/freenode-%23writeas-blue.svg)](http://webchat.freenode.net/?channels=writeas) [![Public Slack discussion](http://slack.write.as/badge.svg)](http://slack.write.as/)

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
       Directory where text files will be stored. If not supplied, will try to
       use database for storage (see Environment Variables).
  -s
       Directory where required static files exist (like the banner).
  -p
       Port to listen on.

Environment Variables:
  WA_USER
       Database user.
  WA_PASSWORD
       Database password.
  WA_HOST
       Database host. Default: localhost
  WA_DB
       Database name.
```

The default configuration (without any flags) is essentially the following line. **You'll need to supply the `-o` flag or database env variables to store posts**.

```
cmd -s ./static -p 8080
```

## How it works
The user's input is simply written to a flat file in a given directory. To provide web access, a web server (sold separately) serves all files in this directory as `plain/text`. That's it!

## License
This project is licensed under the MIT open source license.
