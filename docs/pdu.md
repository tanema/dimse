# PDU message descriptions
PDU stands for protocol data unit, and this doc overall describes the underlying
protocol that establishes an association with an AE, issues a command, and responds
to any subsequent commands.

## PDU DIMSE Flow
It seems like since you can only have 128(255/2) presentation contexts you should
mainly do one call at a time, or batch very few calls and then release. So one
full command process looks like:

- -> Assoc Request with one or more presentation context with ID, SOP and transfer syntax for types of data you will be handling.
- <- Assoc Accept with accepted presentation contexts.
- -> PDU PData Command(s) with AffectedSOPClassUID field that was added to the assoc request.
- -> PDU PData Data
- -> ...
- -> PDU Last Data
- <- PDU PData Command Resp
- <- PDU PData Data Resp
- <- ...
- <- PDU Last Data Resp
- -> Release Request
- <- Release Response

# PDU Definitions

## ASSOCIATE-RQ
association request

| bytes     | Field name         | Datatype   | Description of field                |
|-----------|--------------------|------------|-------------------------------------|
| 1         | PDU-type           | `uint8`    | `0x01`                              |
| 2         | Reserved           | `uint8`    | `0x00`                              |
| 3-6       | PDU-length         | `uint16`   |  length til end                     |
| 7-8       | Protocol-version   | `uint16`   |                                     |
| 9-10      | Reserved           | `uint16`   | `0x0000`                            |
| 11-26     | Called-AE-title    | [16]byte   | Destination DICOM Application Name. It shall be encoded as 16 characters as defined by the ISO 646:1990-Basic G0 Set with leading and trailing spaces being non-significant. The value made of 16 spaces meaning "no Application Name specified" shall not be used.
| 27-42     | Calling-AE-title   | [16]byte   | Source DICOM Application Name. It shall be encoded as 16 characters as defined by the ISO 646:1990-Basic G0 Set with leading and trailing spaces being non-significant. The value made of 16 spaces meaning "no Application Name specified" shall not be used.
| 43-74     | Reserved           | `uint32`   | `0x00`                              |
| 75-xxx    | Variable items     |            | [Application Context Item](#application-context-item), one or more [Presentation Context Items](#presentation-context-item-fields) and one [User Information Item](#user-information-item-fields). |

## ASSOCIATE-AC
association accepted

| bytes     | Field name         | Datatype   | Description of field                |
|-----------|--------------------|------------|-------------------------------------|
| 1         | PDU-type           | `uint8`    | `0x02`                              |
| 2         | Reserved           | `uint8`    | `0x00`                              |
| 3-6       | PDU-length         | `uint16`   | length til end                      |
| 7-8       | Protocol-version   | `uint16`   |                                     |
| 9-10      | Reserved           | `uint16`   | `0x0000`                            |
| 11-26     | Called-AE-title    | [16]byte   | Destination DICOM Application Name. It shall be encoded as 16 characters as defined by the ISO 646:1990-Basic G0 Set with leading and trailing spaces being non-significant. The value made of 16 spaces meaning "no Application Name specified" shall not be used.
| 27-42     | Calling-AE-title   | [16]byte   | Source DICOM Application Name. It shall be encoded as 16 characters as defined by the ISO 646:1990-Basic G0 Set with leading and trailing spaces being non-significant. The value made of 16 spaces meaning "no Application Name specified" shall not be used.
| 43-74     | Reserved           | `uint32`   | `0x00`                              |
| 75-xxx    | Variable items     |            | [Application Context Item](#application-context-item), one or more [Presentation Context Items](#presentation-context-item-fields) and one [User Information Item](#user-information-item-fields). |

## A-ASSOCIATE-RJ
Association request rejected

| bytes     | Field name         | Datatype   | Description of field                |
|-----------|--------------------|------------|-------------------------------------|
| 1         | PDU-type           | `uint8`    | `0x03`                              |
| 2         | Reserved           | `uint8`    | `0x00`                              |
| 3-6       | PDU-length         | `uint16`   |  length til end                     |
| 7         | Reserved           | `uint8`    | `0x00`                              |
| 8         | Result             | `uint8`    | 1: permanent, 2: transient          |
| 9         | Source             | `uint8`    | One of the following values shall be used: 1: DICOM UL service-user, 2: DICOM UL service-provider (ACSE related function), 3: DICOM UL service-provider (Presentation related function) |
| 10        | Reason/Diag.       | `uint8`    | [Reasons](#reject-reasons)          |

## P-DATA-TF
Used once an association has been established to send DIMSE message data.

| bytes     | Field name         | Datatype   | Description of field                |
|-----------|--------------------|------------|-------------------------------------|
| 1         | PDU-type           | `uint8`    | `0x04`                              |
| 2         | Reserved           | `uint8`    | `0x00`                              |
| 3-6       | PDU-length         | `uint16`   |  length til end                     |
| 7-xxx     | Presentation-data-value Item(s)|| contains one or more [Presentation-Data-Value Item Fields](#presentation-data-value-item-fields) |

## A-RELEASE-RQ and A-RELEASE-RP
Release request and response are exactly the same

| bytes     | Field name         | Datatype   | Description of field                |
|-----------|--------------------|------------|-------------------------------------|
| 1         | PDU-type           | `uint8`    | `0x05`                              |
| 2         | Reserved           | `uint8`    | `0x00`                              |
| 3-6       | PDU-length         | `uint16`   |  length til end                     |
| 7-10      | Reserved           | `uint32`   | `0x00000000`                        |

## A-ABORT

| bytes     | Field name         | Datatype   | Description of field                |
|-----------|--------------------|------------|-------------------------------------|
| 1         | PDU-type           | `uint8`    | `0x07`                              |
| 2         | Reserved           | `uint8`    | `0x00`                              |
| 3-6       | PDU-length         | `uint16`   |  length til end                     |
| 7         | Reserved           | `uint8`    | `0x00`                              |
| 8         | Reserved           | `uint8`    | `0x00`                              |
| 9         | Source             | `uint8`    | 0: service-user, 1: reserved, 2: service-provider |
| 10        | Reason/Diag.       | `uint8`    | [Reasons](#abort-reasons)           |

## Sub field definitions

### Application Context Item
| bytes     | Field name         | Datatype   | Description of field                |
|-----------|--------------------|------------|-------------------------------------|
| 1         | Item-type          | `uint8`    | `0x10`                              |
| 2         | Reserved           | `uint8`    | `0x00`                              |
| 3-6       | Item-length        | `uint16`   | length til end                      |
| 5-xxx     | Application-context-name|       | A valid Application-context-name    |

### Presentation Context Item Fields

| bytes     | Field name         | Datatype   | Description of field                |
|-----------|--------------------|------------|-------------------------------------|
| 1         | Item-type          | `uint8`    | `0x21`                              |
| 2         | Reserved           | `uint8`    | `0x00`                              |
| 3-6       | Item-length        | `uint16`   | length til end                      |
| 5         | Presentation-context-ID| `uint8`| odd integers between 1 and 255.     |
| 6         | Reserved           | `uint8`    | `0x00`                              |
| 7         | Result/Reason      | `uint8`    | 0: acceptance, 1: user-rejection, 2: no-reason, 3: abstract-syntax-not-supported, 4: transfer-syntaxes-not-supported |
| 8         | Reserved           | `uint8`    | `0x00`                              |
| 9-xxx     | Transfer syntax sub-item|        | one [Transfer Syntax Sub-Item](#transfer-syntax-sub-item). When the Result/Reason field has a value other than acceptance (0), this field shall not be significant and its value shall not be tested when received.

### Transfer Syntax Sub-Item

| bytes     | Field name         | Datatype   | Description of field                |
|-----------|--------------------|------------|-------------------------------------|
| 1         | Item-type          | `uint8`    | `0x40`                              |
| 2         | Reserved           | `uint8`    | `0x00`                              |
| 3-6       | Item-length        | `uint16`   | length til end                      |
| 5-xxx     | Transfer-syntax-name(s)|        | This variable field shall contain the Transfer-syntax-name proposed for this presentation context.

### Abstract Syntax Sub-Item

| bytes     | Field name         | Datatype   | Description of field                |
|-----------|--------------------|------------|-------------------------------------|
| 1         | Item-type          | `uint8`    | `0x30`                              |
| 2         | Reserved           | `uint8`    | `0x00`                              |
| 3-6       | Item-length        | `uint16`   | length til end                      |
| 5-xxx     | Abstract-syntax-name|           | This variable field shall contain the Abstract-syntax-name related to the proposed presentation context. |

### User Information Item Fields

| bytes     | Field name         | Datatype   | Description of field                |
|-----------|--------------------|------------|-------------------------------------|
| 1         | Item-type          | `uint8`    | `0x50`                              |
| 2         | Reserved           | `uint8`    | `0x00`                              |
| 3-6       | Item-length        | `uint16`   | length til end                      |
| 5-xxx     | User-data          |            | User-data sub-items as defined by the DICOM Application Entity. |

### Presentation-Data-Value Item Fields

| bytes     | Field name         | Datatype   | Description of field                |
|-----------|--------------------|------------|-------------------------------------|
| 1-4       | Item-length        | `uint32`    | number of bytes from the first byte of the following field to the last byte of the Presentation-data-value field. It shall be encoded as an unsigned binary number.
| 5         | Presentation-context-ID|`uint8` | odd integers between 1 and 255, encoded
| 6-xxx     | Presentation-data-value|        | contain DICOM message information (command and/or Data Set) with a message control header

### Reject Reasons
- If Source is 1
    - 1: no-reason-given
    - 2: application-context-name-not-supported
    - 3: calling-AE-title-not-recognized
    - 4-6: reserved
    - 7: called-AE-title-not-recognized
    - 8-10: reserved
- If the Source is 2
    - 1: no-reason-given
    - 2: protocol-version-not-supported
- If the Source is 3
    - 0: reserved
    - 1: temporary-congestion
    - 2: local-limit-exceeded
    - 3-7: reserved

### Abort Reasons
- If the Source field has the value (0) "DICOM UL service-user", this reason field
  shall not be significant. It shall be sent with a value `0x00` but not tested to
  this value when received.
- If the Source field has the value (2)
    - 0: reason-not-specified
    - 1: unrecognized-PDU
    - 2: unexpected-PDU
    - 3: reserved
    - 4: unrecognized-PDU parameter
    - 5: unexpected-PDU parameter
    - 6: invalid-PDU-parameter value
