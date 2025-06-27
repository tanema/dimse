# Commands
Commands are wrapped inside a P-Data-Tf that allows you to operate on DICOM files
in the PACS

## Type of Commands
- C-Echo: Just a ping command to ensure the connection/association is working.
- C-Find: Query for files that match a query, and receive back only the matches.
- C-Get: Kind of Deprecated, you can use c-move instead. Download a Dicom file on the same association.
- C-Move: Move dicom file, can be used to download dicom file just move file to yourself.
- C-Store: Send a dicom file to another AE

## C-Store

### C-Store-RQ

| Message Field                            | Tag        | VR | VM | Ind | Description of Field |
|------------------------------------------|------------|----|----|-----|----------------------|
| Command Group Length                     |(0000,0000) | UL | 1  | M   | The even number of bytes from the end of the value field to the beginning of the next group.
| Affected SOP Class UID                   |(0000,0002) | UI | 1  | M   | SOP Class UID of the SOP Instance to be stored.
| Command Field                            |(0000,0100) | US | 1  | MF  | 0x0001
| Message ID                               |(0000,0110) | US | 1  | M   |
| Priority                                 |(0000,0700) | US | 1  | M   | LOW = 0x0002 MEDIUM = 0x0000 HIGH = 0x0001
| Command Data Set Type                    |(0000,0800) | US | 1  | MF  | 0x0001 (NonNull)
| Affected SOP Instance UID                |(0000,1000) | UI | 1  | M   | UID of the SOP Instance to be stored.
| Move Originator Application Entity Title |(0000,1030) | AE | 1  | U   | DICOM AE Title of the DICOM AE that invoked the C-MOVE operation from which this C-STORE sub-operation is being performed.
| Move Originator Message ID               |(0000,1031) | US | 1  | U   | Message ID (0000,0110) of the C-MOVE-RQ Message from which this C-STORE sub-operations is being performed.
| Data Set                                 |(no tag)    | -  | -  | M   | Application-specific Data Set.

### C-Store-Rsp

| Message Field                            | Tag        | VR | VM | Ind | Description of Field |
|------------------------------------------|------------|----|----|-----|----------------------|
| Command Group Length                     |(0000,0000) | UL | 1  | M   | The even number of bytes from the end of the value field to the beginning of the next group.
| Affected SOP Class UID                   |(0000,0002) | UI | 1  | M   | SOP Class UID of the SOP Instance to be stored.
| Command Field                            |(0000,0100) | US | 1  | MF  | 0x8001
| Message ID                               |(0000,0110) | US | 1  | U   |
| Message ID Being Responded To            |(0000,0120) | US | 1  | M   |
| Command Data Set Type                    |(0000,0800) | US | 1  | MF  | 0x0101 (Null)
| Affected SOP Instance UID                |(0000,1000) | UI | 1  | M   | UID of the SOP Instance to be stored.
| Status                                   |(0000,0900) | US | 1  | M   | DICOM AE Title of the DICOM AE that invoked the C-MOVE operation from which this C-STORE sub-operation is being performed.

## C-Find

### C-Find-RQ

| Message Field                            | Tag        | VR | VM | Ind | Description of Field |
|------------------------------------------|------------|----|----|-----|----------------------|
| Command Group Length                     |(0000,0000) | UL | 1  | M   | The even number of bytes from the end of the value field to the beginning of the next group.
| Affected SOP Class UID                   |(0000,0002) | UI | 1  | M   | SOP Class UID of the SOP Instance to be stored.
| Command Field                            |(0000,0100) | US | 1  | MF  | 0x0020
| Message ID                               |(0000,0110) | US | 1  | M   |
| Priority                                 |(0000,0700) | US | 1  | M   | LOW = 0x0002 MEDIUM = 0x0000 HIGH = 0x0001
| Command Data Set Type                    |(0000,0800) | US | 1  | MF  | 0x0001 (NonNull)
| Identifier                               |(no tag)    | -  | -  | M   | A Data Set that encodes the Identifier to be matched.

