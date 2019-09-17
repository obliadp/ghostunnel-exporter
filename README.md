# Ghostunnel exporter

Annoyed with having to deal with certs for metrics.

This opens up port 8080, and accepts ghostunnels metrics POST request.

Any GET will produce prometheus metrics available from the latest POST.
