# DIMSE for golang

>[!CAUTION]
> WIP do not use

Trying to use [github.com/suyashkumar/dicom](http://github.com/suyashkumar/dicom) as the basis for a DIMSE implementation

*Current Issue* : dicom lib does not support reading elements without a header, and specifying
your own transfer syntax. Does not expose much that is of use at all. May need to
open an issue.

## Reference
- [PDU Notes](./docs/pdu.md)
- [Dev DICOM server](https://dicomserver.co.uk/logs/)
- [github.com/grailbio/go-netdicom](https://github.com/grailbio/go-netdicom)
- [github.com/pydicom/pynetdicom](https://github.com/pydicom/pynetdicom/blob/main/pynetdicom/association.py)
- [spec](https://dicom.nema.org/medical/dicom/current/output/chtml/part08/PS3.8.html)

## Client Design

- Client
  - Connection Pool
  - Command
    - context.Context for cancel
    - Get connection
    - Gather SOPs for Command
    - Associate with SOPs
    - Command
      - with AffectedSOPClassUID
      - Chunk Commands into PDUs if large
      - Send P-Data PDUs
      - Gather Received Chunks
      - Release
      - Abort if context.Cancel()