### C-Find-Rsp

| Message Field                            | Tag        | VR | VM | Ind | Description of Field |
|------------------------------------------|------------|----|----|-----|----------------------|
| Command Group Length                     |(0000,0000) | UL | 1  | M   | The even number of bytes from the end of the value field to the beginning of the next group.
| Affected SOP Class UID                   |(0000,0002) | UI | 1  | M   | SOP Class UID of the SOP Instance to be stored.
| Command Field                            |(0000,0100) | US | 1  | MF  | 0x8020
| Message ID                               |(0000,0110) | US | 1  | U   |
| Message ID Being Responded To            |(0000,0120) | US | 1  | M   |
| Command Data Set Type                    |(0000,0800) | US | 1  | MF  | 0x0101 (Null) if no match 0x0001 (NonNull) if there was a match
| Affected SOP Instance UID                |(0000,1000) | UI | 1  | M   | UID of the SOP Instance to be stored.
| Status                                   |(0000,0900) | US | 1  | M   | DICOM AE Title of the DICOM AE that invoked the C-MOVE operation from which this C-STORE sub-operation is being performed.
| Identifier                               |(no tag)    | -  | -  | M   | A Data Set that encodes the Identifier that was matched.

## C-Get

### C-Get-RQ

| Message Field                            | Tag        | VR | VM | Ind | Description of Field |
|------------------------------------------|------------|----|----|-----|----------------------|
| Command Group Length                     |(0000,0000) | UL | 1  | M   | The even number of bytes from the end of the value field to the beginning of the next group.
| Affected SOP Class UID                   |(0000,0002) | UI | 1  | M   | SOP Class UID of the SOP Instance to be stored.
| Command Field                            |(0000,0100) | US | 1  | MF  | 0x0010
| Message ID                               |(0000,0110) | US | 1  | M   |
| Priority                                 |(0000,0700) | US | 1  | M   | LOW = 0x0002 MEDIUM = 0x0000 HIGH = 0x0001
| Command Data Set Type                    |(0000,0800) | US | 1  | MF  | 0x0001 (NonNull)
| Identifier                               |(no tag)    | -  | -  | M   | A Data Set that encodes attributes providing status information about the C-GET operation

### C-Get-Rsp

| Message Field                            | Tag        | VR | VM | Ind | Description of Field |
|------------------------------------------|------------|----|----|-----|----------------------|
| Command Group Length                     |(0000,0000) | UL | 1  | M   | The even number of bytes from the end of the value field to the beginning of the next group.
| Affected SOP Class UID                   |(0000,0002) | UI | 1  | M   | SOP Class UID of the SOP Instance to be stored.
| Command Field                            |(0000,0100) | US | 1  | MF  | 0x8020
| Message ID                               |(0000,0110) | US | 1  | U   |
| Message ID Being Responded To            |(0000,0120) | US | 1  | M   |
| Command Data Set Type                    |(0000,0800) | US | 1  | MF  | 0x0101 (Null) if no match 0x0001 (NonNull) if there was a match
| Affected SOP Instance UID                |(0000,1000) | UI | 1  | M   | UID of the SOP Instance to be stored.
| Status                                   |(0000,0900) | US | 1  | M   | DICOM AE Title of the DICOM AE that invoked the C-MOVE operation from which this C-STORE sub-operation is being performed.
| Number of Remaining Sub-operations       |(0000,1020) | US | 1  | C   | The number of remaining C-STORE sub-operations to be invoked for this C-GET operation.
| Number of Completed Sub-operations       |(0000,1021) | US | 1  | C   | The number of C-STORE sub-operations invoked by this C-GET operation that have completed successfully.
| Number of Failed Sub-operations          |(0000,1022) | US | 1  | C   | The number of C-STORE sub-operations invoked by this C-GET operation that have failed.
| Number of Warning Sub-operations         |(0000,1023) | US | 1  | C   | The number of C-STORE sub-operations invoked by this C-GET operation that generated warning responses.
| Identifier                               |(no tag)    | -  | -  | M   | A Data Set that encodes the Identifier that was matched.

