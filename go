#!/bin/bash
# import multiple ESRI ASCII grids into DSS

JARS=$( echo jar/*.jar|tr ' ' ':')

java -Djava.library.path=.  -classpath ${JARS}  hec.heclib.grid.Asc2DssGrid BATCH=batch.input

