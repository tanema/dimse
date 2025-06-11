# PDU message descriptions

## ASSOCIATE-RQ
association request

| bytes     | Field name                      | Description of field                |
|-----------|---------------------------------|-------------------------------------|
| 1         | PDU-type                        | `01H`                               |
| 2         | Reserved                        | `00H`                               |
| 3-6       | PDU-length                      | uint16 number of bytes from the first byte of the following field to the last byte of the variable field
| 7-8       | Protocol-version                | uint16                              |
| 9-10      | Reserved                        | `0000H`                             |
| 11-26     | Called-AE-title                 | Destination DICOM Application Name. It shall be encoded as 16 characters as defined by the ISO 646:1990-Basic G0 Set with leading and trailing spaces (20H) being non-significant. The value made of 16 spaces (20H) meaning "no Application Name specified" shall not be used.
| 27-42     | Calling-AE-title                | Source DICOM Application Name. It shall be encoded as 16 characters as defined by the ISO 646:1990-Basic G0 Set with leading and trailing spaces (20H) being non-significant. The value made of 16 spaces (20H) meaning "no Application Name specified" shall not be used.
| 43-74     | Reserved                        | `00H`                               |
| 75-xxx    | Variable items                  | [Application Context Item](#application-context-item), one or more [Presentation Context Items](#presentation-context-item-fields) and one [User Information Item](#user-information-item-fields). |

## ASSOCIATE-AC
association accepted

| bytes     | Field name                      | Description of field                |
|-----------|---------------------------------|-------------------------------------|
| 1         | PDU-type                        | `02H`                               |
| 2         | Reserved                        | `00H`                               |
| 3-6       | PDU-length                      | uint16 number of bytes from the first byte of the following field to the last byte of the variable field.
| 7-8       | Protocol-version                | uint16                              |
| 9-10      | Reserved                        | `0000H`                             |
| 11-26     | Called-AE-title                 | Destination DICOM Application Name. It shall be encoded as 16 characters as defined by the ISO 646:1990-Basic G0 Set with leading and trailing spaces (20H) being non-significant. The value made of 16 spaces (20H) meaning "no Application Name specified" shall not be used.
| 27-42     | Calling-AE-title                | Source DICOM Application Name. It shall be encoded as 16 characters as defined by the ISO 646:1990-Basic G0 Set with leading and trailing spaces (20H) being non-significant. The value made of 16 spaces (20H) meaning "no Application Name specified" shall not be used.
| 43-74     | Reserved                        | `00H`                               |
| 75-xxx    | Variable items                  | [Application Context Item](#application-context-item), one or more [Presentation Context Items](#presentation-context-item-fields) and one [User Information Item](#user-information-item-fields). |

## A-ASSOCIATE-RJ
Association request rejected

| bytes     | Field name                      | Description of field                |
|-----------|---------------------------------|-------------------------------------|
| 1         | PDU-type                        | `03H`                               |
| 2         | Reserved                        | `00H`                               |
| 3-6       | PDU-length                      | uint32 number of bytes from the first byte of the following field to the last byte of the Reason/Diag. field. In the case of this PDU, it shall have the fixed value of `00000004H` encoded as an unsigned binary number.
| 7         | Reserved                        | `00H`                               |
| 8         | Result                          | uint8. 1: permanent, 2: transient
| 9         | Source                          | uint8. One of the following values shall be used: 1: DICOM UL service-user, 2: DICOM UL service-provider (ACSE related function), 3: DICOM UL service-provider (Presentation related function)
| 10        | Reason/Diag.                    | uint8. [See Reasons](#reasons)

## P-DATA-TF
Used once an association has been established to send DIMSE message data.

| bytes     | Field name                      | Description of field                |
|-----------|---------------------------------|-------------------------------------|
| 1         | PDU-type                        | `04H`
| 2         | Reserved                        | reserved field should be sent with a value `00H` but not tested to this value when received.
| 3-6       | PDU-length                      | uint16 number of bytes from the first byte of the following field to the last byte of the variable field.
| 7-xxx     | Presentation-data-value Item(s) | contains one or more Presentation-data-value Items(s)

## Sub field definitions

### Application Context Item
| bytes     | Field name                      | Description of field                |
|-----------|---------------------------------|-------------------------------------|
| 1         | Item-type                       | 10H                                 |
| 2         | Reserved                        | `00H`                               |
| 3-4       | Item-length                     | uint16 number of bytes from the first byte of the following field to the last byte of the Application-context-name field.
| 5-xxx     | Application-context-name        | A valid Application-context-name    |

### Transfer Syntax Sub-Item

| bytes     | Field name                      | Description of field                |
|-----------|---------------------------------|-------------------------------------|
| 1         | Item-type                       | `40H`                               |
| 2         | Reserved                        | `00H`                               |
| 3-4       | Item-length                     | uint16 number of bytes from the first byte of the following field to the last byte of the Transfer-syntax-name field(s).
| 5-xxx     | Transfer-syntax-name(s)         | This variable field shall contain the Transfer-syntax-name proposed for this presentation context.

### Presentation Context Item Fields

| bytes     | Field name                      | Description of field                |
|-----------|---------------------------------|-------------------------------------|
| 1         | Item-type                       | `21H`                               |
| 2         | Reserved                        | `00H`                               |
| 3-4       | Item-length                     | uint16 number of bytes from the first byte of the following field to the last byte of the Transfer Syntax Sub-Item.
| 5         | Presentation-context-ID         | uint8 values shall be odd integers between 1 and 255.
| 6         | Reserved                        | `00H`                               |
| 7         | Result/Reason                   | uint8 0: acceptance, 1: user-rejection, 2: no-reason, 3: abstract-syntax-not-supported, 4: transfer-syntaxes-not-supported
| 8         | Reserved                        | `00H`                               |
| 9-xxx     | Transfer syntax sub-item        | one Transfer Syntax Sub-Item. When the Result/Reason field has a value other than acceptance (0), this field shall not be significant and its value shall not be tested when received.

### Abstract Syntax Sub-Item

| bytes     | Field name                      | Description of field                |
|-----------|---------------------------------|-------------------------------------|
| 1         | Item-type                       | `30H`                               |
| 2         | Reserved                        | `00H`                               |
| 3-4       | Item-length                     | uint16 number of bytes from the first byte of the following field to the last byte of the Abstract-syntax-name field.
| 5-xxx     | Abstract-syntax-name            | This variable field shall contain the Abstract-syntax-name related to the proposed presentation context.

### User Information Item Fields

| bytes     | Field name                      | Description of field                |
|-----------|---------------------------------|-------------------------------------|
| 1         | Item-type                       | 50H                                 |
| 2         | Reserved                        | `00H`                               |
| 3-4       | Item-length                     | uint16 number of bytes from the first byte of the following field to the last byte of the User-data-information field(s)
| 5-xxx     | User-data                       | This variable field shall contain User-data sub-items as defined by the DICOM Application Entity.

### Presentation-Data-Value Item Fields

| bytes     | Field name                      | Description of field                |
|-----------|---------------------------------|-------------------------------------|
| 1-4       | Item-length                     | number of bytes from the first byte of the following field to the last byte of the Presentation-data-value field. It shall be encoded as an unsigned binary number.
| 5         | Presentation-context-ID         | odd integers between 1 and 255, encoded as an unsigned binary number
| 6-xxx     | Presentation-data-value         | contain DICOM message information (command and/or Data Set) with a message control header

### Reasons
If Source is 1
- 1: no-reason-given
- 2: application-context-name-not-supported
- 3: calling-AE-title-not-recognized
- 4-6: reserved
- 7: called-AE-title-not-recognized
- 8-10: reserved

If the Source is 2
- 1: no-reason-given
- 2: protocol-version-not-supported

If the Source is 3
- 0: reserved
- 1: temporary-congestion
- 2: local-limit-exceeded
- 3-7: reserved
