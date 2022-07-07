# Zeitmessungen für verschiedene Algorithmen

## Branch Prediction

Im Package `branchprediction` gibt es eine Funktion `sumGreater()`, die aus einer Liste
von Zahlen die Summe aller Elemente berechnet, die größer als ein gegebener Wert sind.

Diese Funktion wird von zwei anderen Funktionen verwendet, die beide fast das gleiche
tun: Sie erzeugen beide eine Liste mit Zufallszahlen im Intervall `[0, max]` und
berechnen dann die Summe aller Elemente, die größer als `max` sind.
Der Unterschied ist, dass eine der Funktionen die Liste vorher sortiert und die andere
nicht.

Wir messen die Zeit und stellen fest, dass die sortierende Liste bei großen Listen
wesentlich schneller ist. Zumindest gilt dies auf typischen Intel-Prozessoren.

Der Grund dafür ist, dass der Prozessor ein Feature namens *Branch Prediction*
einsetzt, das bei einem `if` oder einer Schleife versucht, vorherzusagen, ob die
Bedingung wahr sein wird.
Auf dieser Basis führt der Prozessor Anweisungen im Voraus, sozusagen "auf Verdacht",
wenn er sonst gerade nichts zu tun hat. Liegt er mit seiner Vorhersage richtig,
verbessert er so die Performance. Liegt er jedoch falsch, ist dies relativ teuer.

Ist die Liste sortiert, gelingt in unserem Beispiel die Branch Prediction in vielen
Fällen, bei einer unsortierten Liste liegt sie im Mittel bei jedem 2. Mal falsch.
Dadurch ist die Summenfunktion mit einer sortierten Liste erheblich schneller.
