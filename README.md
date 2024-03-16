![](https://github.com/dome385/goGoldWatch/blob/main/ExampleGIF.gif)

# Gold & Silber Portfolio Manager

Der "Gold & Silber Portfolio Manager" ist eine Desktop-Anwendung, die mit dem Fyne-Toolkit für die Go-Programmiersprache entwickelt wurde. 
Diese Anwendung ermöglicht es Benutzern, Echtzeit-Preisdiagramme für Gold- und Silberhandel anzuzeigen und ihre Bestände an diesen Edelmetallen zu verwalten. 
Die Daten werden sicher in einer SQLite-Datenbank gespeichert, die eine einfache und effiziente Datenverwaltung ermöglicht.
# Hauptfunktionen:

- Echtzeit-Preisdiagramme: Nutzer können aktuelle Preisdiagramme für Gold und Silber einsehen, die ihnen helfen, Markttrends zu verfolgen und informierte Handelsentscheidungen zu treffen.
- Portfolioverwaltung: Benutzer können ihre Bestände an Gold und Silber eingeben und verwalten. Die Anwendung berechnet den Gesamtwert des Portfolios basierend auf den aktuellen Marktpreisen.
- Datenbankintegration: Alle Benutzerdaten, einschließlich der Portfolioinformationen, werden in einer SQLite-Datenbank gespeichert. Dies bietet eine schnelle und zuverlässige Möglichkeit, Daten zu speichern und abzurufen.
- Benutzerfreundliche Oberfläche: Die Anwendung bietet eine intuitive und ansprechende Benutzeroberfläche, die es auch unerfahrenen Benutzern ermöglicht, ihre Portfolios effektiv zu verwalten.
- Plattformübergreifende Kompatibilität: Dank Fyne können Benutzer die Anwendung auf verschiedenen Betriebssystemen wie Windows, macOS und Linux ohne zusätzliche Anpassungen ausführen.

## Technische Details:
- Programmiersprache: Go (Golang) mit dem Fyne-Toolkit, das eine einfache API für die Erstellung moderner GUI-Anwendungen bietet
- Datenbank: SQLite, eine leichte, dateibasierte Datenbank, ideal für Anwendungen mit geringem Ressourcenbedarf und einfacher Einrichtung
- Diagramm- und Grafikunterstützung: Obwohl Fyne zum Zeitpunkt des Wissensstands keine integrierte Chart- oder Graph-Widget-Unterstützung bietet, können Entwickler externe Go-Bibliotheken wie go-chart verwenden, um Diagramme zu rendern und in die Fyne-Anwendung zu integrieren
- Cross-Compilation: Fyne unterstützt die Cross-Compilation, sodass Entwickler die Anwendung für verschiedene Plattformen aus einer Codebasis heraus kompilieren können
- Datenbindung und Widgets: Fyne bietet Datenbindungsfunktionen und eine Vielzahl von Widgets, die für die Erstellung interaktiver Benutzeroberflächen erforderlich sind

## Entwicklung und Bereitstellung:
Die Entwicklung des "Gold & Silber Portfolio Managers" folgt den Best Practices für Go-Projekte, einschließlich der Verwendung von Modulen für die Abhängigkeitsverwaltung und der Einhaltung der Go-Codekonventionen. 
Die Anwendung kann mit dem Go-Compiler in ausführbare Dateien für die unterstützten Betriebssysteme kompiliert werden.
Die Bereitstellung der Anwendung ist dank der statischen Kompilierung von Go und der eigenständigen Natur von Fyne-Anwendungen unkompliziert.