## C-Move

### C-Move-RQ

| Message Field                            | Tag        | VR | VM | Ind | Description of Field |
|------------------------------------------|------------|----|----|-----|----------------------|
| Command Group Length                     |(0000,0000) | UL | 1  | M   | The even number of bytes from the end of the value field to the beginning of the next group.
| Affected SOP Class UID                   |(0000,0002) | UI | 1  | M   | SOP Class UID of the SOP Instance to be stored.
| Command Field                            |(0000,0100) | US | 1  | MF  | 0x0021
| Message ID                               |(0000,0110) | US | 1  | M   |
| Priority                                 |(0000,0700) | US | 1  | M   | LOW = 0x0002 MEDIUM = 0x0000 HIGH = 0x0001
| Command Data Set Type                    |(0000,0800) | US | 1  | MF  | 0x0001 (NonNull)
| Move Destination                         |(0000,0600) | AE | 1  | M   | Shall be set to the DICOM AE Title of the destination DICOM AE to which the C-STORE sub-operations are being performed.
| Identifier                               |(no tag)    | -  | -  | M   | A Data Set that encodes attributes providing status information about the C-GET operation

### C-Move-Rsp

| Message Field                            | Tag        | VR | VM | Ind | Description of Field |
|------------------------------------------|------------|----|----|-----|----------------------|
| Command Group Length                     |(0000,0000) | UL | 1  | M   | The even number of bytes from the end of the value field to the beginning of the next group.
| Affected SOP Class UID                   |(0000,0002) | UI | 1  | M   | SOP Class UID of the SOP Instance to be stored.
| Command Field                            |(0000,0100) | US | 1  | MF  | 0x8021
| Message ID                               |(0000,0110) | US | 1  | U   |
| Message ID Being Responded To            |(0000,0120) | US | 1  | M   |
| Command Data Set Type                    |(0000,0800) | US | 1  | MF  | 0x0101 (Null) if no match 0x0001 (NonNull) if there was a match
| Affected SOP Instance UID                |(0000,1000) | UI | 1  | M   | UID of the SOP Instance to be stored.
| Status                                   |(0000,0900) | US | 1  | M   | DICOM AE Title of the DICOM AE that invoked the C-MOVE operation from which this C-STORE sub-operation is being performed.
| Number of Remaining Sub-operations       |(0000,1020) | US | 1  | C   | The number of remaining C-STORE sub-operations to be invoked for this C-GET operation.
| Number of Completed Sub-operations       |(0000,1021) | US | 1  | C   | The number of C-STORE sub-operations invoked by this C-GET operation that have completed successfully.
| Number of Failed Sub-operations          |(0000,1022) | US | 1  | C   | The number of C-STORE sub-operations invoked by this C-GET operation that have failed.
| Number of Warning Sub-operations         |(0000,1023) | US | 1  | C   | The number of C-STORE sub-operations invoked by this C-GET operation that generated warning responses.
| Identifier                               |(no tag)    | -  | -  | M   | A Data Set that encodes the Identifier that was matched.

## C-Echo

### C-Echo-RQ

| Message Field                            | Tag        | VR | VM | Ind | Description of Field |
|------------------------------------------|------------|----|----|-----|----------------------|
| Command Group Length                     |(0000,0000) | UL | 1  | M   | The even number of bytes from the end of the value field to the beginning of the next group.
| Affected SOP Class UID                   |(0000,0002) | UI | 1  | M   | SOP Class UID of the SOP Instance to be stored.
| Command Field                            |(0000,0100) | US | 1  | MF  | 0x0030
| Message ID                               |(0000,0110) | US | 1  | M   |
| Command Data Set Type                    |(0000,0800) | US | 1  | MF  | 0x0001 (NonNull)

### C-Echo-Rsp

