WikidataQuiz
============

[![GoDoc](https://godoc.org/github.com/FraBle/WikidataQuiz?status.png)](https://godoc.org/github.com/FraBle/WikidataQuiz)

## What is WikidataQuiz?

WikidataQuiz is an interactive quiz game for exploring open knowledge sources from [Wikidata](http://www.wikidata.org/ "Wikidata").
It was developed during the [Youth Science Hack Day 2014](http://re-publica.de/news/call-projects-youth-science-hack-day "Youth Science Hack Day 2014") in Berlin in the context of the [re:publica](http://re-publica.de/ "re:publica") conference.

Two players control the game using a [Leap Motion Controller](https://www.leapmotion.com/ "Leap Motion Controller") which acts as buzzer and answer input.
A connected [Arduino](http://www.arduino.cc/ "Arduino") with a LED stripe gives ambient visual feedback.
The provided questions and answers are generated using the [Wikidata](https://www.wikidata.org/w/api.php "Wikidata") and [DBpedia](http://dbpedia.org/sparql "DBpedia") API.
There are currently 3 question categories available:

- FIFA World Cup winners
- Capitals
- Day of death of Nobel Prize winners

## Images

![WikidataQuiz1](http://i.imgur.com/rqBixOgl.jpg "WikidataQuiz1")

![WikidataQuiz2](http://i.imgur.com/BpvMcZrl.jpg "WikidataQuiz2")

![WikidataQuiz3](http://i.imgur.com/Z3DDhEXl.jpg "WikidataQuiz3")

## How to get it running?
You need:

- [Go](http://golang.org/ "Go")
- [Git](http://git-scm.com/ "Git")
- [Mercurial](http://mercurial.selenic.com/ "Mercurial")
- [Bazaar](http://bazaar.canonical.com/en/ "Bazaar")
- An [Arduino](http://www.arduino.cc/ "Arduino")
- A [Leap Motion Controller](https://www.leapmotion.com/ "Leap Motion Controller")

If everything is set up well, you just need to run `go get github.com/FraBle/WikidataQuiz` and adjust the config.yaml.

## Developers
- [Frank Blechschmidt](https://github.com/FraBle "Frank Blechschmidt")
- [Stephan Schultz](https://github.com/Steppschuh "Stephan Schultz")
- [Jan-Peter Heuzeroth](https://github.com/peterHeuz "Jan-Peter Heuzeroth")
- [Fredrik Teschke](https://github.com/ftes "Fredrik Teschke")
- [Tobias Rohloff](https://github.com/rhlff "Tobias Rohloff")

## Used APIs
- [Wikidata](https://www.wikidata.org/w/api.php "Wikidata")
- [Wikidata Query](http://wdq.wmflabs.org/ "Wikidata Query")
- [DBpedia SPARQL](http://dbpedia.org/sparql "DBpedia SPARQL")

## Used packages
- [github.com/gorilla/mux](http://www.gorillatoolkit.org/pkg/mux "github.com/gorilla/mux")
- [github.com/knakk/fenster/sparql](github.com/knakk/fenster/sparql "github.com/knakk/fenster/sparql")
- [github.com/distributed/sers](github.com/distributed/sers "github.com/distributed/sers")
- [bitbucket.org/kardianos/osext](http://www.bitbucket.org/kardianos/osext "bitbucket.org/kardianos/osext")
- [launchpad.net/goyaml](launchpad.net/goyaml "launchpad.net/goyaml")
