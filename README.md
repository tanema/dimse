# DIMSE for golang

>[!CAUTION]
> WIP do not use

Trying to use [github.com/suyashkumar/dicom](http://github.com/suyashkumar/dicom)
as the basis for a DIMSE implementation

## Implementation Progress
- [x] C-GET
- [x] C-FIND
- [x] C-STORE
- [x] C-ECHO
- [ ] C-MOVE
    - Need to be running a local PACS to test c-move since it will open a new connection
      to initiate a c-store.
- [ ] Testing

## Reference
- [Abbreviations](./docs/abbreviations.md)
- [PDU Notes](./docs/pdu.md)
- [Commands](./docs/commands.md)
- [Dev DICOM server](https://dicomserver.co.uk/logs/)
- [github.com/grailbio/go-netdicom](https://github.com/grailbio/go-netdicom)
- [github.com/pydicom/pynetdicom](https://github.com/pydicom/pynetdicom)
- [dicomstandard](https://www.dicomstandard.org/standards/view/message-exchange)
- [spec](https://dicom.nema.org/medical/dicom/current/output/chtml/part08/PS3.8.html)