| Message Field                            | Tag        | VR | VM | Ind | Description of Field |
|------------------------------------------|------------|----|----|-----|----------------------|
| Command Group Length                     |(0000,0000) | UL | 1  | M   | The even number of bytes from the end of the value field to the beginning of the next group.
| Affected SOP Class UID                   |(0000,0002) | UI | 1  | M   | SOP Class UID of the SOP Instance to be stored.
| Command Field                            |(0000,0100) | US | 1  | MF  | 0x8030
| Message ID                               |(0000,0110) | US | 1  | U   |
| Message ID Being Responded To            |(0000,0120) | US | 1  | M   |
| Command Data Set Type                    |(0000,0800) | US | 1  | MF  | 0x0101 (Null)
| Status                                   |(0000,0900) | US | 1  | M   | DICOM AE Title of the DICOM AE that invoked the C-MOVE operation from which this C-STORE sub-operation is being performed.

## Usage Inds

| Param | Meaning                       |
|-------|-------------------------------|
| (=)   | resp field match req field    |
| C     | conditional                   |
| M     | mandatory                     |
| MF    | mandatory with a fixed value  |
| U     | optional                      |
| UF    | optional with a fixed value   |

## Value Representations

| VR | VR Name            | Length   |  Definition       |
|----|--------------------|----------|-------------------|
| AE | Application Entity | 16b max  | String that identifies an Application Entity with leading and trailing spaces (20H) being non-significant. A value consisting solely of spaces shall not be used.
| AS | Age String         | 4b fixed | A string of characters with one of the following formats -- nnnD, nnnW, nnnM, nnnY; where nnn shall contain the number of days for D, weeks for W, months for M, or years for Y. Example: "018M" would represent an age of 18 months.
| AT | Attribute Tag      | 4b fixed | Ordered pair of 16-bit unsigned integers that is the value of a Data Element Tag.
| CS | Code String        | 16b max  | A string of characters with leading or trailing spaces (20H) being non-significant.
| DA | Date               | 8b fixed | A string of characters of the format YYYYMMDD; where YYYY shall contain year, MM shall contain the month, and DD shall contain the day, interpreted as a date of the Gregorian calendar system. "19930822" would represent August 22, 1993.
| DS | Decimal String     | 16b max  | A string of characters representing either a fixed point number or a floating point number. A fixed point number shall contain only the characters 0-9 with an optional leading "+" or "-" and an optional "." to mark the decimal point. A floating point number shall be conveyed as defined in ANSI X3.9, with an "E" or "e" to indicate the start of the exponent. Decimal Strings may be padded with leading or trailing spaces. Embedded spaces are not allowed.
| DT | Date Time          | 26b max  | A concatenated date-time character string in the format: YYYYMMDDHHMMSS.FFFFFF&ZZXX
| FL | Float Point Single | 4b fixed | Single precision binary floating point number represented in IEEE 754:1985 32-bit Floating Point Number Format.
| FD | Float Point Double | 8b fixed | Double precision binary floating point number represented in IEEE 754:1985 64-bit Floating Point Number Format.
| IS | Integer String     | 12b max  | A string of characters representing an Integer in base-10 (decimal), shall contain only the characters 0 - 9, with an optional leading "+" or "-". It may be padded with leading and/or trailing spaces. Embedded spaces are not allowed.
| LO | Long String        | 64chrMax | A character string that may be padded with leading and/or trailing spaces. The character code 5CH (the BACKSLASH "\" in ISO-IR 6) shall not be present, as it is used as the delimiter between values in multiple valued data elements. The string shall not have Control Characters except for ESC.
| LT | Long Text          |10240chrMx| A character string that may contain one or more paragraphs. It may contain the Graphic Character set and the Control Characters, CR, LF, FF, and ESC. It may be padded with trailing spaces, which may be ignored, but leading spaces are considered to be significant. Data Elements with this VR shall not be multi-valued and therefore character code 5CH (the BACKSLASH "\" in ISO-IR 6) may be used.
| OB | Other Byte String  |          | A string of bytes where the encoding of the contents is specified by the negotiated Transfer Syntax. OB is a VR that is insensitive to Little/Big Endian byte ordering (see Section 7.3). The string of bytes shall be padded with a single trailing NULL byte value (00H) when necessary to achieve even length.
| OD | Other Double String|232-8b max| A string of 64-bit IEEE 754:1985 floating point words. OD is a VR that requires byte swapping within each 64-bit word when changing between Little Endian and Big Endian byte ordering (see Section 7.3).
| OF | Other Float String |232-4b max| A string of 32-bit IEEE 754:1985 floating point words. OF is a VR that requires byte swapping within each 32-bit word when changing between Little Endian and Big Endian byte ordering (see Section 7.3).
| OW | Other Word String  |          | A string of 16-bit words where the encoding of the contents is specified by the negotiated Transfer Syntax. OW is a VR that requires byte swapping within each word when changing between Little Endian and Big Endian byte ordering (see Section 7.3).
| PN | Person Name        |64chr max | A character string encoded using a 5 component convention. The character code 5CH (the BACKSLASH "\" in ISO-IR 6) shall not be present, as it is used as the delimiter between values in multiple valued data elements. The string may be padded with trailing spaces. For human use, the five components in their order of occurrence are: family name complex, given name complex, middle name, name prefix, name suffix.
| SH | Short String       |64chr max | A character string that may be padded with leading and/or trailing spaces. The character code 05CH (the BACKSLASH "\" in ISO-IR 6) shall not be present, as it is used as the delimiter between values for multiple data elements. The string shall not have Control Characters except ESC.
| SL | Signed Long        | 4b fixed | Signed binary integer 32 bits long in 2's complement form. Represents an integer, n, in the range: - 231<= n <= 231-1.
| SQ | Sequence of Items  |          | Value is a Sequence of zero or more Items, as defined in Section 7.5.
| SS | Signed Short       | 2b fixed | Signed binary integer 16 bits long in 2's complement form. Represents an integer n in the range: -215<= n <= 215-1.
| ST | Short Text         |1024chrMax| A character string that may contain one or more paragraphs. It may contain the Graphic Character set and the Control Characters, CR, LF, FF, and ESC. It may be padded with trailing spaces, which may be ignored, but leading spaces are considered to be significant. Data Elements with this VR shall not be multi-valued and therefore character code 5CH (the BACKSLASH "\" in ISO-IR 6) may be used.
| TM | Time               | 16b max  | A string of characters of the format HHMMSS.FFFFFF; where HH contains hours (range "00" - "23"), MM contains minutes (range "00" - "59"), SS contains seconds (range "00" - "60"), and FFFFFF contains a fractional part of a second as small as 1 millionth of a second (range "000000" - "999999"). A 24-hour clock is used. Midnight shall be represented by only "0000" since "2400" would violate the hour range. The string may be padded with trailing spaces. Leading and embedded spaces are not allowed.
| UI | UID                | 64b max  | A character string containing a UID that is used to uniquely identify a wide variety of items. The UID is a series of numeric components separated by the period "." character. If a Value Field containing one or more UIDs is an odd number of bytes in length, the Value Field shall be padded with a single trailing NULL (00H) character to ensure that the Value Field is an even number of bytes in length. See Section 9 and Annex B for a complete specification and examples.
| UL | Unsigned Long      | 4b fixed | Unsigned binary integer 32 bits long. Represents an integer n in the range:
| UN | Unknown            |          | A string of bytes where the encoding of the contents is unknown (see Section 6.2.2).
| US | Unsigned Short     | 2b fixed | Unsigned binary integer 16 bits long. Represents integer n in the range:
| UT | Unlimited Text     |232-2b max| A character string that may contain one or more paragraphs. It may contain the Graphic Character set and the Control Characters, CR, LF, FF, and ESC. It may be padded with trailing spaces, which may be ignored, but leading spaces are considered to be significant. Data Elements with this VR shall not be multi-valued and therefore character code 5CH (the BACKSLASH "\" in ISO-IR 6) may be used.
