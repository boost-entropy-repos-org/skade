package pefiles

// the final PE File structure that should hold all information
type PEFile struct {
	FileName          string
	DosHeader         DosHeader
	NtHeader          NtHeader
	FileHeader        FileHeader
	OptionalHeader    OptionalHeader
	OptionalHeader64  OptionalHeader64
	Sections          []SectionHeader
	importDescriptors ImportDescriptors
	ExportDirectory   ExportDirectory

	dataLen   uint32
	headerEnd uint32
}

type DosHeader struct {
	Data       DosHeaderData
	FileOffset uint32
	Flags      map[string]bool
	Size       uint32
}

type NtHeader struct {
	Data       NtHeaderData
	FileOffset uint32
	Flags      map[string]bool
	Size       uint32
}

type FileHeader struct {
	Data       FileHeaderData
	FileOffset uint32
	Flags      map[string]bool
	Size       uint32
}

type OptionalHeader struct {
	Data       OptionalHeaderData
	FileOffset uint32
	Flags      map[string]bool
	Size       uint32
}

type OptionalHeader64 struct {
}

type ImportDescriptors struct {
}

type SectionHeader struct {
}

type ExportDirectory struct {
}

type NtHeaderData struct {
	Signature uint32
}

type OptionalHeaderData struct {
	Magic                       uint16
	MajorLinkerVersion          uint8
	MinorLinkerVersion          uint8
	SizeOfCode                  uint32
	SizeOfInitializedData       uint32
	SizeOfUninitializedData     uint32
	AddressOfEntryPoint         uint32
	BaseOfCode                  uint32
	BaseOfData                  uint32
	ImageBase                   uint32
	SectionAlignment            uint32
	FileAlignment               uint32
	MajorOperatingSystemVersion uint16
	MinorOperatingSystemVersion uint16
	MajorImageVersion           uint16
	MinorImageVersion           uint16
	MajorSubsystemVersion       uint16
	MinorSubsystemVersion       uint16
	Reserved1                   uint32
	SizeOfImage                 uint32
	SizeOfHeaders               uint32
	CheckSum                    uint32
	Subsystem                   uint16
	DllCharacteristics          uint16
	SizeOfStackReserve          uint32
	SizeOfStackCommit           uint32
	SizeOfHeapReserve           uint32
	SizeOfHeapCommit            uint32
	LoaderFlags                 uint32
	NumberOfRvaAndSizes         uint32
}

type FileHeaderData struct {
	Machine              uint16
	NumberOfSections     uint16
	TimeDateStamp        uint32
	PointerToSymbolTable uint32
	NumberOfSymbols      uint32
	SizeOfOptionalHeader uint16
	Characteristics      uint16
}

type DosHeaderData struct {
	E_magic    uint16    // Magic bytes aka MZ aka 0x5A4D
	E_cblp     uint16    // Number of bytes in the last page.
	E_cp       uint16    // Number of Pages in file
	E_crlc     uint16    // Number of entries in the relocation table.
	E_cparhd   uint16    // Number of paragraphs taken by the header - header size
	E_minalloc uint16    // Number of paragraphs require by the program - there has to be a free block this big
	E_maxalloc uint16    // Number of paragraphs requested by the program - if no block is this big, the closest one is taken
	E_ss       uint16    // Reolcatable Segment address for SS
	E_sp       uint16    // Inital value for SP
	E_csum     uint16    // When added to the sum of all other words in the file, the result should be zero.
	E_ip       uint16    // Inital value of IP
	E_cs       uint16    // Relocatable segment address for CS
	E_lfarlc   uint16    // Offset for the relocation table
	E_ovno     uint16    // Overlay number - If zero, this is the main executable.
	E_res      [8]uint8  // Reserved words
	E_oemid    uint16    // OEM identifier
	E_oeminfo  uint16    // OEM information
	E_res2     [20]uint8 // Reserved Words
	E_lfanew   uint32    // File address of the new exe header
}
