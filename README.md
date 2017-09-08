# Printmaps Webservice

Printmaps ist ein Webservice der die Erstellung großformatiger, OSM-basierter Karten in Druckqualität erlaubt.
Die Karte kann dabei um beliebige Benutzerelemente und Benutzerdaten ergänzt werden.

* großformatige Karten in Druckqualität
* verschiedene Kartenstile (osm-carto, schwarzplan+, ...)
* verschiedene Ausgabeformate (png, [pdf, svg])
* aktuelle OpenStreetMap-Kartendaten
* Kartendaten verfügbar für Europa
* benutzerdefinierte Zusatzelemente (Rahmen, Legende, Maßstabsbalken, ...)
* benutzerdefinierte Datenobjekte (gpx, kml, shape, csv, ...)

Printmaps kann genutzt werden
* als Webservice
* via CLI-Client

![](sample1.png)

## Status
Printmaps befindet sich aktuell im Status "Beta 1".

## Einschränkungen
Karten im PDF- oder SVG-Format werden im Moment nicht unterstützt. Die entsprechenden Programmteile sind nicht getestet.

## Wie anfangen?
* durcharbeiten der Webseite [Printmaps website](http://printmaps-osm.de:8080/)
* erzeugen einiger Karten mit dem Printmaps-CLI-Client
* sich vertraut machen mit den Printmaps-Eigenschaften und -möglichkeiten 
* aufsetzen einer (minimalen) eigenen Printmaps-Testumgebung (ohne DB)
* aufsetzen einer (vollumfänglichen) Printmaps-Umgebung (mit DB)

## Printmaps besteht aus folgenden Komponenten

### Server (printmaps_webservice, printmaps_buildservice, Nik4)
* printmaps_webservice
* printmaps_buildservice
* Nik4

Siehe README.md in den entsprechenden Verzeichnissen.

### Client (printmaps_client)
Siehe README.md im printmaps_client-Verzeichnis.

### Kartenstile (printmaps_styles)
* osm-carto
* raster10
* schwarzplan
* schwarzplan+

Siehe README.md im printmaps_styles-Verzeichnis.

### Icons (printmaps_icons)
Siehe README.md im printmaps_icons-Verzeichnis.

### Pattern (printmaps_pattern1, printmaps_pattern2)
Siehe README.md in den entsprechenden Verzeichnissen.

### Datenpflege (printmaps_updater, printmaps_purger)
* printmaps_updater
* printmaps_purger

Siehe README.md in den entsprechenden Verzeichnissen.

### Webseite (printmaps-osm.de)
Siehe README.md im printmaps-osm.de-Verzeichnis.

## Empfohlene (Server-) Verzeichnisstruktur

    printmaps
        printmaps
            orders
            maps
        printstyles
            markers
            osm-carto
            osm-carto-ele
            osm-carto-mono
            raster10
            schwarzplan
            schwarzplan+
        printdata
            updater
            purger
        Nik4

## Vorbedingungen der Serverumgebung
Folgende Programme und Tools müssen installiert sein:
* python
* mapnik
* mapnik-utils
* python-mapnik
* zip

Beispiel für Ubuntu:

    sudo apt install python 
    sudo apt install libmapnik3.0 libmapnik-dev mapnik-utils python-mapnik
    sudo apt install zip

## Test der Serverumgebung
Es empfiehlt sich, zunächst mit einer minimalen Serverumgebung (ohne Datenbank) zu starten:

    printmaps
        printmaps
            orders
            maps
        printstyles
            markers
            raster10
        Nik4

Folgende Konfiguration sind anzupassen:
* printmaps_webservice.yaml
* printmaps_buildservice.yaml

Desweiteren ist in der Kartenkonfiguration (map.yaml für den CLI-Client) die URL der Testumgebung einzutragen. 

---

to be done - english translation
