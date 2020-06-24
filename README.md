# batch-Asc2Grid

import ESRI ASCII format grids into DSS  http://resources.esri.com/help/9.3/arcgisengine/java/GP_ToolRef/spatial_analyst_tools/esri_ascii_raster_format.htm

tested on CentOS 7, and Windows 10.

DSS binaries are here:

https://www.hec.usace.army.mil/nexus/#browse/browse:heclib:7-HT%2FLinux.zip

You will also need these jars:   (from latest DssVue)
hec.jar  lookup.jar  rma.jar

To control the DSS version (6 or 7) create an empty DSS file and use that in destination dss file argument in batch.input (dss=t1.dss)
