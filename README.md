# DIMSE for golang

>[!CAUTION]
> WIP do not use

Trying to use [github.com/suyashkumar/dicom](http://github.com/suyashkumar/dicom)
as the basis for a DIMSE implementation

## Reference
- [PDU Notes](./docs/pdu.md)
- [dicomstandard](https://www.dicomstandard.org/standards/view/message-exchange)
- [Dev DICOM server](https://dicomserver.co.uk/logs/)
- [github.com/grailbio/go-netdicom](https://github.com/grailbio/go-netdicom)
- [github.com/pydicom/pynetdicom](https://github.com/pydicom/pynetdicom/blob/main/pynetdicom/association.py)
- [spec](https://dicom.nema.org/medical/dicom/current/output/chtml/part08/PS3.8.html)

## Terminology
- SCP: Service Class Provider (DICOM server)
- SCU: Service Class User     (client app or other DICOM server acting on SCP)
- PACS: picture archiving and communication systems
- DICOM: Digital Imaging & Communication in Medicine
- DIMSE: DICOM Message Service Element
- PDU: Protocol Data Unit
- ACSE: Association Control Service Element
- AE: Application Entity
- HL7: Health Level 7
- PDV: Presentation Data Value
- SOP: Service-Object Pair
- VM: Value Multiplicity specifies the number of Values that can be encoded in the Value Field of that Data Element.
- VR: Value Representation
- UID: Unique Identifier
- UL: Upper Layer
