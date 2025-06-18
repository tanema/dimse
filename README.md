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

## Scratch Area
Server sending this in response to find but not reading it all for some reason.
```
PCID = 1
(0000,0000) UL : 76 (4CH)
(0000,0002) UI : 1.2.840.10008.5.1.4.1.2.1.1 (PatientRootQR_FIND)
(0000,0100) US : 32800 (8020H)
(0000,0120) US : 2 (2H)
(0000,0800) US : 258 (102H)
(0000,0900) US : 65280 (FF00H)
(0008,0005) CS : {null}
(0008,0052) CS : PATIENT
(0008,0054) AE : Array of 1 elements (anon-called-ae)
(0010,0020) LO : 3af4bf39-601f-4917-a577-9bbbc8b99366
PCID = 1
(0000,0000) UL : 76 (4CH)
(0000,0002) UI : 1.2.840.10008.5.1.4.1.2.1.1 (PatientRootQR_FIND)
(0000,0100) US : 32800 (8020H)
(0000,0120) US : 2 (2H)
(0000,0800) US : 257 (101H)
(0000,0900) US : 0 (0H)
```

Reading this from connection
```
(0000,0000) UL : [76]
(0000,0002) UI : [1.2.840.10008.5.1.4.1.2.1.1]
(0000,0100) US : [32800]
(0000,0120) US : [2]
(0000,0800) US : [258]
(0000,0900) US : [65280]
```

