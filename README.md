gox
===

Golang functions library

- archivex
  - zipx
    - Compress
    - UnCompress
- bytex
  - IsEmpty
  - IsBlank
  - ToString
  - StartsWith
  - EndsWith
  - Contains
- cryptox
  - Crc32
  - Md5
  - Sha1
- extractx
  - Number
  - Numbers
  - Float64
  - Float32
  - Int64
  - Int32
  - Int16
  - Int8
  - Int
- filepathx
  - Dirs
  - Files
  - GenerateDirNames
  - Ext
- filex
  - IsDir
  - IsFile
  - Exists
  - Size
- fmtx
  - SprettyPrint
  - PrettyPrint
  - PrettyPrintln
- htmlx
  - Strip
  - Spaceless
  - Clean
  - Tag
- inx
  - In
  - StringIn
  - IntIn
- ipx
  - RemoteAddr
  - LocalAddr
  - IsPrivate
  - IsPublic
  - Number
  - String
- isx
  - Number
  - Empty
  - Equal
  - SafeCharacters
  - HttpURL
  - OS
- jsonx
    - ToRawMessage
    - ToJson
    - ToPrettyJson
    - EmptyObjectRawMessage
    - EmptyArrayRawMessage
    - IsEmptyRawMessage
    - NewParser
      - Exists 
      - Find 
        - Interface
        - String
        - Int
        - Int64
        - Float32
        - Float64
        - Bool
- keyx
    - Generate
- map
    - Keys
    - StringMapStringEncode
- net
    - urlx
      - NewURL
      - GetValue
      - SetValue
      - AddValue
      - DelKey
      - HasKey
      - String
      - IsAbsolute
      - IsRelative
- nullx
    - StringFrom
    - NullString
    - TimeFrom
    - NullTime
- pathx
    - FilenameWithoutExt
- randx
  - Letter
  - Number
  - Any
- setx
  - ToSet
  - ToStringSet
  - ToIntSet
- slicex
  - ToInterface
  - StringToInterface
  - IntToInterface
  - StringSliceEqual
  - IntSliceEqual
  - StringSliceReverse
  - IntSliceReverse
  - Diff
  - StringSliceDiff
  - IntSliceDiff
  - Chunk
- spreedsheetx
  - NewColumn()
```go
  column := NewColumn("A")
  column.Next() // Return `B` if successful
  column.RightShift(26) // Return `AB` if successful
  column.LeftShift(1) // Return `AA` if successful
``` 


- stringx
  - IsEmpty
  - IsBlank
  - ToNumber
  - ContainsChinese
  - ToNarrow
  - ToWiden
  - Split
  - String
  - RemoveEmoji
  - TrimAny
  - RemoveExtraSpace
  - SequentialWordFields
  - ToBytes
  - WordMatched
  - StartsWith
  - EndsWith
  - Contains
  - QuoteMeta
  - HexToByte
  - Len
  - UpperFirst
  - LowerFirst
- timex
  - IsAmericaSummerTime
  - ChineseTimeLocation
  - Between
  - DayStart
  - DayEnd
  - MonthStart
  - MonthEnd
  - IsAM
  - IsPM
  - WeekStart
  - WeekEnd
  - YearWeeksByWeek
  - YearWeeksByTime
  - XISOWeek