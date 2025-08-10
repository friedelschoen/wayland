package xkb

//lint:file-ignore ST1003 Converted from C, partly case-sensitive (K_a vs K_A)

type Key int

const (
	K_NoSymbol   Key = 0x0
	K_VoidSymbol Key = 0xffffff /* Void symbol */

	/*
	 * TTY function keys, cleverly chosen to map to ASCII, for convenience of
	 * programming, but could have been arbitrary (at the cost of lookup
	 * tables in client code).
	 */

	K_BackSpace   Key = 0xff08 /* U+0008 BACKSPACE */
	K_Tab         Key = 0xff09 /* U+0009 CHARACTER TABULATION */
	K_Linefeed    Key = 0xff0a /* U+000A LINE FEED */
	K_Clear       Key = 0xff0b /* U+000B LINE TABULATION */
	K_Return      Key = 0xff0d /* U+000D CARRIAGE RETURN */
	K_Pause       Key = 0xff13 /* Pause, hold */
	K_Scroll_Lock Key = 0xff14
	K_Sys_Req     Key = 0xff15
	K_Escape      Key = 0xff1b /* U+001B ESCAPE */
	K_Delete      Key = 0xffff /* U+007F DELETE */

	/* International & multi-key character composition */

	K_Multi_key         Key = 0xff20 /* Multi-key character compose */
	K_Codeinput         Key = 0xff37
	K_SingleCandidate   Key = 0xff3c
	K_MultipleCandidate Key = 0xff3d
	K_PreviousCandidate Key = 0xff3e

	/* Japanese keyboard support */

	K_Kanji             Key = 0xff21 /* Kanji, Kanji convert */
	K_Muhenkan          Key = 0xff22 /* Cancel Conversion */
	K_Henkan_Mode       Key = 0xff23 /* Start/Stop Conversion */
	K_Henkan            Key = 0xff23 /* non-deprecated alias for Henkan_Mode */
	K_Romaji            Key = 0xff24 /* to Romaji */
	K_Hiragana          Key = 0xff25 /* to Hiragana */
	K_Katakana          Key = 0xff26 /* to Katakana */
	K_Hiragana_Katakana Key = 0xff27 /* Hiragana/Katakana toggle */
	K_Zenkaku           Key = 0xff28 /* to Zenkaku */
	K_Hankaku           Key = 0xff29 /* to Hankaku */
	K_Zenkaku_Hankaku   Key = 0xff2a /* Zenkaku/Hankaku toggle */
	K_Touroku           Key = 0xff2b /* Add to Dictionary */
	K_Massyo            Key = 0xff2c /* Delete from Dictionary */
	K_Kana_Lock         Key = 0xff2d /* Kana Lock */
	K_Kana_Shift        Key = 0xff2e /* Kana Shift */
	K_Eisu_Shift        Key = 0xff2f /* Alphanumeric Shift */
	K_Eisu_toggle       Key = 0xff30 /* Alphanumeric toggle */
	K_Kanji_Bangou      Key = 0xff37 /* Codeinput */
	K_Zen_Koho          Key = 0xff3d /* Multiple/All Candidate(s) */
	K_Mae_Koho          Key = 0xff3e /* Previous Candidate */

	/* 0xff31 thru 0xff3f are under XK_KOREAN */

	/* Cursor control & motion */

	K_Home      Key = 0xff50
	K_Left      Key = 0xff51 /* Move left, left arrow */
	K_Up        Key = 0xff52 /* Move up, up arrow */
	K_Right     Key = 0xff53 /* Move right, right arrow */
	K_Down      Key = 0xff54 /* Move down, down arrow */
	K_Prior     Key = 0xff55 /* Prior, previous */
	K_Page_Up   Key = 0xff55 /* deprecated alias for Prior */
	K_Next      Key = 0xff56 /* Next */
	K_Page_Down Key = 0xff56 /* deprecated alias for Next */
	K_End       Key = 0xff57 /* EOL */
	K_Begin     Key = 0xff58 /* BOL */

	/* Misc functions */

	K_Select        Key = 0xff60 /* Select, mark */
	K_Print         Key = 0xff61
	K_Execute       Key = 0xff62 /* Execute, run, do */
	K_Insert        Key = 0xff63 /* Insert, insert here */
	K_Undo          Key = 0xff65
	K_Redo          Key = 0xff66 /* Redo, again */
	K_Menu          Key = 0xff67
	K_Find          Key = 0xff68 /* Find, search */
	K_Cancel        Key = 0xff69 /* Cancel, stop, abort, exit */
	K_Help          Key = 0xff6a /* Help */
	K_Break         Key = 0xff6b
	K_Mode_switch   Key = 0xff7e /* Character set switch */
	K_script_switch Key = 0xff7e /* non-deprecated alias for Mode_switch */
	K_Num_Lock      Key = 0xff7f

	/* Keypad functions, keypad numbers cleverly chosen to map to ASCII */

	K_KP_Space     Key = 0xff80 /*<U+0020 SPACE>*/
	K_KP_Tab       Key = 0xff89 /*<U+0009 CHARACTER TABULATION>*/
	K_KP_Enter     Key = 0xff8d /*<U+000D CARRIAGE RETURN>*/
	K_KP_F1        Key = 0xff91 /* PF1, KP_A, ... */
	K_KP_F2        Key = 0xff92
	K_KP_F3        Key = 0xff93
	K_KP_F4        Key = 0xff94
	K_KP_Home      Key = 0xff95
	K_KP_Left      Key = 0xff96
	K_KP_Up        Key = 0xff97
	K_KP_Right     Key = 0xff98
	K_KP_Down      Key = 0xff99
	K_KP_Prior     Key = 0xff9a
	K_KP_Page_Up   Key = 0xff9a /* deprecated alias for KP_Prior */
	K_KP_Next      Key = 0xff9b
	K_KP_Page_Down Key = 0xff9b /* deprecated alias for KP_Next */
	K_KP_End       Key = 0xff9c
	K_KP_Begin     Key = 0xff9d
	K_KP_Insert    Key = 0xff9e
	K_KP_Delete    Key = 0xff9f
	K_KP_Equal     Key = 0xffbd /*<U+003D EQUALS SIGN>*/
	K_KP_Multiply  Key = 0xffaa /*<U+002A ASTERISK>*/
	K_KP_Add       Key = 0xffab /*<U+002B PLUS SIGN>*/
	K_KP_Separator Key = 0xffac /*<U+002C COMMA>*/
	K_KP_Subtract  Key = 0xffad /*<U+002D HYPHEN-MINUS>*/
	K_KP_Decimal   Key = 0xffae /*<U+002E FULL STOP>*/
	K_KP_Divide    Key = 0xffaf /*<U+002F SOLIDUS>*/

	K_KP_0 Key = 0xffb0 /*<U+0030 DIGIT ZERO>*/
	K_KP_1 Key = 0xffb1 /*<U+0031 DIGIT ONE>*/
	K_KP_2 Key = 0xffb2 /*<U+0032 DIGIT TWO>*/
	K_KP_3 Key = 0xffb3 /*<U+0033 DIGIT THREE>*/
	K_KP_4 Key = 0xffb4 /*<U+0034 DIGIT FOUR>*/
	K_KP_5 Key = 0xffb5 /*<U+0035 DIGIT FIVE>*/
	K_KP_6 Key = 0xffb6 /*<U+0036 DIGIT SIX>*/
	K_KP_7 Key = 0xffb7 /*<U+0037 DIGIT SEVEN>*/
	K_KP_8 Key = 0xffb8 /*<U+0038 DIGIT EIGHT>*/
	K_KP_9 Key = 0xffb9 /*<U+0039 DIGIT NINE>*/

	/*
	 * Auxiliary functions; note the duplicate definitions for left and right
	 * function keys;  Sun keyboards and a few other manufacturers have such
	 * function key groups on the left and/or right sides of the keyboard.
	 * We've not found a keyboard with more than 35 function keys total.
	 */

	K_F1  Key = 0xffbe
	K_F2  Key = 0xffbf
	K_F3  Key = 0xffc0
	K_F4  Key = 0xffc1
	K_F5  Key = 0xffc2
	K_F6  Key = 0xffc3
	K_F7  Key = 0xffc4
	K_F8  Key = 0xffc5
	K_F9  Key = 0xffc6
	K_F10 Key = 0xffc7
	K_F11 Key = 0xffc8
	K_L1  Key = 0xffc8 /* deprecated alias for F11 */
	K_F12 Key = 0xffc9
	K_L2  Key = 0xffc9 /* deprecated alias for F12 */
	K_F13 Key = 0xffca
	K_L3  Key = 0xffca /* deprecated alias for F13 */
	K_F14 Key = 0xffcb
	K_L4  Key = 0xffcb /* deprecated alias for F14 */
	K_F15 Key = 0xffcc
	K_L5  Key = 0xffcc /* deprecated alias for F15 */
	K_F16 Key = 0xffcd
	K_L6  Key = 0xffcd /* deprecated alias for F16 */
	K_F17 Key = 0xffce
	K_L7  Key = 0xffce /* deprecated alias for F17 */
	K_F18 Key = 0xffcf
	K_L8  Key = 0xffcf /* deprecated alias for F18 */
	K_F19 Key = 0xffd0
	K_L9  Key = 0xffd0 /* deprecated alias for F19 */
	K_F20 Key = 0xffd1
	K_L10 Key = 0xffd1 /* deprecated alias for F20 */
	K_F21 Key = 0xffd2
	K_R1  Key = 0xffd2 /* deprecated alias for F21 */
	K_F22 Key = 0xffd3
	K_R2  Key = 0xffd3 /* deprecated alias for F22 */
	K_F23 Key = 0xffd4
	K_R3  Key = 0xffd4 /* deprecated alias for F23 */
	K_F24 Key = 0xffd5
	K_R4  Key = 0xffd5 /* deprecated alias for F24 */
	K_F25 Key = 0xffd6
	K_R5  Key = 0xffd6 /* deprecated alias for F25 */
	K_F26 Key = 0xffd7
	K_R6  Key = 0xffd7 /* deprecated alias for F26 */
	K_F27 Key = 0xffd8
	K_R7  Key = 0xffd8 /* deprecated alias for F27 */
	K_F28 Key = 0xffd9
	K_R8  Key = 0xffd9 /* deprecated alias for F28 */
	K_F29 Key = 0xffda
	K_R9  Key = 0xffda /* deprecated alias for F29 */
	K_F30 Key = 0xffdb
	K_R10 Key = 0xffdb /* deprecated alias for F30 */
	K_F31 Key = 0xffdc
	K_R11 Key = 0xffdc /* deprecated alias for F31 */
	K_F32 Key = 0xffdd
	K_R12 Key = 0xffdd /* deprecated alias for F32 */
	K_F33 Key = 0xffde
	K_R13 Key = 0xffde /* deprecated alias for F33 */
	K_F34 Key = 0xffdf
	K_R14 Key = 0xffdf /* deprecated alias for F34 */
	K_F35 Key = 0xffe0
	K_R15 Key = 0xffe0 /* deprecated alias for F35 */

	/* Modifiers */
	K_Shift_L    Key = 0xffe1 /* Left shift */
	K_Shift_R    Key = 0xffe2 /* Right shift */
	K_Control_L  Key = 0xffe3 /* Left control */
	K_Control_R  Key = 0xffe4 /* Right control */
	K_Caps_Lock  Key = 0xffe5 /* Caps lock */
	K_Shift_Lock Key = 0xffe6 /* Shift lock */

	K_Meta_L  Key = 0xffe7 /* Left meta */
	K_Meta_R  Key = 0xffe8 /* Right meta */
	K_Alt_L   Key = 0xffe9 /* Left alt */
	K_Alt_R   Key = 0xffea /* Right alt */
	K_Super_L Key = 0xffeb /* Left super */
	K_Super_R Key = 0xffec /* Right super */
	K_Hyper_L Key = 0xffed /* Left hyper */
	K_Hyper_R Key = 0xffee /* Right hyper */

	/*
	 * Keyboard (XKB) Extension function and modifier keys
	 * (from Appendix C of "The X Keyboard Extension: Protocol Specification")
	 * Byte 3 = 0xfe
	 */

	K_ISO_Lock             Key = 0xfe01
	K_ISO_Level2_Latch     Key = 0xfe02
	K_ISO_Level3_Shift     Key = 0xfe03
	K_ISO_Level3_Latch     Key = 0xfe04
	K_ISO_Level3_Lock      Key = 0xfe05
	K_ISO_Level5_Shift     Key = 0xfe11
	K_ISO_Level5_Latch     Key = 0xfe12
	K_ISO_Level5_Lock      Key = 0xfe13
	K_ISO_Group_Shift      Key = 0xff7e /* non-deprecated alias for Mode_switch */
	K_ISO_Group_Latch      Key = 0xfe06
	K_ISO_Group_Lock       Key = 0xfe07
	K_ISO_Next_Group       Key = 0xfe08
	K_ISO_Next_Group_Lock  Key = 0xfe09
	K_ISO_Prev_Group       Key = 0xfe0a
	K_ISO_Prev_Group_Lock  Key = 0xfe0b
	K_ISO_First_Group      Key = 0xfe0c
	K_ISO_First_Group_Lock Key = 0xfe0d
	K_ISO_Last_Group       Key = 0xfe0e
	K_ISO_Last_Group_Lock  Key = 0xfe0f

	K_ISO_Left_Tab                Key = 0xfe20
	K_ISO_Move_Line_Up            Key = 0xfe21
	K_ISO_Move_Line_Down          Key = 0xfe22
	K_ISO_Partial_Line_Up         Key = 0xfe23
	K_ISO_Partial_Line_Down       Key = 0xfe24
	K_ISO_Partial_Space_Left      Key = 0xfe25
	K_ISO_Partial_Space_Right     Key = 0xfe26
	K_ISO_Set_Margin_Left         Key = 0xfe27
	K_ISO_Set_Margin_Right        Key = 0xfe28
	K_ISO_Release_Margin_Left     Key = 0xfe29
	K_ISO_Release_Margin_Right    Key = 0xfe2a
	K_ISO_Release_Both_Margins    Key = 0xfe2b
	K_ISO_Fast_Cursor_Left        Key = 0xfe2c
	K_ISO_Fast_Cursor_Right       Key = 0xfe2d
	K_ISO_Fast_Cursor_Up          Key = 0xfe2e
	K_ISO_Fast_Cursor_Down        Key = 0xfe2f
	K_ISO_Continuous_Underline    Key = 0xfe30
	K_ISO_Discontinuous_Underline Key = 0xfe31
	K_ISO_Emphasize               Key = 0xfe32
	K_ISO_Center_Object           Key = 0xfe33
	K_ISO_Enter                   Key = 0xfe34

	K_dead_grave              Key = 0xfe50
	K_dead_acute              Key = 0xfe51
	K_dead_circumflex         Key = 0xfe52
	K_dead_tilde              Key = 0xfe53
	K_dead_perispomeni        Key = 0xfe53 /* non-deprecated alias for dead_tilde */
	K_dead_macron             Key = 0xfe54
	K_dead_breve              Key = 0xfe55
	K_dead_abovedot           Key = 0xfe56
	K_dead_diaeresis          Key = 0xfe57
	K_dead_abovering          Key = 0xfe58
	K_dead_doubleacute        Key = 0xfe59
	K_dead_caron              Key = 0xfe5a
	K_dead_cedilla            Key = 0xfe5b
	K_dead_ogonek             Key = 0xfe5c
	K_dead_iota               Key = 0xfe5d
	K_dead_voiced_sound       Key = 0xfe5e
	K_dead_semivoiced_sound   Key = 0xfe5f
	K_dead_belowdot           Key = 0xfe60
	K_dead_hook               Key = 0xfe61
	K_dead_horn               Key = 0xfe62
	K_dead_stroke             Key = 0xfe63
	K_dead_abovecomma         Key = 0xfe64
	K_dead_psili              Key = 0xfe64 /* non-deprecated alias for dead_abovecomma */
	K_dead_abovereversedcomma Key = 0xfe65
	K_dead_dasia              Key = 0xfe65 /* non-deprecated alias for dead_abovereversedcomma */
	K_dead_doublegrave        Key = 0xfe66
	K_dead_belowring          Key = 0xfe67
	K_dead_belowmacron        Key = 0xfe68
	K_dead_belowcircumflex    Key = 0xfe69
	K_dead_belowtilde         Key = 0xfe6a
	K_dead_belowbreve         Key = 0xfe6b
	K_dead_belowdiaeresis     Key = 0xfe6c
	K_dead_invertedbreve      Key = 0xfe6d
	K_dead_belowcomma         Key = 0xfe6e
	K_dead_currency           Key = 0xfe6f

	/* extra dead elements for German T3 layout */
	K_dead_lowline            Key = 0xfe90
	K_dead_aboveverticalline  Key = 0xfe91
	K_dead_belowverticalline  Key = 0xfe92
	K_dead_longsolidusoverlay Key = 0xfe93

	/* dead vowels for universal syllable entry */
	K_dead_a             Key = 0xfe80
	K_dead_A             Key = 0xfe81
	K_dead_e             Key = 0xfe82
	K_dead_E             Key = 0xfe83
	K_dead_i             Key = 0xfe84
	K_dead_I             Key = 0xfe85
	K_dead_o             Key = 0xfe86
	K_dead_O             Key = 0xfe87
	K_dead_u             Key = 0xfe88
	K_dead_U             Key = 0xfe89
	K_dead_small_schwa   Key = 0xfe8a /* deprecated alias for dead_schwa */
	K_dead_schwa         Key = 0xfe8a
	K_dead_capital_schwa Key = 0xfe8b /* deprecated alias for dead_SCHWA */
	K_dead_SCHWA         Key = 0xfe8b

	K_dead_greek Key = 0xfe8c
	K_dead_hamza Key = 0xfe8d

	K_First_Virtual_Screen Key = 0xfed0
	K_Prev_Virtual_Screen  Key = 0xfed1
	K_Next_Virtual_Screen  Key = 0xfed2
	K_Last_Virtual_Screen  Key = 0xfed4
	K_Terminate_Server     Key = 0xfed5

	K_AccessX_Enable          Key = 0xfe70
	K_AccessX_Feedback_Enable Key = 0xfe71
	K_RepeatKeys_Enable       Key = 0xfe72
	K_SlowKeys_Enable         Key = 0xfe73
	K_BounceKeys_Enable       Key = 0xfe74
	K_StickyKeys_Enable       Key = 0xfe75
	K_MouseKeys_Enable        Key = 0xfe76
	K_MouseKeys_Accel_Enable  Key = 0xfe77
	K_Overlay1_Enable         Key = 0xfe78
	K_Overlay2_Enable         Key = 0xfe79
	K_AudibleBell_Enable      Key = 0xfe7a

	K_Pointer_Left          Key = 0xfee0
	K_Pointer_Right         Key = 0xfee1
	K_Pointer_Up            Key = 0xfee2
	K_Pointer_Down          Key = 0xfee3
	K_Pointer_UpLeft        Key = 0xfee4
	K_Pointer_UpRight       Key = 0xfee5
	K_Pointer_DownLeft      Key = 0xfee6
	K_Pointer_DownRight     Key = 0xfee7
	K_Pointer_Button_Dflt   Key = 0xfee8
	K_Pointer_Button1       Key = 0xfee9
	K_Pointer_Button2       Key = 0xfeea
	K_Pointer_Button3       Key = 0xfeeb
	K_Pointer_Button4       Key = 0xfeec
	K_Pointer_Button5       Key = 0xfeed
	K_Pointer_DblClick_Dflt Key = 0xfeee
	K_Pointer_DblClick1     Key = 0xfeef
	K_Pointer_DblClick2     Key = 0xfef0
	K_Pointer_DblClick3     Key = 0xfef1
	K_Pointer_DblClick4     Key = 0xfef2
	K_Pointer_DblClick5     Key = 0xfef3
	K_Pointer_Drag_Dflt     Key = 0xfef4
	K_Pointer_Drag1         Key = 0xfef5
	K_Pointer_Drag2         Key = 0xfef6
	K_Pointer_Drag3         Key = 0xfef7
	K_Pointer_Drag4         Key = 0xfef8
	K_Pointer_Drag5         Key = 0xfefd

	K_Pointer_EnableKeys  Key = 0xfef9
	K_Pointer_Accelerate  Key = 0xfefa
	K_Pointer_DfltBtnNext Key = 0xfefb
	K_Pointer_DfltBtnPrev Key = 0xfefc

	/* Single-Stroke Multiple-Character N-Graph Keysyms For The X Input Method */

	K_ch  Key = 0xfea0
	K_Ch  Key = 0xfea1
	K_CH  Key = 0xfea2
	K_c_h Key = 0xfea3
	K_C_h Key = 0xfea4
	K_C_H Key = 0xfea5

	/*
	 * 3270 Terminal Keys
	 * Byte 3 = 0xfd
	 */

	K_3270_Duplicate    Key = 0xfd01
	K_3270_FieldMark    Key = 0xfd02
	K_3270_Right2       Key = 0xfd03
	K_3270_Left2        Key = 0xfd04
	K_3270_BackTab      Key = 0xfd05
	K_3270_EraseEOF     Key = 0xfd06
	K_3270_EraseInput   Key = 0xfd07
	K_3270_Reset        Key = 0xfd08
	K_3270_Quit         Key = 0xfd09
	K_3270_PA1          Key = 0xfd0a
	K_3270_PA2          Key = 0xfd0b
	K_3270_PA3          Key = 0xfd0c
	K_3270_Test         Key = 0xfd0d
	K_3270_Attn         Key = 0xfd0e
	K_3270_CursorBlink  Key = 0xfd0f
	K_3270_AltCursor    Key = 0xfd10
	K_3270_KeyClick     Key = 0xfd11
	K_3270_Jump         Key = 0xfd12
	K_3270_Ident        Key = 0xfd13
	K_3270_Rule         Key = 0xfd14
	K_3270_Copy         Key = 0xfd15
	K_3270_Play         Key = 0xfd16
	K_3270_Setup        Key = 0xfd17
	K_3270_Record       Key = 0xfd18
	K_3270_ChangeScreen Key = 0xfd19
	K_3270_DeleteWord   Key = 0xfd1a
	K_3270_ExSelect     Key = 0xfd1b
	K_3270_CursorSelect Key = 0xfd1c
	K_3270_PrintScreen  Key = 0xfd1d
	K_3270_Enter        Key = 0xfd1e

	/*
	 * Latin 1
	 * (ISO/IEC 8859-1 = Unicode U+0020..U+00FF)
	 * Byte 3 = 0
	 */
	K_space        Key = 0x0020 /* U+0020 SPACE */
	K_exclam       Key = 0x0021 /* U+0021 EXCLAMATION MARK */
	K_quotedbl     Key = 0x0022 /* U+0022 QUOTATION MARK */
	K_numbersign   Key = 0x0023 /* U+0023 NUMBER SIGN */
	K_dollar       Key = 0x0024 /* U+0024 DOLLAR SIGN */
	K_percent      Key = 0x0025 /* U+0025 PERCENT SIGN */
	K_ampersand    Key = 0x0026 /* U+0026 AMPERSAND */
	K_apostrophe   Key = 0x0027 /* U+0027 APOSTROPHE */
	K_quoteright   Key = 0x0027 /* deprecated */
	K_parenleft    Key = 0x0028 /* U+0028 LEFT PARENTHESIS */
	K_parenright   Key = 0x0029 /* U+0029 RIGHT PARENTHESIS */
	K_asterisk     Key = 0x002a /* U+002A ASTERISK */
	K_plus         Key = 0x002b /* U+002B PLUS SIGN */
	K_comma        Key = 0x002c /* U+002C COMMA */
	K_minus        Key = 0x002d /* U+002D HYPHEN-MINUS */
	K_period       Key = 0x002e /* U+002E FULL STOP */
	K_slash        Key = 0x002f /* U+002F SOLIDUS */
	K_0            Key = 0x0030 /* U+0030 DIGIT ZERO */
	K_1            Key = 0x0031 /* U+0031 DIGIT ONE */
	K_2            Key = 0x0032 /* U+0032 DIGIT TWO */
	K_3            Key = 0x0033 /* U+0033 DIGIT THREE */
	K_4            Key = 0x0034 /* U+0034 DIGIT FOUR */
	K_5            Key = 0x0035 /* U+0035 DIGIT FIVE */
	K_6            Key = 0x0036 /* U+0036 DIGIT SIX */
	K_7            Key = 0x0037 /* U+0037 DIGIT SEVEN */
	K_8            Key = 0x0038 /* U+0038 DIGIT EIGHT */
	K_9            Key = 0x0039 /* U+0039 DIGIT NINE */
	K_colon        Key = 0x003a /* U+003A COLON */
	K_semicolon    Key = 0x003b /* U+003B SEMICOLON */
	K_less         Key = 0x003c /* U+003C LESS-THAN SIGN */
	K_equal        Key = 0x003d /* U+003D EQUALS SIGN */
	K_greater      Key = 0x003e /* U+003E GREATER-THAN SIGN */
	K_question     Key = 0x003f /* U+003F QUESTION MARK */
	K_at           Key = 0x0040 /* U+0040 COMMERCIAL AT */
	K_A            Key = 0x0041 /* U+0041 LATIN CAPITAL LETTER A */
	K_B            Key = 0x0042 /* U+0042 LATIN CAPITAL LETTER B */
	K_C            Key = 0x0043 /* U+0043 LATIN CAPITAL LETTER C */
	K_D            Key = 0x0044 /* U+0044 LATIN CAPITAL LETTER D */
	K_E            Key = 0x0045 /* U+0045 LATIN CAPITAL LETTER E */
	K_F            Key = 0x0046 /* U+0046 LATIN CAPITAL LETTER F */
	K_G            Key = 0x0047 /* U+0047 LATIN CAPITAL LETTER G */
	K_H            Key = 0x0048 /* U+0048 LATIN CAPITAL LETTER H */
	K_I            Key = 0x0049 /* U+0049 LATIN CAPITAL LETTER I */
	K_J            Key = 0x004a /* U+004A LATIN CAPITAL LETTER J */
	K_K            Key = 0x004b /* U+004B LATIN CAPITAL LETTER K */
	K_L            Key = 0x004c /* U+004C LATIN CAPITAL LETTER L */
	K_M            Key = 0x004d /* U+004D LATIN CAPITAL LETTER M */
	K_N            Key = 0x004e /* U+004E LATIN CAPITAL LETTER N */
	K_O            Key = 0x004f /* U+004F LATIN CAPITAL LETTER O */
	K_P            Key = 0x0050 /* U+0050 LATIN CAPITAL LETTER P */
	K_Q            Key = 0x0051 /* U+0051 LATIN CAPITAL LETTER Q */
	K_R            Key = 0x0052 /* U+0052 LATIN CAPITAL LETTER R */
	K_S            Key = 0x0053 /* U+0053 LATIN CAPITAL LETTER S */
	K_T            Key = 0x0054 /* U+0054 LATIN CAPITAL LETTER T */
	K_U            Key = 0x0055 /* U+0055 LATIN CAPITAL LETTER U */
	K_V            Key = 0x0056 /* U+0056 LATIN CAPITAL LETTER V */
	K_W            Key = 0x0057 /* U+0057 LATIN CAPITAL LETTER W */
	K_X            Key = 0x0058 /* U+0058 LATIN CAPITAL LETTER X */
	K_Y            Key = 0x0059 /* U+0059 LATIN CAPITAL LETTER Y */
	K_Z            Key = 0x005a /* U+005A LATIN CAPITAL LETTER Z */
	K_bracketleft  Key = 0x005b /* U+005B LEFT SQUARE BRACKET */
	K_backslash    Key = 0x005c /* U+005C REVERSE SOLIDUS */
	K_bracketright Key = 0x005d /* U+005D RIGHT SQUARE BRACKET */
	K_asciicircum  Key = 0x005e /* U+005E CIRCUMFLEX ACCENT */
	K_underscore   Key = 0x005f /* U+005F LOW LINE */
	K_grave        Key = 0x0060 /* U+0060 GRAVE ACCENT */
	K_quoteleft    Key = 0x0060 /* deprecated */
	K_a            Key = 0x0061 /* U+0061 LATIN SMALL LETTER A */
	K_b            Key = 0x0062 /* U+0062 LATIN SMALL LETTER B */
	K_c            Key = 0x0063 /* U+0063 LATIN SMALL LETTER C */
	K_d            Key = 0x0064 /* U+0064 LATIN SMALL LETTER D */
	K_e            Key = 0x0065 /* U+0065 LATIN SMALL LETTER E */
	K_f            Key = 0x0066 /* U+0066 LATIN SMALL LETTER F */
	K_g            Key = 0x0067 /* U+0067 LATIN SMALL LETTER G */
	K_h            Key = 0x0068 /* U+0068 LATIN SMALL LETTER H */
	K_i            Key = 0x0069 /* U+0069 LATIN SMALL LETTER I */
	K_j            Key = 0x006a /* U+006A LATIN SMALL LETTER J */
	K_k            Key = 0x006b /* U+006B LATIN SMALL LETTER K */
	K_l            Key = 0x006c /* U+006C LATIN SMALL LETTER L */
	K_m            Key = 0x006d /* U+006D LATIN SMALL LETTER M */
	K_n            Key = 0x006e /* U+006E LATIN SMALL LETTER N */
	K_o            Key = 0x006f /* U+006F LATIN SMALL LETTER O */
	K_p            Key = 0x0070 /* U+0070 LATIN SMALL LETTER P */
	K_q            Key = 0x0071 /* U+0071 LATIN SMALL LETTER Q */
	K_r            Key = 0x0072 /* U+0072 LATIN SMALL LETTER R */
	K_s            Key = 0x0073 /* U+0073 LATIN SMALL LETTER S */
	K_t            Key = 0x0074 /* U+0074 LATIN SMALL LETTER T */
	K_u            Key = 0x0075 /* U+0075 LATIN SMALL LETTER U */
	K_v            Key = 0x0076 /* U+0076 LATIN SMALL LETTER V */
	K_w            Key = 0x0077 /* U+0077 LATIN SMALL LETTER W */
	K_x            Key = 0x0078 /* U+0078 LATIN SMALL LETTER X */
	K_y            Key = 0x0079 /* U+0079 LATIN SMALL LETTER Y */
	K_z            Key = 0x007a /* U+007A LATIN SMALL LETTER Z */
	K_braceleft    Key = 0x007b /* U+007B LEFT CURLY BRACKET */
	K_bar          Key = 0x007c /* U+007C VERTICAL LINE */
	K_braceright   Key = 0x007d /* U+007D RIGHT CURLY BRACKET */
	K_asciitilde   Key = 0x007e /* U+007E TILDE */

	K_nobreakspace   Key = 0x00a0 /* U+00A0 NO-BREAK SPACE */
	K_exclamdown     Key = 0x00a1 /* U+00A1 INVERTED EXCLAMATION MARK */
	K_cent           Key = 0x00a2 /* U+00A2 CENT SIGN */
	K_sterling       Key = 0x00a3 /* U+00A3 POUND SIGN */
	K_currency       Key = 0x00a4 /* U+00A4 CURRENCY SIGN */
	K_yen            Key = 0x00a5 /* U+00A5 YEN SIGN */
	K_brokenbar      Key = 0x00a6 /* U+00A6 BROKEN BAR */
	K_section        Key = 0x00a7 /* U+00A7 SECTION SIGN */
	K_diaeresis      Key = 0x00a8 /* U+00A8 DIAERESIS */
	K_copyright      Key = 0x00a9 /* U+00A9 COPYRIGHT SIGN */
	K_ordfeminine    Key = 0x00aa /* U+00AA FEMININE ORDINAL INDICATOR */
	K_guillemotleft  Key = 0x00ab /* deprecated alias for guillemetleft (misspelling) */
	K_guillemetleft  Key = 0x00ab /* U+00AB LEFT-POINTING DOUBLE ANGLE QUOTATION MARK */
	K_notsign        Key = 0x00ac /* U+00AC NOT SIGN */
	K_hyphen         Key = 0x00ad /* U+00AD SOFT HYPHEN */
	K_registered     Key = 0x00ae /* U+00AE REGISTERED SIGN */
	K_macron         Key = 0x00af /* U+00AF MACRON */
	K_degree         Key = 0x00b0 /* U+00B0 DEGREE SIGN */
	K_plusminus      Key = 0x00b1 /* U+00B1 PLUS-MINUS SIGN */
	K_twosuperior    Key = 0x00b2 /* U+00B2 SUPERSCRIPT TWO */
	K_threesuperior  Key = 0x00b3 /* U+00B3 SUPERSCRIPT THREE */
	K_acute          Key = 0x00b4 /* U+00B4 ACUTE ACCENT */
	K_mu             Key = 0x00b5 /* U+00B5 MICRO SIGN */
	K_paragraph      Key = 0x00b6 /* U+00B6 PILCROW SIGN */
	K_periodcentered Key = 0x00b7 /* U+00B7 MIDDLE DOT */
	K_cedilla        Key = 0x00b8 /* U+00B8 CEDILLA */
	K_onesuperior    Key = 0x00b9 /* U+00B9 SUPERSCRIPT ONE */
	K_masculine      Key = 0x00ba /* deprecated alias for ordmasculine (inconsistent name) */
	K_ordmasculine   Key = 0x00ba /* U+00BA MASCULINE ORDINAL INDICATOR */
	K_guillemotright Key = 0x00bb /* deprecated alias for guillemetright (misspelling) */
	K_guillemetright Key = 0x00bb /* U+00BB RIGHT-POINTING DOUBLE ANGLE QUOTATION MARK */
	K_onequarter     Key = 0x00bc /* U+00BC VULGAR FRACTION ONE QUARTER */
	K_onehalf        Key = 0x00bd /* U+00BD VULGAR FRACTION ONE HALF */
	K_threequarters  Key = 0x00be /* U+00BE VULGAR FRACTION THREE QUARTERS */
	K_questiondown   Key = 0x00bf /* U+00BF INVERTED QUESTION MARK */
	K_Agrave         Key = 0x00c0 /* U+00C0 LATIN CAPITAL LETTER A WITH GRAVE */
	K_Aacute         Key = 0x00c1 /* U+00C1 LATIN CAPITAL LETTER A WITH ACUTE */
	K_Acircumflex    Key = 0x00c2 /* U+00C2 LATIN CAPITAL LETTER A WITH CIRCUMFLEX */
	K_Atilde         Key = 0x00c3 /* U+00C3 LATIN CAPITAL LETTER A WITH TILDE */
	K_Adiaeresis     Key = 0x00c4 /* U+00C4 LATIN CAPITAL LETTER A WITH DIAERESIS */
	K_Aring          Key = 0x00c5 /* U+00C5 LATIN CAPITAL LETTER A WITH RING ABOVE */
	K_AE             Key = 0x00c6 /* U+00C6 LATIN CAPITAL LETTER AE */
	K_Ccedilla       Key = 0x00c7 /* U+00C7 LATIN CAPITAL LETTER C WITH CEDILLA */
	K_Egrave         Key = 0x00c8 /* U+00C8 LATIN CAPITAL LETTER E WITH GRAVE */
	K_Eacute         Key = 0x00c9 /* U+00C9 LATIN CAPITAL LETTER E WITH ACUTE */
	K_Ecircumflex    Key = 0x00ca /* U+00CA LATIN CAPITAL LETTER E WITH CIRCUMFLEX */
	K_Ediaeresis     Key = 0x00cb /* U+00CB LATIN CAPITAL LETTER E WITH DIAERESIS */
	K_Igrave         Key = 0x00cc /* U+00CC LATIN CAPITAL LETTER I WITH GRAVE */
	K_Iacute         Key = 0x00cd /* U+00CD LATIN CAPITAL LETTER I WITH ACUTE */
	K_Icircumflex    Key = 0x00ce /* U+00CE LATIN CAPITAL LETTER I WITH CIRCUMFLEX */
	K_Idiaeresis     Key = 0x00cf /* U+00CF LATIN CAPITAL LETTER I WITH DIAERESIS */
	K_ETH            Key = 0x00d0 /* U+00D0 LATIN CAPITAL LETTER ETH */
	K_Eth            Key = 0x00d0 /* deprecated */
	K_Ntilde         Key = 0x00d1 /* U+00D1 LATIN CAPITAL LETTER N WITH TILDE */
	K_Ograve         Key = 0x00d2 /* U+00D2 LATIN CAPITAL LETTER O WITH GRAVE */
	K_Oacute         Key = 0x00d3 /* U+00D3 LATIN CAPITAL LETTER O WITH ACUTE */
	K_Ocircumflex    Key = 0x00d4 /* U+00D4 LATIN CAPITAL LETTER O WITH CIRCUMFLEX */
	K_Otilde         Key = 0x00d5 /* U+00D5 LATIN CAPITAL LETTER O WITH TILDE */
	K_Odiaeresis     Key = 0x00d6 /* U+00D6 LATIN CAPITAL LETTER O WITH DIAERESIS */
	K_multiply       Key = 0x00d7 /* U+00D7 MULTIPLICATION SIGN */
	K_Oslash         Key = 0x00d8 /* U+00D8 LATIN CAPITAL LETTER O WITH STROKE */
	K_Ooblique       Key = 0x00d8 /* deprecated alias for Oslash */
	K_Ugrave         Key = 0x00d9 /* U+00D9 LATIN CAPITAL LETTER U WITH GRAVE */
	K_Uacute         Key = 0x00da /* U+00DA LATIN CAPITAL LETTER U WITH ACUTE */
	K_Ucircumflex    Key = 0x00db /* U+00DB LATIN CAPITAL LETTER U WITH CIRCUMFLEX */
	K_Udiaeresis     Key = 0x00dc /* U+00DC LATIN CAPITAL LETTER U WITH DIAERESIS */
	K_Yacute         Key = 0x00dd /* U+00DD LATIN CAPITAL LETTER Y WITH ACUTE */
	K_THORN          Key = 0x00de /* U+00DE LATIN CAPITAL LETTER THORN */
	K_Thorn          Key = 0x00de /* deprecated */
	K_ssharp         Key = 0x00df /* U+00DF LATIN SMALL LETTER SHARP S */
	K_agrave         Key = 0x00e0 /* U+00E0 LATIN SMALL LETTER A WITH GRAVE */
	K_aacute         Key = 0x00e1 /* U+00E1 LATIN SMALL LETTER A WITH ACUTE */
	K_acircumflex    Key = 0x00e2 /* U+00E2 LATIN SMALL LETTER A WITH CIRCUMFLEX */
	K_atilde         Key = 0x00e3 /* U+00E3 LATIN SMALL LETTER A WITH TILDE */
	K_adiaeresis     Key = 0x00e4 /* U+00E4 LATIN SMALL LETTER A WITH DIAERESIS */
	K_aring          Key = 0x00e5 /* U+00E5 LATIN SMALL LETTER A WITH RING ABOVE */
	K_ae             Key = 0x00e6 /* U+00E6 LATIN SMALL LETTER AE */
	K_ccedilla       Key = 0x00e7 /* U+00E7 LATIN SMALL LETTER C WITH CEDILLA */
	K_egrave         Key = 0x00e8 /* U+00E8 LATIN SMALL LETTER E WITH GRAVE */
	K_eacute         Key = 0x00e9 /* U+00E9 LATIN SMALL LETTER E WITH ACUTE */
	K_ecircumflex    Key = 0x00ea /* U+00EA LATIN SMALL LETTER E WITH CIRCUMFLEX */
	K_ediaeresis     Key = 0x00eb /* U+00EB LATIN SMALL LETTER E WITH DIAERESIS */
	K_igrave         Key = 0x00ec /* U+00EC LATIN SMALL LETTER I WITH GRAVE */
	K_iacute         Key = 0x00ed /* U+00ED LATIN SMALL LETTER I WITH ACUTE */
	K_icircumflex    Key = 0x00ee /* U+00EE LATIN SMALL LETTER I WITH CIRCUMFLEX */
	K_idiaeresis     Key = 0x00ef /* U+00EF LATIN SMALL LETTER I WITH DIAERESIS */
	K_eth            Key = 0x00f0 /* U+00F0 LATIN SMALL LETTER ETH */
	K_ntilde         Key = 0x00f1 /* U+00F1 LATIN SMALL LETTER N WITH TILDE */
	K_ograve         Key = 0x00f2 /* U+00F2 LATIN SMALL LETTER O WITH GRAVE */
	K_oacute         Key = 0x00f3 /* U+00F3 LATIN SMALL LETTER O WITH ACUTE */
	K_ocircumflex    Key = 0x00f4 /* U+00F4 LATIN SMALL LETTER O WITH CIRCUMFLEX */
	K_otilde         Key = 0x00f5 /* U+00F5 LATIN SMALL LETTER O WITH TILDE */
	K_odiaeresis     Key = 0x00f6 /* U+00F6 LATIN SMALL LETTER O WITH DIAERESIS */
	K_division       Key = 0x00f7 /* U+00F7 DIVISION SIGN */
	K_oslash         Key = 0x00f8 /* U+00F8 LATIN SMALL LETTER O WITH STROKE */
	K_ooblique       Key = 0x00f8 /* deprecated alias for oslash */
	K_ugrave         Key = 0x00f9 /* U+00F9 LATIN SMALL LETTER U WITH GRAVE */
	K_uacute         Key = 0x00fa /* U+00FA LATIN SMALL LETTER U WITH ACUTE */
	K_ucircumflex    Key = 0x00fb /* U+00FB LATIN SMALL LETTER U WITH CIRCUMFLEX */
	K_udiaeresis     Key = 0x00fc /* U+00FC LATIN SMALL LETTER U WITH DIAERESIS */
	K_yacute         Key = 0x00fd /* U+00FD LATIN SMALL LETTER Y WITH ACUTE */
	K_thorn          Key = 0x00fe /* U+00FE LATIN SMALL LETTER THORN */
	K_ydiaeresis     Key = 0x00ff /* U+00FF LATIN SMALL LETTER Y WITH DIAERESIS */

	/*
	 * Latin 2
	 * Byte 3 = 1
	 */

	K_Aogonek      Key = 0x01a1 /* U+0104 LATIN CAPITAL LETTER A WITH OGONEK */
	K_breve        Key = 0x01a2 /* U+02D8 BREVE */
	K_Lstroke      Key = 0x01a3 /* U+0141 LATIN CAPITAL LETTER L WITH STROKE */
	K_Lcaron       Key = 0x01a5 /* U+013D LATIN CAPITAL LETTER L WITH CARON */
	K_Sacute       Key = 0x01a6 /* U+015A LATIN CAPITAL LETTER S WITH ACUTE */
	K_Scaron       Key = 0x01a9 /* U+0160 LATIN CAPITAL LETTER S WITH CARON */
	K_Scedilla     Key = 0x01aa /* U+015E LATIN CAPITAL LETTER S WITH CEDILLA */
	K_Tcaron       Key = 0x01ab /* U+0164 LATIN CAPITAL LETTER T WITH CARON */
	K_Zacute       Key = 0x01ac /* U+0179 LATIN CAPITAL LETTER Z WITH ACUTE */
	K_Zcaron       Key = 0x01ae /* U+017D LATIN CAPITAL LETTER Z WITH CARON */
	K_Zabovedot    Key = 0x01af /* U+017B LATIN CAPITAL LETTER Z WITH DOT ABOVE */
	K_aogonek      Key = 0x01b1 /* U+0105 LATIN SMALL LETTER A WITH OGONEK */
	K_ogonek       Key = 0x01b2 /* U+02DB OGONEK */
	K_lstroke      Key = 0x01b3 /* U+0142 LATIN SMALL LETTER L WITH STROKE */
	K_lcaron       Key = 0x01b5 /* U+013E LATIN SMALL LETTER L WITH CARON */
	K_sacute       Key = 0x01b6 /* U+015B LATIN SMALL LETTER S WITH ACUTE */
	K_caron        Key = 0x01b7 /* U+02C7 CARON */
	K_scaron       Key = 0x01b9 /* U+0161 LATIN SMALL LETTER S WITH CARON */
	K_scedilla     Key = 0x01ba /* U+015F LATIN SMALL LETTER S WITH CEDILLA */
	K_tcaron       Key = 0x01bb /* U+0165 LATIN SMALL LETTER T WITH CARON */
	K_zacute       Key = 0x01bc /* U+017A LATIN SMALL LETTER Z WITH ACUTE */
	K_doubleacute  Key = 0x01bd /* U+02DD DOUBLE ACUTE ACCENT */
	K_zcaron       Key = 0x01be /* U+017E LATIN SMALL LETTER Z WITH CARON */
	K_zabovedot    Key = 0x01bf /* U+017C LATIN SMALL LETTER Z WITH DOT ABOVE */
	K_Racute       Key = 0x01c0 /* U+0154 LATIN CAPITAL LETTER R WITH ACUTE */
	K_Abreve       Key = 0x01c3 /* U+0102 LATIN CAPITAL LETTER A WITH BREVE */
	K_Lacute       Key = 0x01c5 /* U+0139 LATIN CAPITAL LETTER L WITH ACUTE */
	K_Cacute       Key = 0x01c6 /* U+0106 LATIN CAPITAL LETTER C WITH ACUTE */
	K_Ccaron       Key = 0x01c8 /* U+010C LATIN CAPITAL LETTER C WITH CARON */
	K_Eogonek      Key = 0x01ca /* U+0118 LATIN CAPITAL LETTER E WITH OGONEK */
	K_Ecaron       Key = 0x01cc /* U+011A LATIN CAPITAL LETTER E WITH CARON */
	K_Dcaron       Key = 0x01cf /* U+010E LATIN CAPITAL LETTER D WITH CARON */
	K_Dstroke      Key = 0x01d0 /* U+0110 LATIN CAPITAL LETTER D WITH STROKE */
	K_Nacute       Key = 0x01d1 /* U+0143 LATIN CAPITAL LETTER N WITH ACUTE */
	K_Ncaron       Key = 0x01d2 /* U+0147 LATIN CAPITAL LETTER N WITH CARON */
	K_Odoubleacute Key = 0x01d5 /* U+0150 LATIN CAPITAL LETTER O WITH DOUBLE ACUTE */
	K_Rcaron       Key = 0x01d8 /* U+0158 LATIN CAPITAL LETTER R WITH CARON */
	K_Uring        Key = 0x01d9 /* U+016E LATIN CAPITAL LETTER U WITH RING ABOVE */
	K_Udoubleacute Key = 0x01db /* U+0170 LATIN CAPITAL LETTER U WITH DOUBLE ACUTE */
	K_Tcedilla     Key = 0x01de /* U+0162 LATIN CAPITAL LETTER T WITH CEDILLA */
	K_racute       Key = 0x01e0 /* U+0155 LATIN SMALL LETTER R WITH ACUTE */
	K_abreve       Key = 0x01e3 /* U+0103 LATIN SMALL LETTER A WITH BREVE */
	K_lacute       Key = 0x01e5 /* U+013A LATIN SMALL LETTER L WITH ACUTE */
	K_cacute       Key = 0x01e6 /* U+0107 LATIN SMALL LETTER C WITH ACUTE */
	K_ccaron       Key = 0x01e8 /* U+010D LATIN SMALL LETTER C WITH CARON */
	K_eogonek      Key = 0x01ea /* U+0119 LATIN SMALL LETTER E WITH OGONEK */
	K_ecaron       Key = 0x01ec /* U+011B LATIN SMALL LETTER E WITH CARON */
	K_dcaron       Key = 0x01ef /* U+010F LATIN SMALL LETTER D WITH CARON */
	K_dstroke      Key = 0x01f0 /* U+0111 LATIN SMALL LETTER D WITH STROKE */
	K_nacute       Key = 0x01f1 /* U+0144 LATIN SMALL LETTER N WITH ACUTE */
	K_ncaron       Key = 0x01f2 /* U+0148 LATIN SMALL LETTER N WITH CARON */
	K_odoubleacute Key = 0x01f5 /* U+0151 LATIN SMALL LETTER O WITH DOUBLE ACUTE */
	K_rcaron       Key = 0x01f8 /* U+0159 LATIN SMALL LETTER R WITH CARON */
	K_uring        Key = 0x01f9 /* U+016F LATIN SMALL LETTER U WITH RING ABOVE */
	K_udoubleacute Key = 0x01fb /* U+0171 LATIN SMALL LETTER U WITH DOUBLE ACUTE */
	K_tcedilla     Key = 0x01fe /* U+0163 LATIN SMALL LETTER T WITH CEDILLA */
	K_abovedot     Key = 0x01ff /* U+02D9 DOT ABOVE */

	/*
	 * Latin 3
	 * Byte 3 = 2
	 */

	K_Hstroke     Key = 0x02a1 /* U+0126 LATIN CAPITAL LETTER H WITH STROKE */
	K_Hcircumflex Key = 0x02a6 /* U+0124 LATIN CAPITAL LETTER H WITH CIRCUMFLEX */
	K_Iabovedot   Key = 0x02a9 /* U+0130 LATIN CAPITAL LETTER I WITH DOT ABOVE */
	K_Gbreve      Key = 0x02ab /* U+011E LATIN CAPITAL LETTER G WITH BREVE */
	K_Jcircumflex Key = 0x02ac /* U+0134 LATIN CAPITAL LETTER J WITH CIRCUMFLEX */
	K_hstroke     Key = 0x02b1 /* U+0127 LATIN SMALL LETTER H WITH STROKE */
	K_hcircumflex Key = 0x02b6 /* U+0125 LATIN SMALL LETTER H WITH CIRCUMFLEX */
	K_idotless    Key = 0x02b9 /* U+0131 LATIN SMALL LETTER DOTLESS I */
	K_gbreve      Key = 0x02bb /* U+011F LATIN SMALL LETTER G WITH BREVE */
	K_jcircumflex Key = 0x02bc /* U+0135 LATIN SMALL LETTER J WITH CIRCUMFLEX */
	K_Cabovedot   Key = 0x02c5 /* U+010A LATIN CAPITAL LETTER C WITH DOT ABOVE */
	K_Ccircumflex Key = 0x02c6 /* U+0108 LATIN CAPITAL LETTER C WITH CIRCUMFLEX */
	K_Gabovedot   Key = 0x02d5 /* U+0120 LATIN CAPITAL LETTER G WITH DOT ABOVE */
	K_Gcircumflex Key = 0x02d8 /* U+011C LATIN CAPITAL LETTER G WITH CIRCUMFLEX */
	K_Ubreve      Key = 0x02dd /* U+016C LATIN CAPITAL LETTER U WITH BREVE */
	K_Scircumflex Key = 0x02de /* U+015C LATIN CAPITAL LETTER S WITH CIRCUMFLEX */
	K_cabovedot   Key = 0x02e5 /* U+010B LATIN SMALL LETTER C WITH DOT ABOVE */
	K_ccircumflex Key = 0x02e6 /* U+0109 LATIN SMALL LETTER C WITH CIRCUMFLEX */
	K_gabovedot   Key = 0x02f5 /* U+0121 LATIN SMALL LETTER G WITH DOT ABOVE */
	K_gcircumflex Key = 0x02f8 /* U+011D LATIN SMALL LETTER G WITH CIRCUMFLEX */
	K_ubreve      Key = 0x02fd /* U+016D LATIN SMALL LETTER U WITH BREVE */
	K_scircumflex Key = 0x02fe /* U+015D LATIN SMALL LETTER S WITH CIRCUMFLEX */

	/*
	 * Latin 4
	 * Byte 3 = 3
	 */

	K_kra       Key = 0x03a2 /* U+0138 LATIN SMALL LETTER KRA */
	K_kappa     Key = 0x03a2 /* deprecated */
	K_Rcedilla  Key = 0x03a3 /* U+0156 LATIN CAPITAL LETTER R WITH CEDILLA */
	K_Itilde    Key = 0x03a5 /* U+0128 LATIN CAPITAL LETTER I WITH TILDE */
	K_Lcedilla  Key = 0x03a6 /* U+013B LATIN CAPITAL LETTER L WITH CEDILLA */
	K_Emacron   Key = 0x03aa /* U+0112 LATIN CAPITAL LETTER E WITH MACRON */
	K_Gcedilla  Key = 0x03ab /* U+0122 LATIN CAPITAL LETTER G WITH CEDILLA */
	K_Tslash    Key = 0x03ac /* U+0166 LATIN CAPITAL LETTER T WITH STROKE */
	K_rcedilla  Key = 0x03b3 /* U+0157 LATIN SMALL LETTER R WITH CEDILLA */
	K_itilde    Key = 0x03b5 /* U+0129 LATIN SMALL LETTER I WITH TILDE */
	K_lcedilla  Key = 0x03b6 /* U+013C LATIN SMALL LETTER L WITH CEDILLA */
	K_emacron   Key = 0x03ba /* U+0113 LATIN SMALL LETTER E WITH MACRON */
	K_gcedilla  Key = 0x03bb /* U+0123 LATIN SMALL LETTER G WITH CEDILLA */
	K_tslash    Key = 0x03bc /* U+0167 LATIN SMALL LETTER T WITH STROKE */
	K_ENG       Key = 0x03bd /* U+014A LATIN CAPITAL LETTER ENG */
	K_eng       Key = 0x03bf /* U+014B LATIN SMALL LETTER ENG */
	K_Amacron   Key = 0x03c0 /* U+0100 LATIN CAPITAL LETTER A WITH MACRON */
	K_Iogonek   Key = 0x03c7 /* U+012E LATIN CAPITAL LETTER I WITH OGONEK */
	K_Eabovedot Key = 0x03cc /* U+0116 LATIN CAPITAL LETTER E WITH DOT ABOVE */
	K_Imacron   Key = 0x03cf /* U+012A LATIN CAPITAL LETTER I WITH MACRON */
	K_Ncedilla  Key = 0x03d1 /* U+0145 LATIN CAPITAL LETTER N WITH CEDILLA */
	K_Omacron   Key = 0x03d2 /* U+014C LATIN CAPITAL LETTER O WITH MACRON */
	K_Kcedilla  Key = 0x03d3 /* U+0136 LATIN CAPITAL LETTER K WITH CEDILLA */
	K_Uogonek   Key = 0x03d9 /* U+0172 LATIN CAPITAL LETTER U WITH OGONEK */
	K_Utilde    Key = 0x03dd /* U+0168 LATIN CAPITAL LETTER U WITH TILDE */
	K_Umacron   Key = 0x03de /* U+016A LATIN CAPITAL LETTER U WITH MACRON */
	K_amacron   Key = 0x03e0 /* U+0101 LATIN SMALL LETTER A WITH MACRON */
	K_iogonek   Key = 0x03e7 /* U+012F LATIN SMALL LETTER I WITH OGONEK */
	K_eabovedot Key = 0x03ec /* U+0117 LATIN SMALL LETTER E WITH DOT ABOVE */
	K_imacron   Key = 0x03ef /* U+012B LATIN SMALL LETTER I WITH MACRON */
	K_ncedilla  Key = 0x03f1 /* U+0146 LATIN SMALL LETTER N WITH CEDILLA */
	K_omacron   Key = 0x03f2 /* U+014D LATIN SMALL LETTER O WITH MACRON */
	K_kcedilla  Key = 0x03f3 /* U+0137 LATIN SMALL LETTER K WITH CEDILLA */
	K_uogonek   Key = 0x03f9 /* U+0173 LATIN SMALL LETTER U WITH OGONEK */
	K_utilde    Key = 0x03fd /* U+0169 LATIN SMALL LETTER U WITH TILDE */
	K_umacron   Key = 0x03fe /* U+016B LATIN SMALL LETTER U WITH MACRON */

	/*
	 * Latin 8
	 */
	K_Wcircumflex Key = 0x1000174 /* U+0174 LATIN CAPITAL LETTER W WITH CIRCUMFLEX */
	K_wcircumflex Key = 0x1000175 /* U+0175 LATIN SMALL LETTER W WITH CIRCUMFLEX */
	K_Ycircumflex Key = 0x1000176 /* U+0176 LATIN CAPITAL LETTER Y WITH CIRCUMFLEX */
	K_ycircumflex Key = 0x1000177 /* U+0177 LATIN SMALL LETTER Y WITH CIRCUMFLEX */
	K_Babovedot   Key = 0x1001e02 /* U+1E02 LATIN CAPITAL LETTER B WITH DOT ABOVE */
	K_babovedot   Key = 0x1001e03 /* U+1E03 LATIN SMALL LETTER B WITH DOT ABOVE */
	K_Dabovedot   Key = 0x1001e0a /* U+1E0A LATIN CAPITAL LETTER D WITH DOT ABOVE */
	K_dabovedot   Key = 0x1001e0b /* U+1E0B LATIN SMALL LETTER D WITH DOT ABOVE */
	K_Fabovedot   Key = 0x1001e1e /* U+1E1E LATIN CAPITAL LETTER F WITH DOT ABOVE */
	K_fabovedot   Key = 0x1001e1f /* U+1E1F LATIN SMALL LETTER F WITH DOT ABOVE */
	K_Mabovedot   Key = 0x1001e40 /* U+1E40 LATIN CAPITAL LETTER M WITH DOT ABOVE */
	K_mabovedot   Key = 0x1001e41 /* U+1E41 LATIN SMALL LETTER M WITH DOT ABOVE */
	K_Pabovedot   Key = 0x1001e56 /* U+1E56 LATIN CAPITAL LETTER P WITH DOT ABOVE */
	K_pabovedot   Key = 0x1001e57 /* U+1E57 LATIN SMALL LETTER P WITH DOT ABOVE */
	K_Sabovedot   Key = 0x1001e60 /* U+1E60 LATIN CAPITAL LETTER S WITH DOT ABOVE */
	K_sabovedot   Key = 0x1001e61 /* U+1E61 LATIN SMALL LETTER S WITH DOT ABOVE */
	K_Tabovedot   Key = 0x1001e6a /* U+1E6A LATIN CAPITAL LETTER T WITH DOT ABOVE */
	K_tabovedot   Key = 0x1001e6b /* U+1E6B LATIN SMALL LETTER T WITH DOT ABOVE */
	K_Wgrave      Key = 0x1001e80 /* U+1E80 LATIN CAPITAL LETTER W WITH GRAVE */
	K_wgrave      Key = 0x1001e81 /* U+1E81 LATIN SMALL LETTER W WITH GRAVE */
	K_Wacute      Key = 0x1001e82 /* U+1E82 LATIN CAPITAL LETTER W WITH ACUTE */
	K_wacute      Key = 0x1001e83 /* U+1E83 LATIN SMALL LETTER W WITH ACUTE */
	K_Wdiaeresis  Key = 0x1001e84 /* U+1E84 LATIN CAPITAL LETTER W WITH DIAERESIS */
	K_wdiaeresis  Key = 0x1001e85 /* U+1E85 LATIN SMALL LETTER W WITH DIAERESIS */
	K_Ygrave      Key = 0x1001ef2 /* U+1EF2 LATIN CAPITAL LETTER Y WITH GRAVE */
	K_ygrave      Key = 0x1001ef3 /* U+1EF3 LATIN SMALL LETTER Y WITH GRAVE */

	/*
	 * Latin 9
	 * Byte 3 = 0x13
	 */

	K_OE         Key = 0x13bc /* U+0152 LATIN CAPITAL LIGATURE OE */
	K_oe         Key = 0x13bd /* U+0153 LATIN SMALL LIGATURE OE */
	K_Ydiaeresis Key = 0x13be /* U+0178 LATIN CAPITAL LETTER Y WITH DIAERESIS */

	/*
	 * Katakana
	 * Byte 3 = 4
	 */

	K_overline            Key = 0x047e /* U+203E OVERLINE */
	K_kana_fullstop       Key = 0x04a1 /* U+3002 IDEOGRAPHIC FULL STOP */
	K_kana_openingbracket Key = 0x04a2 /* U+300C LEFT CORNER BRACKET */
	K_kana_closingbracket Key = 0x04a3 /* U+300D RIGHT CORNER BRACKET */
	K_kana_comma          Key = 0x04a4 /* U+3001 IDEOGRAPHIC COMMA */
	K_kana_conjunctive    Key = 0x04a5 /* U+30FB KATAKANA MIDDLE DOT */
	K_kana_middledot      Key = 0x04a5 /* deprecated */
	K_kana_WO             Key = 0x04a6 /* U+30F2 KATAKANA LETTER WO */
	K_kana_a              Key = 0x04a7 /* U+30A1 KATAKANA LETTER SMALL A */
	K_kana_i              Key = 0x04a8 /* U+30A3 KATAKANA LETTER SMALL I */
	K_kana_u              Key = 0x04a9 /* U+30A5 KATAKANA LETTER SMALL U */
	K_kana_e              Key = 0x04aa /* U+30A7 KATAKANA LETTER SMALL E */
	K_kana_o              Key = 0x04ab /* U+30A9 KATAKANA LETTER SMALL O */
	K_kana_ya             Key = 0x04ac /* U+30E3 KATAKANA LETTER SMALL YA */
	K_kana_yu             Key = 0x04ad /* U+30E5 KATAKANA LETTER SMALL YU */
	K_kana_yo             Key = 0x04ae /* U+30E7 KATAKANA LETTER SMALL YO */
	K_kana_tsu            Key = 0x04af /* U+30C3 KATAKANA LETTER SMALL TU */
	K_kana_tu             Key = 0x04af /* deprecated */
	K_prolongedsound      Key = 0x04b0 /* U+30FC KATAKANA-HIRAGANA PROLONGED SOUND MARK */
	K_kana_A              Key = 0x04b1 /* U+30A2 KATAKANA LETTER A */
	K_kana_I              Key = 0x04b2 /* U+30A4 KATAKANA LETTER I */
	K_kana_U              Key = 0x04b3 /* U+30A6 KATAKANA LETTER U */
	K_kana_E              Key = 0x04b4 /* U+30A8 KATAKANA LETTER E */
	K_kana_O              Key = 0x04b5 /* U+30AA KATAKANA LETTER O */
	K_kana_KA             Key = 0x04b6 /* U+30AB KATAKANA LETTER KA */
	K_kana_KI             Key = 0x04b7 /* U+30AD KATAKANA LETTER KI */
	K_kana_KU             Key = 0x04b8 /* U+30AF KATAKANA LETTER KU */
	K_kana_KE             Key = 0x04b9 /* U+30B1 KATAKANA LETTER KE */
	K_kana_KO             Key = 0x04ba /* U+30B3 KATAKANA LETTER KO */
	K_kana_SA             Key = 0x04bb /* U+30B5 KATAKANA LETTER SA */
	K_kana_SHI            Key = 0x04bc /* U+30B7 KATAKANA LETTER SI */
	K_kana_SU             Key = 0x04bd /* U+30B9 KATAKANA LETTER SU */
	K_kana_SE             Key = 0x04be /* U+30BB KATAKANA LETTER SE */
	K_kana_SO             Key = 0x04bf /* U+30BD KATAKANA LETTER SO */
	K_kana_TA             Key = 0x04c0 /* U+30BF KATAKANA LETTER TA */
	K_kana_CHI            Key = 0x04c1 /* U+30C1 KATAKANA LETTER TI */
	K_kana_TI             Key = 0x04c1 /* deprecated */
	K_kana_TSU            Key = 0x04c2 /* U+30C4 KATAKANA LETTER TU */
	K_kana_TU             Key = 0x04c2 /* deprecated */
	K_kana_TE             Key = 0x04c3 /* U+30C6 KATAKANA LETTER TE */
	K_kana_TO             Key = 0x04c4 /* U+30C8 KATAKANA LETTER TO */
	K_kana_NA             Key = 0x04c5 /* U+30CA KATAKANA LETTER NA */
	K_kana_NI             Key = 0x04c6 /* U+30CB KATAKANA LETTER NI */
	K_kana_NU             Key = 0x04c7 /* U+30CC KATAKANA LETTER NU */
	K_kana_NE             Key = 0x04c8 /* U+30CD KATAKANA LETTER NE */
	K_kana_NO             Key = 0x04c9 /* U+30CE KATAKANA LETTER NO */
	K_kana_HA             Key = 0x04ca /* U+30CF KATAKANA LETTER HA */
	K_kana_HI             Key = 0x04cb /* U+30D2 KATAKANA LETTER HI */
	K_kana_FU             Key = 0x04cc /* U+30D5 KATAKANA LETTER HU */
	K_kana_HU             Key = 0x04cc /* deprecated */
	K_kana_HE             Key = 0x04cd /* U+30D8 KATAKANA LETTER HE */
	K_kana_HO             Key = 0x04ce /* U+30DB KATAKANA LETTER HO */
	K_kana_MA             Key = 0x04cf /* U+30DE KATAKANA LETTER MA */
	K_kana_MI             Key = 0x04d0 /* U+30DF KATAKANA LETTER MI */
	K_kana_MU             Key = 0x04d1 /* U+30E0 KATAKANA LETTER MU */
	K_kana_ME             Key = 0x04d2 /* U+30E1 KATAKANA LETTER ME */
	K_kana_MO             Key = 0x04d3 /* U+30E2 KATAKANA LETTER MO */
	K_kana_YA             Key = 0x04d4 /* U+30E4 KATAKANA LETTER YA */
	K_kana_YU             Key = 0x04d5 /* U+30E6 KATAKANA LETTER YU */
	K_kana_YO             Key = 0x04d6 /* U+30E8 KATAKANA LETTER YO */
	K_kana_RA             Key = 0x04d7 /* U+30E9 KATAKANA LETTER RA */
	K_kana_RI             Key = 0x04d8 /* U+30EA KATAKANA LETTER RI */
	K_kana_RU             Key = 0x04d9 /* U+30EB KATAKANA LETTER RU */
	K_kana_RE             Key = 0x04da /* U+30EC KATAKANA LETTER RE */
	K_kana_RO             Key = 0x04db /* U+30ED KATAKANA LETTER RO */
	K_kana_WA             Key = 0x04dc /* U+30EF KATAKANA LETTER WA */
	K_kana_N              Key = 0x04dd /* U+30F3 KATAKANA LETTER N */
	K_voicedsound         Key = 0x04de /* U+309B KATAKANA-HIRAGANA VOICED SOUND MARK */
	K_semivoicedsound     Key = 0x04df /* U+309C KATAKANA-HIRAGANA SEMI-VOICED SOUND MARK */
	K_kana_switch         Key = 0xff7e /* non-deprecated alias for Mode_switch */

	/*
	 * Arabic
	 * Byte 3 = 5
	 */

	K_Farsi_0                 Key = 0x10006f0 /* U+06F0 EXTENDED ARABIC-INDIC DIGIT ZERO */
	K_Farsi_1                 Key = 0x10006f1 /* U+06F1 EXTENDED ARABIC-INDIC DIGIT ONE */
	K_Farsi_2                 Key = 0x10006f2 /* U+06F2 EXTENDED ARABIC-INDIC DIGIT TWO */
	K_Farsi_3                 Key = 0x10006f3 /* U+06F3 EXTENDED ARABIC-INDIC DIGIT THREE */
	K_Farsi_4                 Key = 0x10006f4 /* U+06F4 EXTENDED ARABIC-INDIC DIGIT FOUR */
	K_Farsi_5                 Key = 0x10006f5 /* U+06F5 EXTENDED ARABIC-INDIC DIGIT FIVE */
	K_Farsi_6                 Key = 0x10006f6 /* U+06F6 EXTENDED ARABIC-INDIC DIGIT SIX */
	K_Farsi_7                 Key = 0x10006f7 /* U+06F7 EXTENDED ARABIC-INDIC DIGIT SEVEN */
	K_Farsi_8                 Key = 0x10006f8 /* U+06F8 EXTENDED ARABIC-INDIC DIGIT EIGHT */
	K_Farsi_9                 Key = 0x10006f9 /* U+06F9 EXTENDED ARABIC-INDIC DIGIT NINE */
	K_Arabic_percent          Key = 0x100066a /* U+066A ARABIC PERCENT SIGN */
	K_Arabic_superscript_alef Key = 0x1000670 /* U+0670 ARABIC LETTER SUPERSCRIPT ALEF */
	K_Arabic_tteh             Key = 0x1000679 /* U+0679 ARABIC LETTER TTEH */
	K_Arabic_peh              Key = 0x100067e /* U+067E ARABIC LETTER PEH */
	K_Arabic_tcheh            Key = 0x1000686 /* U+0686 ARABIC LETTER TCHEH */
	K_Arabic_ddal             Key = 0x1000688 /* U+0688 ARABIC LETTER DDAL */
	K_Arabic_rreh             Key = 0x1000691 /* U+0691 ARABIC LETTER RREH */
	K_Arabic_comma            Key = 0x05ac    /* U+060C ARABIC COMMA */
	K_Arabic_fullstop         Key = 0x10006d4 /* U+06D4 ARABIC FULL STOP */
	K_Arabic_0                Key = 0x1000660 /* U+0660 ARABIC-INDIC DIGIT ZERO */
	K_Arabic_1                Key = 0x1000661 /* U+0661 ARABIC-INDIC DIGIT ONE */
	K_Arabic_2                Key = 0x1000662 /* U+0662 ARABIC-INDIC DIGIT TWO */
	K_Arabic_3                Key = 0x1000663 /* U+0663 ARABIC-INDIC DIGIT THREE */
	K_Arabic_4                Key = 0x1000664 /* U+0664 ARABIC-INDIC DIGIT FOUR */
	K_Arabic_5                Key = 0x1000665 /* U+0665 ARABIC-INDIC DIGIT FIVE */
	K_Arabic_6                Key = 0x1000666 /* U+0666 ARABIC-INDIC DIGIT SIX */
	K_Arabic_7                Key = 0x1000667 /* U+0667 ARABIC-INDIC DIGIT SEVEN */
	K_Arabic_8                Key = 0x1000668 /* U+0668 ARABIC-INDIC DIGIT EIGHT */
	K_Arabic_9                Key = 0x1000669 /* U+0669 ARABIC-INDIC DIGIT NINE */
	K_Arabic_semicolon        Key = 0x05bb    /* U+061B ARABIC SEMICOLON */
	K_Arabic_question_mark    Key = 0x05bf    /* U+061F ARABIC QUESTION MARK */
	K_Arabic_hamza            Key = 0x05c1    /* U+0621 ARABIC LETTER HAMZA */
	K_Arabic_maddaonalef      Key = 0x05c2    /* U+0622 ARABIC LETTER ALEF WITH MADDA ABOVE */
	K_Arabic_hamzaonalef      Key = 0x05c3    /* U+0623 ARABIC LETTER ALEF WITH HAMZA ABOVE */
	K_Arabic_hamzaonwaw       Key = 0x05c4    /* U+0624 ARABIC LETTER WAW WITH HAMZA ABOVE */
	K_Arabic_hamzaunderalef   Key = 0x05c5    /* U+0625 ARABIC LETTER ALEF WITH HAMZA BELOW */
	K_Arabic_hamzaonyeh       Key = 0x05c6    /* U+0626 ARABIC LETTER YEH WITH HAMZA ABOVE */
	K_Arabic_alef             Key = 0x05c7    /* U+0627 ARABIC LETTER ALEF */
	K_Arabic_beh              Key = 0x05c8    /* U+0628 ARABIC LETTER BEH */
	K_Arabic_tehmarbuta       Key = 0x05c9    /* U+0629 ARABIC LETTER TEH MARBUTA */
	K_Arabic_teh              Key = 0x05ca    /* U+062A ARABIC LETTER TEH */
	K_Arabic_theh             Key = 0x05cb    /* U+062B ARABIC LETTER THEH */
	K_Arabic_jeem             Key = 0x05cc    /* U+062C ARABIC LETTER JEEM */
	K_Arabic_hah              Key = 0x05cd    /* U+062D ARABIC LETTER HAH */
	K_Arabic_khah             Key = 0x05ce    /* U+062E ARABIC LETTER KHAH */
	K_Arabic_dal              Key = 0x05cf    /* U+062F ARABIC LETTER DAL */
	K_Arabic_thal             Key = 0x05d0    /* U+0630 ARABIC LETTER THAL */
	K_Arabic_ra               Key = 0x05d1    /* U+0631 ARABIC LETTER REH */
	K_Arabic_zain             Key = 0x05d2    /* U+0632 ARABIC LETTER ZAIN */
	K_Arabic_seen             Key = 0x05d3    /* U+0633 ARABIC LETTER SEEN */
	K_Arabic_sheen            Key = 0x05d4    /* U+0634 ARABIC LETTER SHEEN */
	K_Arabic_sad              Key = 0x05d5    /* U+0635 ARABIC LETTER SAD */
	K_Arabic_dad              Key = 0x05d6    /* U+0636 ARABIC LETTER DAD */
	K_Arabic_tah              Key = 0x05d7    /* U+0637 ARABIC LETTER TAH */
	K_Arabic_zah              Key = 0x05d8    /* U+0638 ARABIC LETTER ZAH */
	K_Arabic_ain              Key = 0x05d9    /* U+0639 ARABIC LETTER AIN */
	K_Arabic_ghain            Key = 0x05da    /* U+063A ARABIC LETTER GHAIN */
	K_Arabic_tatweel          Key = 0x05e0    /* U+0640 ARABIC TATWEEL */
	K_Arabic_feh              Key = 0x05e1    /* U+0641 ARABIC LETTER FEH */
	K_Arabic_qaf              Key = 0x05e2    /* U+0642 ARABIC LETTER QAF */
	K_Arabic_kaf              Key = 0x05e3    /* U+0643 ARABIC LETTER KAF */
	K_Arabic_lam              Key = 0x05e4    /* U+0644 ARABIC LETTER LAM */
	K_Arabic_meem             Key = 0x05e5    /* U+0645 ARABIC LETTER MEEM */
	K_Arabic_noon             Key = 0x05e6    /* U+0646 ARABIC LETTER NOON */
	K_Arabic_ha               Key = 0x05e7    /* U+0647 ARABIC LETTER HEH */
	K_Arabic_heh              Key = 0x05e7    /* deprecated */
	K_Arabic_waw              Key = 0x05e8    /* U+0648 ARABIC LETTER WAW */
	K_Arabic_alefmaksura      Key = 0x05e9    /* U+0649 ARABIC LETTER ALEF MAKSURA */
	K_Arabic_yeh              Key = 0x05ea    /* U+064A ARABIC LETTER YEH */
	K_Arabic_fathatan         Key = 0x05eb    /* U+064B ARABIC FATHATAN */
	K_Arabic_dammatan         Key = 0x05ec    /* U+064C ARABIC DAMMATAN */
	K_Arabic_kasratan         Key = 0x05ed    /* U+064D ARABIC KASRATAN */
	K_Arabic_fatha            Key = 0x05ee    /* U+064E ARABIC FATHA */
	K_Arabic_damma            Key = 0x05ef    /* U+064F ARABIC DAMMA */
	K_Arabic_kasra            Key = 0x05f0    /* U+0650 ARABIC KASRA */
	K_Arabic_shadda           Key = 0x05f1    /* U+0651 ARABIC SHADDA */
	K_Arabic_sukun            Key = 0x05f2    /* U+0652 ARABIC SUKUN */
	K_Arabic_madda_above      Key = 0x1000653 /* U+0653 ARABIC MADDAH ABOVE */
	K_Arabic_hamza_above      Key = 0x1000654 /* U+0654 ARABIC HAMZA ABOVE */
	K_Arabic_hamza_below      Key = 0x1000655 /* U+0655 ARABIC HAMZA BELOW */
	K_Arabic_jeh              Key = 0x1000698 /* U+0698 ARABIC LETTER JEH */
	K_Arabic_veh              Key = 0x10006a4 /* U+06A4 ARABIC LETTER VEH */
	K_Arabic_keheh            Key = 0x10006a9 /* U+06A9 ARABIC LETTER KEHEH */
	K_Arabic_gaf              Key = 0x10006af /* U+06AF ARABIC LETTER GAF */
	K_Arabic_noon_ghunna      Key = 0x10006ba /* U+06BA ARABIC LETTER NOON GHUNNA */
	K_Arabic_heh_doachashmee  Key = 0x10006be /* U+06BE ARABIC LETTER HEH DOACHASHMEE */
	K_Farsi_yeh               Key = 0x10006cc /* U+06CC ARABIC LETTER FARSI YEH */
	K_Arabic_farsi_yeh        Key = 0x10006cc /* deprecated alias for Farsi_yeh */
	K_Arabic_yeh_baree        Key = 0x10006d2 /* U+06D2 ARABIC LETTER YEH BARREE */
	K_Arabic_heh_goal         Key = 0x10006c1 /* U+06C1 ARABIC LETTER HEH GOAL */
	K_Arabic_switch           Key = 0xff7e    /* non-deprecated alias for Mode_switch */

	/*
	 * Cyrillic
	 * Byte 3 = 6
	 */
	K_Cyrillic_GHE_bar        Key = 0x1000492 /* U+0492 CYRILLIC CAPITAL LETTER GHE WITH STROKE */
	K_Cyrillic_ghe_bar        Key = 0x1000493 /* U+0493 CYRILLIC SMALL LETTER GHE WITH STROKE */
	K_Cyrillic_ZHE_descender  Key = 0x1000496 /* U+0496 CYRILLIC CAPITAL LETTER ZHE WITH DESCENDER */
	K_Cyrillic_zhe_descender  Key = 0x1000497 /* U+0497 CYRILLIC SMALL LETTER ZHE WITH DESCENDER */
	K_Cyrillic_KA_descender   Key = 0x100049a /* U+049A CYRILLIC CAPITAL LETTER KA WITH DESCENDER */
	K_Cyrillic_ka_descender   Key = 0x100049b /* U+049B CYRILLIC SMALL LETTER KA WITH DESCENDER */
	K_Cyrillic_KA_vertstroke  Key = 0x100049c /* U+049C CYRILLIC CAPITAL LETTER KA WITH VERTICAL STROKE */
	K_Cyrillic_ka_vertstroke  Key = 0x100049d /* U+049D CYRILLIC SMALL LETTER KA WITH VERTICAL STROKE */
	K_Cyrillic_EN_descender   Key = 0x10004a2 /* U+04A2 CYRILLIC CAPITAL LETTER EN WITH DESCENDER */
	K_Cyrillic_en_descender   Key = 0x10004a3 /* U+04A3 CYRILLIC SMALL LETTER EN WITH DESCENDER */
	K_Cyrillic_U_straight     Key = 0x10004ae /* U+04AE CYRILLIC CAPITAL LETTER STRAIGHT U */
	K_Cyrillic_u_straight     Key = 0x10004af /* U+04AF CYRILLIC SMALL LETTER STRAIGHT U */
	K_Cyrillic_U_straight_bar Key = 0x10004b0 /* U+04B0 CYRILLIC CAPITAL LETTER STRAIGHT U WITH STROKE */
	K_Cyrillic_u_straight_bar Key = 0x10004b1 /* U+04B1 CYRILLIC SMALL LETTER STRAIGHT U WITH STROKE */
	K_Cyrillic_HA_descender   Key = 0x10004b2 /* U+04B2 CYRILLIC CAPITAL LETTER HA WITH DESCENDER */
	K_Cyrillic_ha_descender   Key = 0x10004b3 /* U+04B3 CYRILLIC SMALL LETTER HA WITH DESCENDER */
	K_Cyrillic_CHE_descender  Key = 0x10004b6 /* U+04B6 CYRILLIC CAPITAL LETTER CHE WITH DESCENDER */
	K_Cyrillic_che_descender  Key = 0x10004b7 /* U+04B7 CYRILLIC SMALL LETTER CHE WITH DESCENDER */
	K_Cyrillic_CHE_vertstroke Key = 0x10004b8 /* U+04B8 CYRILLIC CAPITAL LETTER CHE WITH VERTICAL STROKE */
	K_Cyrillic_che_vertstroke Key = 0x10004b9 /* U+04B9 CYRILLIC SMALL LETTER CHE WITH VERTICAL STROKE */
	K_Cyrillic_SHHA           Key = 0x10004ba /* U+04BA CYRILLIC CAPITAL LETTER SHHA */
	K_Cyrillic_shha           Key = 0x10004bb /* U+04BB CYRILLIC SMALL LETTER SHHA */

	K_Cyrillic_SCHWA    Key = 0x10004d8 /* U+04D8 CYRILLIC CAPITAL LETTER SCHWA */
	K_Cyrillic_schwa    Key = 0x10004d9 /* U+04D9 CYRILLIC SMALL LETTER SCHWA */
	K_Cyrillic_I_macron Key = 0x10004e2 /* U+04E2 CYRILLIC CAPITAL LETTER I WITH MACRON */
	K_Cyrillic_i_macron Key = 0x10004e3 /* U+04E3 CYRILLIC SMALL LETTER I WITH MACRON */
	K_Cyrillic_O_bar    Key = 0x10004e8 /* U+04E8 CYRILLIC CAPITAL LETTER BARRED O */
	K_Cyrillic_o_bar    Key = 0x10004e9 /* U+04E9 CYRILLIC SMALL LETTER BARRED O */
	K_Cyrillic_U_macron Key = 0x10004ee /* U+04EE CYRILLIC CAPITAL LETTER U WITH MACRON */
	K_Cyrillic_u_macron Key = 0x10004ef /* U+04EF CYRILLIC SMALL LETTER U WITH MACRON */

	K_Serbian_dje               Key = 0x06a1 /* U+0452 CYRILLIC SMALL LETTER DJE */
	K_Macedonia_gje             Key = 0x06a2 /* U+0453 CYRILLIC SMALL LETTER GJE */
	K_Cyrillic_io               Key = 0x06a3 /* U+0451 CYRILLIC SMALL LETTER IO */
	K_Ukrainian_ie              Key = 0x06a4 /* U+0454 CYRILLIC SMALL LETTER UKRAINIAN IE */
	K_Ukranian_je               Key = 0x06a4 /* deprecated */
	K_Macedonia_dse             Key = 0x06a5 /* U+0455 CYRILLIC SMALL LETTER DZE */
	K_Ukrainian_i               Key = 0x06a6 /* U+0456 CYRILLIC SMALL LETTER BYELORUSSIAN-UKRAINIAN I */
	K_Ukranian_i                Key = 0x06a6 /* deprecated */
	K_Ukrainian_yi              Key = 0x06a7 /* U+0457 CYRILLIC SMALL LETTER YI */
	K_Ukranian_yi               Key = 0x06a7 /* deprecated */
	K_Cyrillic_je               Key = 0x06a8 /* U+0458 CYRILLIC SMALL LETTER JE */
	K_Serbian_je                Key = 0x06a8 /* deprecated */
	K_Cyrillic_lje              Key = 0x06a9 /* U+0459 CYRILLIC SMALL LETTER LJE */
	K_Serbian_lje               Key = 0x06a9 /* deprecated */
	K_Cyrillic_nje              Key = 0x06aa /* U+045A CYRILLIC SMALL LETTER NJE */
	K_Serbian_nje               Key = 0x06aa /* deprecated */
	K_Serbian_tshe              Key = 0x06ab /* U+045B CYRILLIC SMALL LETTER TSHE */
	K_Macedonia_kje             Key = 0x06ac /* U+045C CYRILLIC SMALL LETTER KJE */
	K_Ukrainian_ghe_with_upturn Key = 0x06ad /* U+0491 CYRILLIC SMALL LETTER GHE WITH UPTURN */
	K_Byelorussian_shortu       Key = 0x06ae /* U+045E CYRILLIC SMALL LETTER SHORT U */
	K_Cyrillic_dzhe             Key = 0x06af /* U+045F CYRILLIC SMALL LETTER DZHE */
	K_Serbian_dze               Key = 0x06af /* deprecated */
	K_numerosign                Key = 0x06b0 /* U+2116 NUMERO SIGN */
	K_Serbian_DJE               Key = 0x06b1 /* U+0402 CYRILLIC CAPITAL LETTER DJE */
	K_Macedonia_GJE             Key = 0x06b2 /* U+0403 CYRILLIC CAPITAL LETTER GJE */
	K_Cyrillic_IO               Key = 0x06b3 /* U+0401 CYRILLIC CAPITAL LETTER IO */
	K_Ukrainian_IE              Key = 0x06b4 /* U+0404 CYRILLIC CAPITAL LETTER UKRAINIAN IE */
	K_Ukranian_JE               Key = 0x06b4 /* deprecated */
	K_Macedonia_DSE             Key = 0x06b5 /* U+0405 CYRILLIC CAPITAL LETTER DZE */
	K_Ukrainian_I               Key = 0x06b6 /* U+0406 CYRILLIC CAPITAL LETTER BYELORUSSIAN-UKRAINIAN I */
	K_Ukranian_I                Key = 0x06b6 /* deprecated */
	K_Ukrainian_YI              Key = 0x06b7 /* U+0407 CYRILLIC CAPITAL LETTER YI */
	K_Ukranian_YI               Key = 0x06b7 /* deprecated */
	K_Cyrillic_JE               Key = 0x06b8 /* U+0408 CYRILLIC CAPITAL LETTER JE */
	K_Serbian_JE                Key = 0x06b8 /* deprecated */
	K_Cyrillic_LJE              Key = 0x06b9 /* U+0409 CYRILLIC CAPITAL LETTER LJE */
	K_Serbian_LJE               Key = 0x06b9 /* deprecated */
	K_Cyrillic_NJE              Key = 0x06ba /* U+040A CYRILLIC CAPITAL LETTER NJE */
	K_Serbian_NJE               Key = 0x06ba /* deprecated */
	K_Serbian_TSHE              Key = 0x06bb /* U+040B CYRILLIC CAPITAL LETTER TSHE */
	K_Macedonia_KJE             Key = 0x06bc /* U+040C CYRILLIC CAPITAL LETTER KJE */
	K_Ukrainian_GHE_WITH_UPTURN Key = 0x06bd /* U+0490 CYRILLIC CAPITAL LETTER GHE WITH UPTURN */
	K_Byelorussian_SHORTU       Key = 0x06be /* U+040E CYRILLIC CAPITAL LETTER SHORT U */
	K_Cyrillic_DZHE             Key = 0x06bf /* U+040F CYRILLIC CAPITAL LETTER DZHE */
	K_Serbian_DZE               Key = 0x06bf /* deprecated */
	K_Cyrillic_yu               Key = 0x06c0 /* U+044E CYRILLIC SMALL LETTER YU */
	K_Cyrillic_a                Key = 0x06c1 /* U+0430 CYRILLIC SMALL LETTER A */
	K_Cyrillic_be               Key = 0x06c2 /* U+0431 CYRILLIC SMALL LETTER BE */
	K_Cyrillic_tse              Key = 0x06c3 /* U+0446 CYRILLIC SMALL LETTER TSE */
	K_Cyrillic_de               Key = 0x06c4 /* U+0434 CYRILLIC SMALL LETTER DE */
	K_Cyrillic_ie               Key = 0x06c5 /* U+0435 CYRILLIC SMALL LETTER IE */
	K_Cyrillic_ef               Key = 0x06c6 /* U+0444 CYRILLIC SMALL LETTER EF */
	K_Cyrillic_ghe              Key = 0x06c7 /* U+0433 CYRILLIC SMALL LETTER GHE */
	K_Cyrillic_ha               Key = 0x06c8 /* U+0445 CYRILLIC SMALL LETTER HA */
	K_Cyrillic_i                Key = 0x06c9 /* U+0438 CYRILLIC SMALL LETTER I */
	K_Cyrillic_shorti           Key = 0x06ca /* U+0439 CYRILLIC SMALL LETTER SHORT I */
	K_Cyrillic_ka               Key = 0x06cb /* U+043A CYRILLIC SMALL LETTER KA */
	K_Cyrillic_el               Key = 0x06cc /* U+043B CYRILLIC SMALL LETTER EL */
	K_Cyrillic_em               Key = 0x06cd /* U+043C CYRILLIC SMALL LETTER EM */
	K_Cyrillic_en               Key = 0x06ce /* U+043D CYRILLIC SMALL LETTER EN */
	K_Cyrillic_o                Key = 0x06cf /* U+043E CYRILLIC SMALL LETTER O */
	K_Cyrillic_pe               Key = 0x06d0 /* U+043F CYRILLIC SMALL LETTER PE */
	K_Cyrillic_ya               Key = 0x06d1 /* U+044F CYRILLIC SMALL LETTER YA */
	K_Cyrillic_er               Key = 0x06d2 /* U+0440 CYRILLIC SMALL LETTER ER */
	K_Cyrillic_es               Key = 0x06d3 /* U+0441 CYRILLIC SMALL LETTER ES */
	K_Cyrillic_te               Key = 0x06d4 /* U+0442 CYRILLIC SMALL LETTER TE */
	K_Cyrillic_u                Key = 0x06d5 /* U+0443 CYRILLIC SMALL LETTER U */
	K_Cyrillic_zhe              Key = 0x06d6 /* U+0436 CYRILLIC SMALL LETTER ZHE */
	K_Cyrillic_ve               Key = 0x06d7 /* U+0432 CYRILLIC SMALL LETTER VE */
	K_Cyrillic_softsign         Key = 0x06d8 /* U+044C CYRILLIC SMALL LETTER SOFT SIGN */
	K_Cyrillic_yeru             Key = 0x06d9 /* U+044B CYRILLIC SMALL LETTER YERU */
	K_Cyrillic_ze               Key = 0x06da /* U+0437 CYRILLIC SMALL LETTER ZE */
	K_Cyrillic_sha              Key = 0x06db /* U+0448 CYRILLIC SMALL LETTER SHA */
	K_Cyrillic_e                Key = 0x06dc /* U+044D CYRILLIC SMALL LETTER E */
	K_Cyrillic_shcha            Key = 0x06dd /* U+0449 CYRILLIC SMALL LETTER SHCHA */
	K_Cyrillic_che              Key = 0x06de /* U+0447 CYRILLIC SMALL LETTER CHE */
	K_Cyrillic_hardsign         Key = 0x06df /* U+044A CYRILLIC SMALL LETTER HARD SIGN */
	K_Cyrillic_YU               Key = 0x06e0 /* U+042E CYRILLIC CAPITAL LETTER YU */
	K_Cyrillic_A                Key = 0x06e1 /* U+0410 CYRILLIC CAPITAL LETTER A */
	K_Cyrillic_BE               Key = 0x06e2 /* U+0411 CYRILLIC CAPITAL LETTER BE */
	K_Cyrillic_TSE              Key = 0x06e3 /* U+0426 CYRILLIC CAPITAL LETTER TSE */
	K_Cyrillic_DE               Key = 0x06e4 /* U+0414 CYRILLIC CAPITAL LETTER DE */
	K_Cyrillic_IE               Key = 0x06e5 /* U+0415 CYRILLIC CAPITAL LETTER IE */
	K_Cyrillic_EF               Key = 0x06e6 /* U+0424 CYRILLIC CAPITAL LETTER EF */
	K_Cyrillic_GHE              Key = 0x06e7 /* U+0413 CYRILLIC CAPITAL LETTER GHE */
	K_Cyrillic_HA               Key = 0x06e8 /* U+0425 CYRILLIC CAPITAL LETTER HA */
	K_Cyrillic_I                Key = 0x06e9 /* U+0418 CYRILLIC CAPITAL LETTER I */
	K_Cyrillic_SHORTI           Key = 0x06ea /* U+0419 CYRILLIC CAPITAL LETTER SHORT I */
	K_Cyrillic_KA               Key = 0x06eb /* U+041A CYRILLIC CAPITAL LETTER KA */
	K_Cyrillic_EL               Key = 0x06ec /* U+041B CYRILLIC CAPITAL LETTER EL */
	K_Cyrillic_EM               Key = 0x06ed /* U+041C CYRILLIC CAPITAL LETTER EM */
	K_Cyrillic_EN               Key = 0x06ee /* U+041D CYRILLIC CAPITAL LETTER EN */
	K_Cyrillic_O                Key = 0x06ef /* U+041E CYRILLIC CAPITAL LETTER O */
	K_Cyrillic_PE               Key = 0x06f0 /* U+041F CYRILLIC CAPITAL LETTER PE */
	K_Cyrillic_YA               Key = 0x06f1 /* U+042F CYRILLIC CAPITAL LETTER YA */
	K_Cyrillic_ER               Key = 0x06f2 /* U+0420 CYRILLIC CAPITAL LETTER ER */
	K_Cyrillic_ES               Key = 0x06f3 /* U+0421 CYRILLIC CAPITAL LETTER ES */
	K_Cyrillic_TE               Key = 0x06f4 /* U+0422 CYRILLIC CAPITAL LETTER TE */
	K_Cyrillic_U                Key = 0x06f5 /* U+0423 CYRILLIC CAPITAL LETTER U */
	K_Cyrillic_ZHE              Key = 0x06f6 /* U+0416 CYRILLIC CAPITAL LETTER ZHE */
	K_Cyrillic_VE               Key = 0x06f7 /* U+0412 CYRILLIC CAPITAL LETTER VE */
	K_Cyrillic_SOFTSIGN         Key = 0x06f8 /* U+042C CYRILLIC CAPITAL LETTER SOFT SIGN */
	K_Cyrillic_YERU             Key = 0x06f9 /* U+042B CYRILLIC CAPITAL LETTER YERU */
	K_Cyrillic_ZE               Key = 0x06fa /* U+0417 CYRILLIC CAPITAL LETTER ZE */
	K_Cyrillic_SHA              Key = 0x06fb /* U+0428 CYRILLIC CAPITAL LETTER SHA */
	K_Cyrillic_E                Key = 0x06fc /* U+042D CYRILLIC CAPITAL LETTER E */
	K_Cyrillic_SHCHA            Key = 0x06fd /* U+0429 CYRILLIC CAPITAL LETTER SHCHA */
	K_Cyrillic_CHE              Key = 0x06fe /* U+0427 CYRILLIC CAPITAL LETTER CHE */
	K_Cyrillic_HARDSIGN         Key = 0x06ff /* U+042A CYRILLIC CAPITAL LETTER HARD SIGN */

	/*
	 * Greek
	 * (based on an early draft of, and not quite identical to, ISO/IEC 8859-7)
	 * Byte 3 = 7
	 */

	K_Greek_ALPHAaccent           Key = 0x07a1 /* U+0386 GREEK CAPITAL LETTER ALPHA WITH TONOS */
	K_Greek_EPSILONaccent         Key = 0x07a2 /* U+0388 GREEK CAPITAL LETTER EPSILON WITH TONOS */
	K_Greek_ETAaccent             Key = 0x07a3 /* U+0389 GREEK CAPITAL LETTER ETA WITH TONOS */
	K_Greek_IOTAaccent            Key = 0x07a4 /* U+038A GREEK CAPITAL LETTER IOTA WITH TONOS */
	K_Greek_IOTAdieresis          Key = 0x07a5 /* U+03AA GREEK CAPITAL LETTER IOTA WITH DIALYTIKA */
	K_Greek_IOTAdiaeresis         Key = 0x07a5 /* deprecated (old typo) */
	K_Greek_OMICRONaccent         Key = 0x07a7 /* U+038C GREEK CAPITAL LETTER OMICRON WITH TONOS */
	K_Greek_UPSILONaccent         Key = 0x07a8 /* U+038E GREEK CAPITAL LETTER UPSILON WITH TONOS */
	K_Greek_UPSILONdieresis       Key = 0x07a9 /* U+03AB GREEK CAPITAL LETTER UPSILON WITH DIALYTIKA */
	K_Greek_OMEGAaccent           Key = 0x07ab /* U+038F GREEK CAPITAL LETTER OMEGA WITH TONOS */
	K_Greek_accentdieresis        Key = 0x07ae /* U+0385 GREEK DIALYTIKA TONOS */
	K_Greek_horizbar              Key = 0x07af /* U+2015 HORIZONTAL BAR */
	K_Greek_alphaaccent           Key = 0x07b1 /* U+03AC GREEK SMALL LETTER ALPHA WITH TONOS */
	K_Greek_epsilonaccent         Key = 0x07b2 /* U+03AD GREEK SMALL LETTER EPSILON WITH TONOS */
	K_Greek_etaaccent             Key = 0x07b3 /* U+03AE GREEK SMALL LETTER ETA WITH TONOS */
	K_Greek_iotaaccent            Key = 0x07b4 /* U+03AF GREEK SMALL LETTER IOTA WITH TONOS */
	K_Greek_iotadieresis          Key = 0x07b5 /* U+03CA GREEK SMALL LETTER IOTA WITH DIALYTIKA */
	K_Greek_iotaaccentdieresis    Key = 0x07b6 /* U+0390 GREEK SMALL LETTER IOTA WITH DIALYTIKA AND TONOS */
	K_Greek_omicronaccent         Key = 0x07b7 /* U+03CC GREEK SMALL LETTER OMICRON WITH TONOS */
	K_Greek_upsilonaccent         Key = 0x07b8 /* U+03CD GREEK SMALL LETTER UPSILON WITH TONOS */
	K_Greek_upsilondieresis       Key = 0x07b9 /* U+03CB GREEK SMALL LETTER UPSILON WITH DIALYTIKA */
	K_Greek_upsilonaccentdieresis Key = 0x07ba /* U+03B0 GREEK SMALL LETTER UPSILON WITH DIALYTIKA AND TONOS */
	K_Greek_omegaaccent           Key = 0x07bb /* U+03CE GREEK SMALL LETTER OMEGA WITH TONOS */
	K_Greek_ALPHA                 Key = 0x07c1 /* U+0391 GREEK CAPITAL LETTER ALPHA */
	K_Greek_BETA                  Key = 0x07c2 /* U+0392 GREEK CAPITAL LETTER BETA */
	K_Greek_GAMMA                 Key = 0x07c3 /* U+0393 GREEK CAPITAL LETTER GAMMA */
	K_Greek_DELTA                 Key = 0x07c4 /* U+0394 GREEK CAPITAL LETTER DELTA */
	K_Greek_EPSILON               Key = 0x07c5 /* U+0395 GREEK CAPITAL LETTER EPSILON */
	K_Greek_ZETA                  Key = 0x07c6 /* U+0396 GREEK CAPITAL LETTER ZETA */
	K_Greek_ETA                   Key = 0x07c7 /* U+0397 GREEK CAPITAL LETTER ETA */
	K_Greek_THETA                 Key = 0x07c8 /* U+0398 GREEK CAPITAL LETTER THETA */
	K_Greek_IOTA                  Key = 0x07c9 /* U+0399 GREEK CAPITAL LETTER IOTA */
	K_Greek_KAPPA                 Key = 0x07ca /* U+039A GREEK CAPITAL LETTER KAPPA */
	K_Greek_LAMDA                 Key = 0x07cb /* U+039B GREEK CAPITAL LETTER LAMDA */
	K_Greek_LAMBDA                Key = 0x07cb /* non-deprecated alias for Greek_LAMDA */
	K_Greek_MU                    Key = 0x07cc /* U+039C GREEK CAPITAL LETTER MU */
	K_Greek_NU                    Key = 0x07cd /* U+039D GREEK CAPITAL LETTER NU */
	K_Greek_XI                    Key = 0x07ce /* U+039E GREEK CAPITAL LETTER XI */
	K_Greek_OMICRON               Key = 0x07cf /* U+039F GREEK CAPITAL LETTER OMICRON */
	K_Greek_PI                    Key = 0x07d0 /* U+03A0 GREEK CAPITAL LETTER PI */
	K_Greek_RHO                   Key = 0x07d1 /* U+03A1 GREEK CAPITAL LETTER RHO */
	K_Greek_SIGMA                 Key = 0x07d2 /* U+03A3 GREEK CAPITAL LETTER SIGMA */
	K_Greek_TAU                   Key = 0x07d4 /* U+03A4 GREEK CAPITAL LETTER TAU */
	K_Greek_UPSILON               Key = 0x07d5 /* U+03A5 GREEK CAPITAL LETTER UPSILON */
	K_Greek_PHI                   Key = 0x07d6 /* U+03A6 GREEK CAPITAL LETTER PHI */
	K_Greek_CHI                   Key = 0x07d7 /* U+03A7 GREEK CAPITAL LETTER CHI */
	K_Greek_PSI                   Key = 0x07d8 /* U+03A8 GREEK CAPITAL LETTER PSI */
	K_Greek_OMEGA                 Key = 0x07d9 /* U+03A9 GREEK CAPITAL LETTER OMEGA */
	K_Greek_alpha                 Key = 0x07e1 /* U+03B1 GREEK SMALL LETTER ALPHA */
	K_Greek_beta                  Key = 0x07e2 /* U+03B2 GREEK SMALL LETTER BETA */
	K_Greek_gamma                 Key = 0x07e3 /* U+03B3 GREEK SMALL LETTER GAMMA */
	K_Greek_delta                 Key = 0x07e4 /* U+03B4 GREEK SMALL LETTER DELTA */
	K_Greek_epsilon               Key = 0x07e5 /* U+03B5 GREEK SMALL LETTER EPSILON */
	K_Greek_zeta                  Key = 0x07e6 /* U+03B6 GREEK SMALL LETTER ZETA */
	K_Greek_eta                   Key = 0x07e7 /* U+03B7 GREEK SMALL LETTER ETA */
	K_Greek_theta                 Key = 0x07e8 /* U+03B8 GREEK SMALL LETTER THETA */
	K_Greek_iota                  Key = 0x07e9 /* U+03B9 GREEK SMALL LETTER IOTA */
	K_Greek_kappa                 Key = 0x07ea /* U+03BA GREEK SMALL LETTER KAPPA */
	K_Greek_lamda                 Key = 0x07eb /* U+03BB GREEK SMALL LETTER LAMDA */
	K_Greek_lambda                Key = 0x07eb /* non-deprecated alias for Greek_lamda */
	K_Greek_mu                    Key = 0x07ec /* U+03BC GREEK SMALL LETTER MU */
	K_Greek_nu                    Key = 0x07ed /* U+03BD GREEK SMALL LETTER NU */
	K_Greek_xi                    Key = 0x07ee /* U+03BE GREEK SMALL LETTER XI */
	K_Greek_omicron               Key = 0x07ef /* U+03BF GREEK SMALL LETTER OMICRON */
	K_Greek_pi                    Key = 0x07f0 /* U+03C0 GREEK SMALL LETTER PI */
	K_Greek_rho                   Key = 0x07f1 /* U+03C1 GREEK SMALL LETTER RHO */
	K_Greek_sigma                 Key = 0x07f2 /* U+03C3 GREEK SMALL LETTER SIGMA */
	K_Greek_finalsmallsigma       Key = 0x07f3 /* U+03C2 GREEK SMALL LETTER FINAL SIGMA */
	K_Greek_tau                   Key = 0x07f4 /* U+03C4 GREEK SMALL LETTER TAU */
	K_Greek_upsilon               Key = 0x07f5 /* U+03C5 GREEK SMALL LETTER UPSILON */
	K_Greek_phi                   Key = 0x07f6 /* U+03C6 GREEK SMALL LETTER PHI */
	K_Greek_chi                   Key = 0x07f7 /* U+03C7 GREEK SMALL LETTER CHI */
	K_Greek_psi                   Key = 0x07f8 /* U+03C8 GREEK SMALL LETTER PSI */
	K_Greek_omega                 Key = 0x07f9 /* U+03C9 GREEK SMALL LETTER OMEGA */
	K_Greek_switch                Key = 0xff7e /* non-deprecated alias for Mode_switch */

	/*
	 * Technical
	 * (from the DEC VT330/VT420 Technical Character Set, http://vt100.net/charsets/technical.html)
	 * Byte 3 = 8
	 */

	K_leftradical               Key = 0x08a1 /* U+23B7 RADICAL SYMBOL BOTTOM */
	K_topleftradical            Key = 0x08a2 /*(U+250C BOX DRAWINGS LIGHT DOWN AND RIGHT)*/
	K_horizconnector            Key = 0x08a3 /*(U+2500 BOX DRAWINGS LIGHT HORIZONTAL)*/
	K_topintegral               Key = 0x08a4 /* U+2320 TOP HALF INTEGRAL */
	K_botintegral               Key = 0x08a5 /* U+2321 BOTTOM HALF INTEGRAL */
	K_vertconnector             Key = 0x08a6 /*(U+2502 BOX DRAWINGS LIGHT VERTICAL)*/
	K_topleftsqbracket          Key = 0x08a7 /* U+23A1 LEFT SQUARE BRACKET UPPER CORNER */
	K_botleftsqbracket          Key = 0x08a8 /* U+23A3 LEFT SQUARE BRACKET LOWER CORNER */
	K_toprightsqbracket         Key = 0x08a9 /* U+23A4 RIGHT SQUARE BRACKET UPPER CORNER */
	K_botrightsqbracket         Key = 0x08aa /* U+23A6 RIGHT SQUARE BRACKET LOWER CORNER */
	K_topleftparens             Key = 0x08ab /* U+239B LEFT PARENTHESIS UPPER HOOK */
	K_botleftparens             Key = 0x08ac /* U+239D LEFT PARENTHESIS LOWER HOOK */
	K_toprightparens            Key = 0x08ad /* U+239E RIGHT PARENTHESIS UPPER HOOK */
	K_botrightparens            Key = 0x08ae /* U+23A0 RIGHT PARENTHESIS LOWER HOOK */
	K_leftmiddlecurlybrace      Key = 0x08af /* U+23A8 LEFT CURLY BRACKET MIDDLE PIECE */
	K_rightmiddlecurlybrace     Key = 0x08b0 /* U+23AC RIGHT CURLY BRACKET MIDDLE PIECE */
	K_topleftsummation          Key = 0x08b1
	K_botleftsummation          Key = 0x08b2
	K_topvertsummationconnector Key = 0x08b3
	K_botvertsummationconnector Key = 0x08b4
	K_toprightsummation         Key = 0x08b5
	K_botrightsummation         Key = 0x08b6
	K_rightmiddlesummation      Key = 0x08b7
	K_lessthanequal             Key = 0x08bc /* U+2264 LESS-THAN OR EQUAL TO */
	K_notequal                  Key = 0x08bd /* U+2260 NOT EQUAL TO */
	K_greaterthanequal          Key = 0x08be /* U+2265 GREATER-THAN OR EQUAL TO */
	K_integral                  Key = 0x08bf /* U+222B INTEGRAL */
	K_therefore                 Key = 0x08c0 /* U+2234 THEREFORE */
	K_variation                 Key = 0x08c1 /* U+221D PROPORTIONAL TO */
	K_infinity                  Key = 0x08c2 /* U+221E INFINITY */
	K_nabla                     Key = 0x08c5 /* U+2207 NABLA */
	K_approximate               Key = 0x08c8 /* U+223C TILDE OPERATOR */
	K_similarequal              Key = 0x08c9 /* U+2243 ASYMPTOTICALLY EQUAL TO */
	K_ifonlyif                  Key = 0x08cd /* U+21D4 LEFT RIGHT DOUBLE ARROW */
	K_implies                   Key = 0x08ce /* U+21D2 RIGHTWARDS DOUBLE ARROW */
	K_identical                 Key = 0x08cf /* U+2261 IDENTICAL TO */
	K_radical                   Key = 0x08d6 /* U+221A SQUARE ROOT */
	K_includedin                Key = 0x08da /* U+2282 SUBSET OF */
	K_includes                  Key = 0x08db /* U+2283 SUPERSET OF */
	K_intersection              Key = 0x08dc /* U+2229 INTERSECTION */
	K_union                     Key = 0x08dd /* U+222A UNION */
	K_logicaland                Key = 0x08de /* U+2227 LOGICAL AND */
	K_logicalor                 Key = 0x08df /* U+2228 LOGICAL OR */
	K_partialderivative         Key = 0x08ef /* U+2202 PARTIAL DIFFERENTIAL */
	K_function                  Key = 0x08f6 /* U+0192 LATIN SMALL LETTER F WITH HOOK */
	K_leftarrow                 Key = 0x08fb /* U+2190 LEFTWARDS ARROW */
	K_uparrow                   Key = 0x08fc /* U+2191 UPWARDS ARROW */
	K_rightarrow                Key = 0x08fd /* U+2192 RIGHTWARDS ARROW */
	K_downarrow                 Key = 0x08fe /* U+2193 DOWNWARDS ARROW */

	/*
	 * Special
	 * (from the DEC VT100 Special Graphics Character Set)
	 * Byte 3 = 9
	 */

	K_blank          Key = 0x09df
	K_soliddiamond   Key = 0x09e0 /* U+25C6 BLACK DIAMOND */
	K_checkerboard   Key = 0x09e1 /* U+2592 MEDIUM SHADE */
	K_ht             Key = 0x09e2 /* U+2409 SYMBOL FOR HORIZONTAL TABULATION */
	K_ff             Key = 0x09e3 /* U+240C SYMBOL FOR FORM FEED */
	K_cr             Key = 0x09e4 /* U+240D SYMBOL FOR CARRIAGE RETURN */
	K_lf             Key = 0x09e5 /* U+240A SYMBOL FOR LINE FEED */
	K_nl             Key = 0x09e8 /* U+2424 SYMBOL FOR NEWLINE */
	K_vt             Key = 0x09e9 /* U+240B SYMBOL FOR VERTICAL TABULATION */
	K_lowrightcorner Key = 0x09ea /* U+2518 BOX DRAWINGS LIGHT UP AND LEFT */
	K_uprightcorner  Key = 0x09eb /* U+2510 BOX DRAWINGS LIGHT DOWN AND LEFT */
	K_upleftcorner   Key = 0x09ec /* U+250C BOX DRAWINGS LIGHT DOWN AND RIGHT */
	K_lowleftcorner  Key = 0x09ed /* U+2514 BOX DRAWINGS LIGHT UP AND RIGHT */
	K_crossinglines  Key = 0x09ee /* U+253C BOX DRAWINGS LIGHT VERTICAL AND HORIZONTAL */
	K_horizlinescan1 Key = 0x09ef /* U+23BA HORIZONTAL SCAN LINE-1 */
	K_horizlinescan3 Key = 0x09f0 /* U+23BB HORIZONTAL SCAN LINE-3 */
	K_horizlinescan5 Key = 0x09f1 /* U+2500 BOX DRAWINGS LIGHT HORIZONTAL */
	K_horizlinescan7 Key = 0x09f2 /* U+23BC HORIZONTAL SCAN LINE-7 */
	K_horizlinescan9 Key = 0x09f3 /* U+23BD HORIZONTAL SCAN LINE-9 */
	K_leftt          Key = 0x09f4 /* U+251C BOX DRAWINGS LIGHT VERTICAL AND RIGHT */
	K_rightt         Key = 0x09f5 /* U+2524 BOX DRAWINGS LIGHT VERTICAL AND LEFT */
	K_bott           Key = 0x09f6 /* U+2534 BOX DRAWINGS LIGHT UP AND HORIZONTAL */
	K_topt           Key = 0x09f7 /* U+252C BOX DRAWINGS LIGHT DOWN AND HORIZONTAL */
	K_vertbar        Key = 0x09f8 /* U+2502 BOX DRAWINGS LIGHT VERTICAL */

	/*
	 * Publishing
	 * (these are probably from a long forgotten DEC Publishing
	 * font that once shipped with DECwrite)
	 * Byte 3 = 0x0a
	 */

	K_emspace              Key = 0x0aa1 /* U+2003 EM SPACE */
	K_enspace              Key = 0x0aa2 /* U+2002 EN SPACE */
	K_em3space             Key = 0x0aa3 /* U+2004 THREE-PER-EM SPACE */
	K_em4space             Key = 0x0aa4 /* U+2005 FOUR-PER-EM SPACE */
	K_digitspace           Key = 0x0aa5 /* U+2007 FIGURE SPACE */
	K_punctspace           Key = 0x0aa6 /* U+2008 PUNCTUATION SPACE */
	K_thinspace            Key = 0x0aa7 /* U+2009 THIN SPACE */
	K_hairspace            Key = 0x0aa8 /* U+200A HAIR SPACE */
	K_emdash               Key = 0x0aa9 /* U+2014 EM DASH */
	K_endash               Key = 0x0aaa /* U+2013 EN DASH */
	K_signifblank          Key = 0x0aac /*(U+2423 OPEN BOX)*/
	K_ellipsis             Key = 0x0aae /* U+2026 HORIZONTAL ELLIPSIS */
	K_doubbaselinedot      Key = 0x0aaf /* U+2025 TWO DOT LEADER */
	K_onethird             Key = 0x0ab0 /* U+2153 VULGAR FRACTION ONE THIRD */
	K_twothirds            Key = 0x0ab1 /* U+2154 VULGAR FRACTION TWO THIRDS */
	K_onefifth             Key = 0x0ab2 /* U+2155 VULGAR FRACTION ONE FIFTH */
	K_twofifths            Key = 0x0ab3 /* U+2156 VULGAR FRACTION TWO FIFTHS */
	K_threefifths          Key = 0x0ab4 /* U+2157 VULGAR FRACTION THREE FIFTHS */
	K_fourfifths           Key = 0x0ab5 /* U+2158 VULGAR FRACTION FOUR FIFTHS */
	K_onesixth             Key = 0x0ab6 /* U+2159 VULGAR FRACTION ONE SIXTH */
	K_fivesixths           Key = 0x0ab7 /* U+215A VULGAR FRACTION FIVE SIXTHS */
	K_careof               Key = 0x0ab8 /* U+2105 CARE OF */
	K_figdash              Key = 0x0abb /* U+2012 FIGURE DASH */
	K_leftanglebracket     Key = 0x0abc /*(U+2329 LEFT-POINTING ANGLE BRACKET)*/
	K_decimalpoint         Key = 0x0abd /*(U+002E FULL STOP)*/
	K_rightanglebracket    Key = 0x0abe /*(U+232A RIGHT-POINTING ANGLE BRACKET)*/
	K_marker               Key = 0x0abf
	K_oneeighth            Key = 0x0ac3 /* U+215B VULGAR FRACTION ONE EIGHTH */
	K_threeeighths         Key = 0x0ac4 /* U+215C VULGAR FRACTION THREE EIGHTHS */
	K_fiveeighths          Key = 0x0ac5 /* U+215D VULGAR FRACTION FIVE EIGHTHS */
	K_seveneighths         Key = 0x0ac6 /* U+215E VULGAR FRACTION SEVEN EIGHTHS */
	K_trademark            Key = 0x0ac9 /* U+2122 TRADE MARK SIGN */
	K_signaturemark        Key = 0x0aca /*(U+2613 SALTIRE)*/
	K_trademarkincircle    Key = 0x0acb
	K_leftopentriangle     Key = 0x0acc /*(U+25C1 WHITE LEFT-POINTING TRIANGLE)*/
	K_rightopentriangle    Key = 0x0acd /*(U+25B7 WHITE RIGHT-POINTING TRIANGLE)*/
	K_emopencircle         Key = 0x0ace /*(U+25CB WHITE CIRCLE)*/
	K_emopenrectangle      Key = 0x0acf /*(U+25AF WHITE VERTICAL RECTANGLE)*/
	K_leftsinglequotemark  Key = 0x0ad0 /* U+2018 LEFT SINGLE QUOTATION MARK */
	K_rightsinglequotemark Key = 0x0ad1 /* U+2019 RIGHT SINGLE QUOTATION MARK */
	K_leftdoublequotemark  Key = 0x0ad2 /* U+201C LEFT DOUBLE QUOTATION MARK */
	K_rightdoublequotemark Key = 0x0ad3 /* U+201D RIGHT DOUBLE QUOTATION MARK */
	K_prescription         Key = 0x0ad4 /* U+211E PRESCRIPTION TAKE */
	K_permille             Key = 0x0ad5 /* U+2030 PER MILLE SIGN */
	K_minutes              Key = 0x0ad6 /* U+2032 PRIME */
	K_seconds              Key = 0x0ad7 /* U+2033 DOUBLE PRIME */
	K_latincross           Key = 0x0ad9 /* U+271D LATIN CROSS */
	K_hexagram             Key = 0x0ada
	K_filledrectbullet     Key = 0x0adb /*(U+25AC BLACK RECTANGLE)*/
	K_filledlefttribullet  Key = 0x0adc /*(U+25C0 BLACK LEFT-POINTING TRIANGLE)*/
	K_filledrighttribullet Key = 0x0add /*(U+25B6 BLACK RIGHT-POINTING TRIANGLE)*/
	K_emfilledcircle       Key = 0x0ade /*(U+25CF BLACK CIRCLE)*/
	K_emfilledrect         Key = 0x0adf /*(U+25AE BLACK VERTICAL RECTANGLE)*/
	K_enopencircbullet     Key = 0x0ae0 /*(U+25E6 WHITE BULLET)*/
	K_enopensquarebullet   Key = 0x0ae1 /*(U+25AB WHITE SMALL SQUARE)*/
	K_openrectbullet       Key = 0x0ae2 /*(U+25AD WHITE RECTANGLE)*/
	K_opentribulletup      Key = 0x0ae3 /*(U+25B3 WHITE UP-POINTING TRIANGLE)*/
	K_opentribulletdown    Key = 0x0ae4 /*(U+25BD WHITE DOWN-POINTING TRIANGLE)*/
	K_openstar             Key = 0x0ae5 /*(U+2606 WHITE STAR)*/
	K_enfilledcircbullet   Key = 0x0ae6 /*(U+2022 BULLET)*/
	K_enfilledsqbullet     Key = 0x0ae7 /*(U+25AA BLACK SMALL SQUARE)*/
	K_filledtribulletup    Key = 0x0ae8 /*(U+25B2 BLACK UP-POINTING TRIANGLE)*/
	K_filledtribulletdown  Key = 0x0ae9 /*(U+25BC BLACK DOWN-POINTING TRIANGLE)*/
	K_leftpointer          Key = 0x0aea /*(U+261C WHITE LEFT POINTING INDEX)*/
	K_rightpointer         Key = 0x0aeb /*(U+261E WHITE RIGHT POINTING INDEX)*/
	K_club                 Key = 0x0aec /* U+2663 BLACK CLUB SUIT */
	K_diamond              Key = 0x0aed /* U+2666 BLACK DIAMOND SUIT */
	K_heart                Key = 0x0aee /* U+2665 BLACK HEART SUIT */
	K_maltesecross         Key = 0x0af0 /* U+2720 MALTESE CROSS */
	K_dagger               Key = 0x0af1 /* U+2020 DAGGER */
	K_doubledagger         Key = 0x0af2 /* U+2021 DOUBLE DAGGER */
	K_checkmark            Key = 0x0af3 /* U+2713 CHECK MARK */
	K_ballotcross          Key = 0x0af4 /* U+2717 BALLOT X */
	K_musicalsharp         Key = 0x0af5 /* U+266F MUSIC SHARP SIGN */
	K_musicalflat          Key = 0x0af6 /* U+266D MUSIC FLAT SIGN */
	K_malesymbol           Key = 0x0af7 /* U+2642 MALE SIGN */
	K_femalesymbol         Key = 0x0af8 /* U+2640 FEMALE SIGN */
	K_telephone            Key = 0x0af9 /* U+260E BLACK TELEPHONE */
	K_telephonerecorder    Key = 0x0afa /* U+2315 TELEPHONE RECORDER */
	K_phonographcopyright  Key = 0x0afb /* U+2117 SOUND RECORDING COPYRIGHT */
	K_caret                Key = 0x0afc /* U+2038 CARET */
	K_singlelowquotemark   Key = 0x0afd /* U+201A SINGLE LOW-9 QUOTATION MARK */
	K_doublelowquotemark   Key = 0x0afe /* U+201E DOUBLE LOW-9 QUOTATION MARK */
	K_cursor               Key = 0x0aff

	/*
	 * APL
	 * Byte 3 = 0x0b
	 */

	K_leftcaret  Key = 0x0ba3 /*(U+003C LESS-THAN SIGN)*/
	K_rightcaret Key = 0x0ba6 /*(U+003E GREATER-THAN SIGN)*/
	K_downcaret  Key = 0x0ba8 /*(U+2228 LOGICAL OR)*/
	K_upcaret    Key = 0x0ba9 /*(U+2227 LOGICAL AND)*/
	K_overbar    Key = 0x0bc0 /*(U+00AF MACRON)*/
	K_downtack   Key = 0x0bc2 /* U+22A4 DOWN TACK */
	K_upshoe     Key = 0x0bc3 /*(U+2229 INTERSECTION)*/
	K_downstile  Key = 0x0bc4 /* U+230A LEFT FLOOR */
	K_underbar   Key = 0x0bc6 /*(U+005F LOW LINE)*/
	K_jot        Key = 0x0bca /* U+2218 RING OPERATOR */
	K_quad       Key = 0x0bcc /* U+2395 APL FUNCTIONAL SYMBOL QUAD */
	K_uptack     Key = 0x0bce /* U+22A5 UP TACK */
	K_circle     Key = 0x0bcf /* U+25CB WHITE CIRCLE */
	K_upstile    Key = 0x0bd3 /* U+2308 LEFT CEILING */
	K_downshoe   Key = 0x0bd6 /*(U+222A UNION)*/
	K_rightshoe  Key = 0x0bd8 /*(U+2283 SUPERSET OF)*/
	K_leftshoe   Key = 0x0bda /*(U+2282 SUBSET OF)*/
	K_lefttack   Key = 0x0bdc /* U+22A3 LEFT TACK */
	K_righttack  Key = 0x0bfc /* U+22A2 RIGHT TACK */

	/*
	 * Hebrew
	 * Byte 3 = 0x0c
	 */

	K_hebrew_doublelowline Key = 0x0cdf /* U+2017 DOUBLE LOW LINE */
	K_hebrew_aleph         Key = 0x0ce0 /* U+05D0 HEBREW LETTER ALEF */
	K_hebrew_bet           Key = 0x0ce1 /* U+05D1 HEBREW LETTER BET */
	K_hebrew_beth          Key = 0x0ce1 /* deprecated */
	K_hebrew_gimel         Key = 0x0ce2 /* U+05D2 HEBREW LETTER GIMEL */
	K_hebrew_gimmel        Key = 0x0ce2 /* deprecated */
	K_hebrew_dalet         Key = 0x0ce3 /* U+05D3 HEBREW LETTER DALET */
	K_hebrew_daleth        Key = 0x0ce3 /* deprecated */
	K_hebrew_he            Key = 0x0ce4 /* U+05D4 HEBREW LETTER HE */
	K_hebrew_waw           Key = 0x0ce5 /* U+05D5 HEBREW LETTER VAV */
	K_hebrew_zain          Key = 0x0ce6 /* U+05D6 HEBREW LETTER ZAYIN */
	K_hebrew_zayin         Key = 0x0ce6 /* deprecated */
	K_hebrew_chet          Key = 0x0ce7 /* U+05D7 HEBREW LETTER HET */
	K_hebrew_het           Key = 0x0ce7 /* deprecated */
	K_hebrew_tet           Key = 0x0ce8 /* U+05D8 HEBREW LETTER TET */
	K_hebrew_teth          Key = 0x0ce8 /* deprecated */
	K_hebrew_yod           Key = 0x0ce9 /* U+05D9 HEBREW LETTER YOD */
	K_hebrew_finalkaph     Key = 0x0cea /* U+05DA HEBREW LETTER FINAL KAF */
	K_hebrew_kaph          Key = 0x0ceb /* U+05DB HEBREW LETTER KAF */
	K_hebrew_lamed         Key = 0x0cec /* U+05DC HEBREW LETTER LAMED */
	K_hebrew_finalmem      Key = 0x0ced /* U+05DD HEBREW LETTER FINAL MEM */
	K_hebrew_mem           Key = 0x0cee /* U+05DE HEBREW LETTER MEM */
	K_hebrew_finalnun      Key = 0x0cef /* U+05DF HEBREW LETTER FINAL NUN */
	K_hebrew_nun           Key = 0x0cf0 /* U+05E0 HEBREW LETTER NUN */
	K_hebrew_samech        Key = 0x0cf1 /* U+05E1 HEBREW LETTER SAMEKH */
	K_hebrew_samekh        Key = 0x0cf1 /* deprecated */
	K_hebrew_ayin          Key = 0x0cf2 /* U+05E2 HEBREW LETTER AYIN */
	K_hebrew_finalpe       Key = 0x0cf3 /* U+05E3 HEBREW LETTER FINAL PE */
	K_hebrew_pe            Key = 0x0cf4 /* U+05E4 HEBREW LETTER PE */
	K_hebrew_finalzade     Key = 0x0cf5 /* U+05E5 HEBREW LETTER FINAL TSADI */
	K_hebrew_finalzadi     Key = 0x0cf5 /* deprecated */
	K_hebrew_zade          Key = 0x0cf6 /* U+05E6 HEBREW LETTER TSADI */
	K_hebrew_zadi          Key = 0x0cf6 /* deprecated */
	K_hebrew_qoph          Key = 0x0cf7 /* U+05E7 HEBREW LETTER QOF */
	K_hebrew_kuf           Key = 0x0cf7 /* deprecated */
	K_hebrew_resh          Key = 0x0cf8 /* U+05E8 HEBREW LETTER RESH */
	K_hebrew_shin          Key = 0x0cf9 /* U+05E9 HEBREW LETTER SHIN */
	K_hebrew_taw           Key = 0x0cfa /* U+05EA HEBREW LETTER TAV */
	K_hebrew_taf           Key = 0x0cfa /* deprecated */
	K_Hebrew_switch        Key = 0xff7e /* non-deprecated alias for Mode_switch */

	/*
	 * Thai
	 * Byte 3 = 0x0d
	 */

	K_Thai_kokai             Key = 0x0da1 /* U+0E01 THAI CHARACTER KO KAI */
	K_Thai_khokhai           Key = 0x0da2 /* U+0E02 THAI CHARACTER KHO KHAI */
	K_Thai_khokhuat          Key = 0x0da3 /* U+0E03 THAI CHARACTER KHO KHUAT */
	K_Thai_khokhwai          Key = 0x0da4 /* U+0E04 THAI CHARACTER KHO KHWAI */
	K_Thai_khokhon           Key = 0x0da5 /* U+0E05 THAI CHARACTER KHO KHON */
	K_Thai_khorakhang        Key = 0x0da6 /* U+0E06 THAI CHARACTER KHO RAKHANG */
	K_Thai_ngongu            Key = 0x0da7 /* U+0E07 THAI CHARACTER NGO NGU */
	K_Thai_chochan           Key = 0x0da8 /* U+0E08 THAI CHARACTER CHO CHAN */
	K_Thai_choching          Key = 0x0da9 /* U+0E09 THAI CHARACTER CHO CHING */
	K_Thai_chochang          Key = 0x0daa /* U+0E0A THAI CHARACTER CHO CHANG */
	K_Thai_soso              Key = 0x0dab /* U+0E0B THAI CHARACTER SO SO */
	K_Thai_chochoe           Key = 0x0dac /* U+0E0C THAI CHARACTER CHO CHOE */
	K_Thai_yoying            Key = 0x0dad /* U+0E0D THAI CHARACTER YO YING */
	K_Thai_dochada           Key = 0x0dae /* U+0E0E THAI CHARACTER DO CHADA */
	K_Thai_topatak           Key = 0x0daf /* U+0E0F THAI CHARACTER TO PATAK */
	K_Thai_thothan           Key = 0x0db0 /* U+0E10 THAI CHARACTER THO THAN */
	K_Thai_thonangmontho     Key = 0x0db1 /* U+0E11 THAI CHARACTER THO NANGMONTHO */
	K_Thai_thophuthao        Key = 0x0db2 /* U+0E12 THAI CHARACTER THO PHUTHAO */
	K_Thai_nonen             Key = 0x0db3 /* U+0E13 THAI CHARACTER NO NEN */
	K_Thai_dodek             Key = 0x0db4 /* U+0E14 THAI CHARACTER DO DEK */
	K_Thai_totao             Key = 0x0db5 /* U+0E15 THAI CHARACTER TO TAO */
	K_Thai_thothung          Key = 0x0db6 /* U+0E16 THAI CHARACTER THO THUNG */
	K_Thai_thothahan         Key = 0x0db7 /* U+0E17 THAI CHARACTER THO THAHAN */
	K_Thai_thothong          Key = 0x0db8 /* U+0E18 THAI CHARACTER THO THONG */
	K_Thai_nonu              Key = 0x0db9 /* U+0E19 THAI CHARACTER NO NU */
	K_Thai_bobaimai          Key = 0x0dba /* U+0E1A THAI CHARACTER BO BAIMAI */
	K_Thai_popla             Key = 0x0dbb /* U+0E1B THAI CHARACTER PO PLA */
	K_Thai_phophung          Key = 0x0dbc /* U+0E1C THAI CHARACTER PHO PHUNG */
	K_Thai_fofa              Key = 0x0dbd /* U+0E1D THAI CHARACTER FO FA */
	K_Thai_phophan           Key = 0x0dbe /* U+0E1E THAI CHARACTER PHO PHAN */
	K_Thai_fofan             Key = 0x0dbf /* U+0E1F THAI CHARACTER FO FAN */
	K_Thai_phosamphao        Key = 0x0dc0 /* U+0E20 THAI CHARACTER PHO SAMPHAO */
	K_Thai_moma              Key = 0x0dc1 /* U+0E21 THAI CHARACTER MO MA */
	K_Thai_yoyak             Key = 0x0dc2 /* U+0E22 THAI CHARACTER YO YAK */
	K_Thai_rorua             Key = 0x0dc3 /* U+0E23 THAI CHARACTER RO RUA */
	K_Thai_ru                Key = 0x0dc4 /* U+0E24 THAI CHARACTER RU */
	K_Thai_loling            Key = 0x0dc5 /* U+0E25 THAI CHARACTER LO LING */
	K_Thai_lu                Key = 0x0dc6 /* U+0E26 THAI CHARACTER LU */
	K_Thai_wowaen            Key = 0x0dc7 /* U+0E27 THAI CHARACTER WO WAEN */
	K_Thai_sosala            Key = 0x0dc8 /* U+0E28 THAI CHARACTER SO SALA */
	K_Thai_sorusi            Key = 0x0dc9 /* U+0E29 THAI CHARACTER SO RUSI */
	K_Thai_sosua             Key = 0x0dca /* U+0E2A THAI CHARACTER SO SUA */
	K_Thai_hohip             Key = 0x0dcb /* U+0E2B THAI CHARACTER HO HIP */
	K_Thai_lochula           Key = 0x0dcc /* U+0E2C THAI CHARACTER LO CHULA */
	K_Thai_oang              Key = 0x0dcd /* U+0E2D THAI CHARACTER O ANG */
	K_Thai_honokhuk          Key = 0x0dce /* U+0E2E THAI CHARACTER HO NOKHUK */
	K_Thai_paiyannoi         Key = 0x0dcf /* U+0E2F THAI CHARACTER PAIYANNOI */
	K_Thai_saraa             Key = 0x0dd0 /* U+0E30 THAI CHARACTER SARA A */
	K_Thai_maihanakat        Key = 0x0dd1 /* U+0E31 THAI CHARACTER MAI HAN-AKAT */
	K_Thai_saraaa            Key = 0x0dd2 /* U+0E32 THAI CHARACTER SARA AA */
	K_Thai_saraam            Key = 0x0dd3 /* U+0E33 THAI CHARACTER SARA AM */
	K_Thai_sarai             Key = 0x0dd4 /* U+0E34 THAI CHARACTER SARA I */
	K_Thai_saraii            Key = 0x0dd5 /* U+0E35 THAI CHARACTER SARA II */
	K_Thai_saraue            Key = 0x0dd6 /* U+0E36 THAI CHARACTER SARA UE */
	K_Thai_sarauee           Key = 0x0dd7 /* U+0E37 THAI CHARACTER SARA UEE */
	K_Thai_sarau             Key = 0x0dd8 /* U+0E38 THAI CHARACTER SARA U */
	K_Thai_sarauu            Key = 0x0dd9 /* U+0E39 THAI CHARACTER SARA UU */
	K_Thai_phinthu           Key = 0x0dda /* U+0E3A THAI CHARACTER PHINTHU */
	K_Thai_maihanakat_maitho Key = 0x0dde /*(U+0E3E Unassigned code point)*/
	K_Thai_baht              Key = 0x0ddf /* U+0E3F THAI CURRENCY SYMBOL BAHT */
	K_Thai_sarae             Key = 0x0de0 /* U+0E40 THAI CHARACTER SARA E */
	K_Thai_saraae            Key = 0x0de1 /* U+0E41 THAI CHARACTER SARA AE */
	K_Thai_sarao             Key = 0x0de2 /* U+0E42 THAI CHARACTER SARA O */
	K_Thai_saraaimaimuan     Key = 0x0de3 /* U+0E43 THAI CHARACTER SARA AI MAIMUAN */
	K_Thai_saraaimaimalai    Key = 0x0de4 /* U+0E44 THAI CHARACTER SARA AI MAIMALAI */
	K_Thai_lakkhangyao       Key = 0x0de5 /* U+0E45 THAI CHARACTER LAKKHANGYAO */
	K_Thai_maiyamok          Key = 0x0de6 /* U+0E46 THAI CHARACTER MAIYAMOK */
	K_Thai_maitaikhu         Key = 0x0de7 /* U+0E47 THAI CHARACTER MAITAIKHU */
	K_Thai_maiek             Key = 0x0de8 /* U+0E48 THAI CHARACTER MAI EK */
	K_Thai_maitho            Key = 0x0de9 /* U+0E49 THAI CHARACTER MAI THO */
	K_Thai_maitri            Key = 0x0dea /* U+0E4A THAI CHARACTER MAI TRI */
	K_Thai_maichattawa       Key = 0x0deb /* U+0E4B THAI CHARACTER MAI CHATTAWA */
	K_Thai_thanthakhat       Key = 0x0dec /* U+0E4C THAI CHARACTER THANTHAKHAT */
	K_Thai_nikhahit          Key = 0x0ded /* U+0E4D THAI CHARACTER NIKHAHIT */
	K_Thai_leksun            Key = 0x0df0 /* U+0E50 THAI DIGIT ZERO */
	K_Thai_leknung           Key = 0x0df1 /* U+0E51 THAI DIGIT ONE */
	K_Thai_leksong           Key = 0x0df2 /* U+0E52 THAI DIGIT TWO */
	K_Thai_leksam            Key = 0x0df3 /* U+0E53 THAI DIGIT THREE */
	K_Thai_leksi             Key = 0x0df4 /* U+0E54 THAI DIGIT FOUR */
	K_Thai_lekha             Key = 0x0df5 /* U+0E55 THAI DIGIT FIVE */
	K_Thai_lekhok            Key = 0x0df6 /* U+0E56 THAI DIGIT SIX */
	K_Thai_lekchet           Key = 0x0df7 /* U+0E57 THAI DIGIT SEVEN */
	K_Thai_lekpaet           Key = 0x0df8 /* U+0E58 THAI DIGIT EIGHT */
	K_Thai_lekkao            Key = 0x0df9 /* U+0E59 THAI DIGIT NINE */

	/*
	 * Korean
	 * Byte 3 = 0x0e
	 */

	K_Hangul                   Key = 0xff31 /* Hangul start/stop(toggle) */
	K_Hangul_Start             Key = 0xff32 /* Hangul start */
	K_Hangul_End               Key = 0xff33 /* Hangul end, English start */
	K_Hangul_Hanja             Key = 0xff34 /* Start Hangul->Hanja Conversion */
	K_Hangul_Jamo              Key = 0xff35 /* Hangul Jamo mode */
	K_Hangul_Romaja            Key = 0xff36 /* Hangul Romaja mode */
	K_Hangul_Codeinput         Key = 0xff37 /* Hangul code input mode */
	K_Hangul_Jeonja            Key = 0xff38 /* Jeonja mode */
	K_Hangul_Banja             Key = 0xff39 /* Banja mode */
	K_Hangul_PreHanja          Key = 0xff3a /* Pre Hanja conversion */
	K_Hangul_PostHanja         Key = 0xff3b /* Post Hanja conversion */
	K_Hangul_SingleCandidate   Key = 0xff3c /* Single candidate */
	K_Hangul_MultipleCandidate Key = 0xff3d /* Multiple candidate */
	K_Hangul_PreviousCandidate Key = 0xff3e /* Previous candidate */
	K_Hangul_Special           Key = 0xff3f /* Special symbols */
	K_Hangul_switch            Key = 0xff7e /* non-deprecated alias for Mode_switch */

	/* Hangul Consonant Characters */
	K_Hangul_Kiyeog      Key = 0x0ea1 /* U+3131 HANGUL LETTER KIYEOK */
	K_Hangul_SsangKiyeog Key = 0x0ea2 /* U+3132 HANGUL LETTER SSANGKIYEOK */
	K_Hangul_KiyeogSios  Key = 0x0ea3 /* U+3133 HANGUL LETTER KIYEOK-SIOS */
	K_Hangul_Nieun       Key = 0x0ea4 /* U+3134 HANGUL LETTER NIEUN */
	K_Hangul_NieunJieuj  Key = 0x0ea5 /* U+3135 HANGUL LETTER NIEUN-CIEUC */
	K_Hangul_NieunHieuh  Key = 0x0ea6 /* U+3136 HANGUL LETTER NIEUN-HIEUH */
	K_Hangul_Dikeud      Key = 0x0ea7 /* U+3137 HANGUL LETTER TIKEUT */
	K_Hangul_SsangDikeud Key = 0x0ea8 /* U+3138 HANGUL LETTER SSANGTIKEUT */
	K_Hangul_Rieul       Key = 0x0ea9 /* U+3139 HANGUL LETTER RIEUL */
	K_Hangul_RieulKiyeog Key = 0x0eaa /* U+313A HANGUL LETTER RIEUL-KIYEOK */
	K_Hangul_RieulMieum  Key = 0x0eab /* U+313B HANGUL LETTER RIEUL-MIEUM */
	K_Hangul_RieulPieub  Key = 0x0eac /* U+313C HANGUL LETTER RIEUL-PIEUP */
	K_Hangul_RieulSios   Key = 0x0ead /* U+313D HANGUL LETTER RIEUL-SIOS */
	K_Hangul_RieulTieut  Key = 0x0eae /* U+313E HANGUL LETTER RIEUL-THIEUTH */
	K_Hangul_RieulPhieuf Key = 0x0eaf /* U+313F HANGUL LETTER RIEUL-PHIEUPH */
	K_Hangul_RieulHieuh  Key = 0x0eb0 /* U+3140 HANGUL LETTER RIEUL-HIEUH */
	K_Hangul_Mieum       Key = 0x0eb1 /* U+3141 HANGUL LETTER MIEUM */
	K_Hangul_Pieub       Key = 0x0eb2 /* U+3142 HANGUL LETTER PIEUP */
	K_Hangul_SsangPieub  Key = 0x0eb3 /* U+3143 HANGUL LETTER SSANGPIEUP */
	K_Hangul_PieubSios   Key = 0x0eb4 /* U+3144 HANGUL LETTER PIEUP-SIOS */
	K_Hangul_Sios        Key = 0x0eb5 /* U+3145 HANGUL LETTER SIOS */
	K_Hangul_SsangSios   Key = 0x0eb6 /* U+3146 HANGUL LETTER SSANGSIOS */
	K_Hangul_Ieung       Key = 0x0eb7 /* U+3147 HANGUL LETTER IEUNG */
	K_Hangul_Jieuj       Key = 0x0eb8 /* U+3148 HANGUL LETTER CIEUC */
	K_Hangul_SsangJieuj  Key = 0x0eb9 /* U+3149 HANGUL LETTER SSANGCIEUC */
	K_Hangul_Cieuc       Key = 0x0eba /* U+314A HANGUL LETTER CHIEUCH */
	K_Hangul_Khieuq      Key = 0x0ebb /* U+314B HANGUL LETTER KHIEUKH */
	K_Hangul_Tieut       Key = 0x0ebc /* U+314C HANGUL LETTER THIEUTH */
	K_Hangul_Phieuf      Key = 0x0ebd /* U+314D HANGUL LETTER PHIEUPH */
	K_Hangul_Hieuh       Key = 0x0ebe /* U+314E HANGUL LETTER HIEUH */

	/* Hangul Vowel Characters */
	K_Hangul_A   Key = 0x0ebf /* U+314F HANGUL LETTER A */
	K_Hangul_AE  Key = 0x0ec0 /* U+3150 HANGUL LETTER AE */
	K_Hangul_YA  Key = 0x0ec1 /* U+3151 HANGUL LETTER YA */
	K_Hangul_YAE Key = 0x0ec2 /* U+3152 HANGUL LETTER YAE */
	K_Hangul_EO  Key = 0x0ec3 /* U+3153 HANGUL LETTER EO */
	K_Hangul_E   Key = 0x0ec4 /* U+3154 HANGUL LETTER E */
	K_Hangul_YEO Key = 0x0ec5 /* U+3155 HANGUL LETTER YEO */
	K_Hangul_YE  Key = 0x0ec6 /* U+3156 HANGUL LETTER YE */
	K_Hangul_O   Key = 0x0ec7 /* U+3157 HANGUL LETTER O */
	K_Hangul_WA  Key = 0x0ec8 /* U+3158 HANGUL LETTER WA */
	K_Hangul_WAE Key = 0x0ec9 /* U+3159 HANGUL LETTER WAE */
	K_Hangul_OE  Key = 0x0eca /* U+315A HANGUL LETTER OE */
	K_Hangul_YO  Key = 0x0ecb /* U+315B HANGUL LETTER YO */
	K_Hangul_U   Key = 0x0ecc /* U+315C HANGUL LETTER U */
	K_Hangul_WEO Key = 0x0ecd /* U+315D HANGUL LETTER WEO */
	K_Hangul_WE  Key = 0x0ece /* U+315E HANGUL LETTER WE */
	K_Hangul_WI  Key = 0x0ecf /* U+315F HANGUL LETTER WI */
	K_Hangul_YU  Key = 0x0ed0 /* U+3160 HANGUL LETTER YU */
	K_Hangul_EU  Key = 0x0ed1 /* U+3161 HANGUL LETTER EU */
	K_Hangul_YI  Key = 0x0ed2 /* U+3162 HANGUL LETTER YI */
	K_Hangul_I   Key = 0x0ed3 /* U+3163 HANGUL LETTER I */

	/* Hangul syllable-final (JongSeong) Characters */
	K_Hangul_J_Kiyeog      Key = 0x0ed4 /* U+11A8 HANGUL JONGSEONG KIYEOK */
	K_Hangul_J_SsangKiyeog Key = 0x0ed5 /* U+11A9 HANGUL JONGSEONG SSANGKIYEOK */
	K_Hangul_J_KiyeogSios  Key = 0x0ed6 /* U+11AA HANGUL JONGSEONG KIYEOK-SIOS */
	K_Hangul_J_Nieun       Key = 0x0ed7 /* U+11AB HANGUL JONGSEONG NIEUN */
	K_Hangul_J_NieunJieuj  Key = 0x0ed8 /* U+11AC HANGUL JONGSEONG NIEUN-CIEUC */
	K_Hangul_J_NieunHieuh  Key = 0x0ed9 /* U+11AD HANGUL JONGSEONG NIEUN-HIEUH */
	K_Hangul_J_Dikeud      Key = 0x0eda /* U+11AE HANGUL JONGSEONG TIKEUT */
	K_Hangul_J_Rieul       Key = 0x0edb /* U+11AF HANGUL JONGSEONG RIEUL */
	K_Hangul_J_RieulKiyeog Key = 0x0edc /* U+11B0 HANGUL JONGSEONG RIEUL-KIYEOK */
	K_Hangul_J_RieulMieum  Key = 0x0edd /* U+11B1 HANGUL JONGSEONG RIEUL-MIEUM */
	K_Hangul_J_RieulPieub  Key = 0x0ede /* U+11B2 HANGUL JONGSEONG RIEUL-PIEUP */
	K_Hangul_J_RieulSios   Key = 0x0edf /* U+11B3 HANGUL JONGSEONG RIEUL-SIOS */
	K_Hangul_J_RieulTieut  Key = 0x0ee0 /* U+11B4 HANGUL JONGSEONG RIEUL-THIEUTH */
	K_Hangul_J_RieulPhieuf Key = 0x0ee1 /* U+11B5 HANGUL JONGSEONG RIEUL-PHIEUPH */
	K_Hangul_J_RieulHieuh  Key = 0x0ee2 /* U+11B6 HANGUL JONGSEONG RIEUL-HIEUH */
	K_Hangul_J_Mieum       Key = 0x0ee3 /* U+11B7 HANGUL JONGSEONG MIEUM */
	K_Hangul_J_Pieub       Key = 0x0ee4 /* U+11B8 HANGUL JONGSEONG PIEUP */
	K_Hangul_J_PieubSios   Key = 0x0ee5 /* U+11B9 HANGUL JONGSEONG PIEUP-SIOS */
	K_Hangul_J_Sios        Key = 0x0ee6 /* U+11BA HANGUL JONGSEONG SIOS */
	K_Hangul_J_SsangSios   Key = 0x0ee7 /* U+11BB HANGUL JONGSEONG SSANGSIOS */
	K_Hangul_J_Ieung       Key = 0x0ee8 /* U+11BC HANGUL JONGSEONG IEUNG */
	K_Hangul_J_Jieuj       Key = 0x0ee9 /* U+11BD HANGUL JONGSEONG CIEUC */
	K_Hangul_J_Cieuc       Key = 0x0eea /* U+11BE HANGUL JONGSEONG CHIEUCH */
	K_Hangul_J_Khieuq      Key = 0x0eeb /* U+11BF HANGUL JONGSEONG KHIEUKH */
	K_Hangul_J_Tieut       Key = 0x0eec /* U+11C0 HANGUL JONGSEONG THIEUTH */
	K_Hangul_J_Phieuf      Key = 0x0eed /* U+11C1 HANGUL JONGSEONG PHIEUPH */
	K_Hangul_J_Hieuh       Key = 0x0eee /* U+11C2 HANGUL JONGSEONG HIEUH */

	/* Ancient Hangul Consonant Characters */
	K_Hangul_RieulYeorinHieuh   Key = 0x0eef /* U+316D HANGUL LETTER RIEUL-YEORINHIEUH */
	K_Hangul_SunkyeongeumMieum  Key = 0x0ef0 /* U+3171 HANGUL LETTER KAPYEOUNMIEUM */
	K_Hangul_SunkyeongeumPieub  Key = 0x0ef1 /* U+3178 HANGUL LETTER KAPYEOUNPIEUP */
	K_Hangul_PanSios            Key = 0x0ef2 /* U+317F HANGUL LETTER PANSIOS */
	K_Hangul_KkogjiDalrinIeung  Key = 0x0ef3 /* U+3181 HANGUL LETTER YESIEUNG */
	K_Hangul_SunkyeongeumPhieuf Key = 0x0ef4 /* U+3184 HANGUL LETTER KAPYEOUNPHIEUPH */
	K_Hangul_YeorinHieuh        Key = 0x0ef5 /* U+3186 HANGUL LETTER YEORINHIEUH */

	/* Ancient Hangul Vowel Characters */
	K_Hangul_AraeA  Key = 0x0ef6 /* U+318D HANGUL LETTER ARAEA */
	K_Hangul_AraeAE Key = 0x0ef7 /* U+318E HANGUL LETTER ARAEAE */

	/* Ancient Hangul syllable-final (JongSeong) Characters */
	K_Hangul_J_PanSios           Key = 0x0ef8 /* U+11EB HANGUL JONGSEONG PANSIOS */
	K_Hangul_J_KkogjiDalrinIeung Key = 0x0ef9 /* U+11F0 HANGUL JONGSEONG YESIEUNG */
	K_Hangul_J_YeorinHieuh       Key = 0x0efa /* U+11F9 HANGUL JONGSEONG YEORINHIEUH */

	/* Korean currency symbol */
	K_Korean_Won Key = 0x0eff /*(U+20A9 WON SIGN)*/

	/*
	 * Armenian
	 */

	K_Armenian_ligature_ew     Key = 0x1000587 /* U+0587 ARMENIAN SMALL LIGATURE ECH YIWN */
	K_Armenian_full_stop       Key = 0x1000589 /* U+0589 ARMENIAN FULL STOP */
	K_Armenian_verjaket        Key = 0x1000589 /* deprecated alias for Armenian_full_stop */
	K_Armenian_separation_mark Key = 0x100055d /* U+055D ARMENIAN COMMA */
	K_Armenian_but             Key = 0x100055d /* deprecated alias for Armenian_separation_mark */
	K_Armenian_hyphen          Key = 0x100058a /* U+058A ARMENIAN HYPHEN */
	K_Armenian_yentamna        Key = 0x100058a /* deprecated alias for Armenian_hyphen */
	K_Armenian_exclam          Key = 0x100055c /* U+055C ARMENIAN EXCLAMATION MARK */
	K_Armenian_amanak          Key = 0x100055c /* deprecated alias for Armenian_exclam */
	K_Armenian_accent          Key = 0x100055b /* U+055B ARMENIAN EMPHASIS MARK */
	K_Armenian_shesht          Key = 0x100055b /* deprecated alias for Armenian_accent */
	K_Armenian_question        Key = 0x100055e /* U+055E ARMENIAN QUESTION MARK */
	K_Armenian_paruyk          Key = 0x100055e /* deprecated alias for Armenian_question */
	K_Armenian_AYB             Key = 0x1000531 /* U+0531 ARMENIAN CAPITAL LETTER AYB */
	K_Armenian_ayb             Key = 0x1000561 /* U+0561 ARMENIAN SMALL LETTER AYB */
	K_Armenian_BEN             Key = 0x1000532 /* U+0532 ARMENIAN CAPITAL LETTER BEN */
	K_Armenian_ben             Key = 0x1000562 /* U+0562 ARMENIAN SMALL LETTER BEN */
	K_Armenian_GIM             Key = 0x1000533 /* U+0533 ARMENIAN CAPITAL LETTER GIM */
	K_Armenian_gim             Key = 0x1000563 /* U+0563 ARMENIAN SMALL LETTER GIM */
	K_Armenian_DA              Key = 0x1000534 /* U+0534 ARMENIAN CAPITAL LETTER DA */
	K_Armenian_da              Key = 0x1000564 /* U+0564 ARMENIAN SMALL LETTER DA */
	K_Armenian_YECH            Key = 0x1000535 /* U+0535 ARMENIAN CAPITAL LETTER ECH */
	K_Armenian_yech            Key = 0x1000565 /* U+0565 ARMENIAN SMALL LETTER ECH */
	K_Armenian_ZA              Key = 0x1000536 /* U+0536 ARMENIAN CAPITAL LETTER ZA */
	K_Armenian_za              Key = 0x1000566 /* U+0566 ARMENIAN SMALL LETTER ZA */
	K_Armenian_E               Key = 0x1000537 /* U+0537 ARMENIAN CAPITAL LETTER EH */
	K_Armenian_e               Key = 0x1000567 /* U+0567 ARMENIAN SMALL LETTER EH */
	K_Armenian_AT              Key = 0x1000538 /* U+0538 ARMENIAN CAPITAL LETTER ET */
	K_Armenian_at              Key = 0x1000568 /* U+0568 ARMENIAN SMALL LETTER ET */
	K_Armenian_TO              Key = 0x1000539 /* U+0539 ARMENIAN CAPITAL LETTER TO */
	K_Armenian_to              Key = 0x1000569 /* U+0569 ARMENIAN SMALL LETTER TO */
	K_Armenian_ZHE             Key = 0x100053a /* U+053A ARMENIAN CAPITAL LETTER ZHE */
	K_Armenian_zhe             Key = 0x100056a /* U+056A ARMENIAN SMALL LETTER ZHE */
	K_Armenian_INI             Key = 0x100053b /* U+053B ARMENIAN CAPITAL LETTER INI */
	K_Armenian_ini             Key = 0x100056b /* U+056B ARMENIAN SMALL LETTER INI */
	K_Armenian_LYUN            Key = 0x100053c /* U+053C ARMENIAN CAPITAL LETTER LIWN */
	K_Armenian_lyun            Key = 0x100056c /* U+056C ARMENIAN SMALL LETTER LIWN */
	K_Armenian_KHE             Key = 0x100053d /* U+053D ARMENIAN CAPITAL LETTER XEH */
	K_Armenian_khe             Key = 0x100056d /* U+056D ARMENIAN SMALL LETTER XEH */
	K_Armenian_TSA             Key = 0x100053e /* U+053E ARMENIAN CAPITAL LETTER CA */
	K_Armenian_tsa             Key = 0x100056e /* U+056E ARMENIAN SMALL LETTER CA */
	K_Armenian_KEN             Key = 0x100053f /* U+053F ARMENIAN CAPITAL LETTER KEN */
	K_Armenian_ken             Key = 0x100056f /* U+056F ARMENIAN SMALL LETTER KEN */
	K_Armenian_HO              Key = 0x1000540 /* U+0540 ARMENIAN CAPITAL LETTER HO */
	K_Armenian_ho              Key = 0x1000570 /* U+0570 ARMENIAN SMALL LETTER HO */
	K_Armenian_DZA             Key = 0x1000541 /* U+0541 ARMENIAN CAPITAL LETTER JA */
	K_Armenian_dza             Key = 0x1000571 /* U+0571 ARMENIAN SMALL LETTER JA */
	K_Armenian_GHAT            Key = 0x1000542 /* U+0542 ARMENIAN CAPITAL LETTER GHAD */
	K_Armenian_ghat            Key = 0x1000572 /* U+0572 ARMENIAN SMALL LETTER GHAD */
	K_Armenian_TCHE            Key = 0x1000543 /* U+0543 ARMENIAN CAPITAL LETTER CHEH */
	K_Armenian_tche            Key = 0x1000573 /* U+0573 ARMENIAN SMALL LETTER CHEH */
	K_Armenian_MEN             Key = 0x1000544 /* U+0544 ARMENIAN CAPITAL LETTER MEN */
	K_Armenian_men             Key = 0x1000574 /* U+0574 ARMENIAN SMALL LETTER MEN */
	K_Armenian_HI              Key = 0x1000545 /* U+0545 ARMENIAN CAPITAL LETTER YI */
	K_Armenian_hi              Key = 0x1000575 /* U+0575 ARMENIAN SMALL LETTER YI */
	K_Armenian_NU              Key = 0x1000546 /* U+0546 ARMENIAN CAPITAL LETTER NOW */
	K_Armenian_nu              Key = 0x1000576 /* U+0576 ARMENIAN SMALL LETTER NOW */
	K_Armenian_SHA             Key = 0x1000547 /* U+0547 ARMENIAN CAPITAL LETTER SHA */
	K_Armenian_sha             Key = 0x1000577 /* U+0577 ARMENIAN SMALL LETTER SHA */
	K_Armenian_VO              Key = 0x1000548 /* U+0548 ARMENIAN CAPITAL LETTER VO */
	K_Armenian_vo              Key = 0x1000578 /* U+0578 ARMENIAN SMALL LETTER VO */
	K_Armenian_CHA             Key = 0x1000549 /* U+0549 ARMENIAN CAPITAL LETTER CHA */
	K_Armenian_cha             Key = 0x1000579 /* U+0579 ARMENIAN SMALL LETTER CHA */
	K_Armenian_PE              Key = 0x100054a /* U+054A ARMENIAN CAPITAL LETTER PEH */
	K_Armenian_pe              Key = 0x100057a /* U+057A ARMENIAN SMALL LETTER PEH */
	K_Armenian_JE              Key = 0x100054b /* U+054B ARMENIAN CAPITAL LETTER JHEH */
	K_Armenian_je              Key = 0x100057b /* U+057B ARMENIAN SMALL LETTER JHEH */
	K_Armenian_RA              Key = 0x100054c /* U+054C ARMENIAN CAPITAL LETTER RA */
	K_Armenian_ra              Key = 0x100057c /* U+057C ARMENIAN SMALL LETTER RA */
	K_Armenian_SE              Key = 0x100054d /* U+054D ARMENIAN CAPITAL LETTER SEH */
	K_Armenian_se              Key = 0x100057d /* U+057D ARMENIAN SMALL LETTER SEH */
	K_Armenian_VEV             Key = 0x100054e /* U+054E ARMENIAN CAPITAL LETTER VEW */
	K_Armenian_vev             Key = 0x100057e /* U+057E ARMENIAN SMALL LETTER VEW */
	K_Armenian_TYUN            Key = 0x100054f /* U+054F ARMENIAN CAPITAL LETTER TIWN */
	K_Armenian_tyun            Key = 0x100057f /* U+057F ARMENIAN SMALL LETTER TIWN */
	K_Armenian_RE              Key = 0x1000550 /* U+0550 ARMENIAN CAPITAL LETTER REH */
	K_Armenian_re              Key = 0x1000580 /* U+0580 ARMENIAN SMALL LETTER REH */
	K_Armenian_TSO             Key = 0x1000551 /* U+0551 ARMENIAN CAPITAL LETTER CO */
	K_Armenian_tso             Key = 0x1000581 /* U+0581 ARMENIAN SMALL LETTER CO */
	K_Armenian_VYUN            Key = 0x1000552 /* U+0552 ARMENIAN CAPITAL LETTER YIWN */
	K_Armenian_vyun            Key = 0x1000582 /* U+0582 ARMENIAN SMALL LETTER YIWN */
	K_Armenian_PYUR            Key = 0x1000553 /* U+0553 ARMENIAN CAPITAL LETTER PIWR */
	K_Armenian_pyur            Key = 0x1000583 /* U+0583 ARMENIAN SMALL LETTER PIWR */
	K_Armenian_KE              Key = 0x1000554 /* U+0554 ARMENIAN CAPITAL LETTER KEH */
	K_Armenian_ke              Key = 0x1000584 /* U+0584 ARMENIAN SMALL LETTER KEH */
	K_Armenian_O               Key = 0x1000555 /* U+0555 ARMENIAN CAPITAL LETTER OH */
	K_Armenian_o               Key = 0x1000585 /* U+0585 ARMENIAN SMALL LETTER OH */
	K_Armenian_FE              Key = 0x1000556 /* U+0556 ARMENIAN CAPITAL LETTER FEH */
	K_Armenian_fe              Key = 0x1000586 /* U+0586 ARMENIAN SMALL LETTER FEH */
	K_Armenian_apostrophe      Key = 0x100055a /* U+055A ARMENIAN APOSTROPHE */

	/*
	 * Georgian
	 */

	K_Georgian_an   Key = 0x10010d0 /* U+10D0 GEORGIAN LETTER AN */
	K_Georgian_ban  Key = 0x10010d1 /* U+10D1 GEORGIAN LETTER BAN */
	K_Georgian_gan  Key = 0x10010d2 /* U+10D2 GEORGIAN LETTER GAN */
	K_Georgian_don  Key = 0x10010d3 /* U+10D3 GEORGIAN LETTER DON */
	K_Georgian_en   Key = 0x10010d4 /* U+10D4 GEORGIAN LETTER EN */
	K_Georgian_vin  Key = 0x10010d5 /* U+10D5 GEORGIAN LETTER VIN */
	K_Georgian_zen  Key = 0x10010d6 /* U+10D6 GEORGIAN LETTER ZEN */
	K_Georgian_tan  Key = 0x10010d7 /* U+10D7 GEORGIAN LETTER TAN */
	K_Georgian_in   Key = 0x10010d8 /* U+10D8 GEORGIAN LETTER IN */
	K_Georgian_kan  Key = 0x10010d9 /* U+10D9 GEORGIAN LETTER KAN */
	K_Georgian_las  Key = 0x10010da /* U+10DA GEORGIAN LETTER LAS */
	K_Georgian_man  Key = 0x10010db /* U+10DB GEORGIAN LETTER MAN */
	K_Georgian_nar  Key = 0x10010dc /* U+10DC GEORGIAN LETTER NAR */
	K_Georgian_on   Key = 0x10010dd /* U+10DD GEORGIAN LETTER ON */
	K_Georgian_par  Key = 0x10010de /* U+10DE GEORGIAN LETTER PAR */
	K_Georgian_zhar Key = 0x10010df /* U+10DF GEORGIAN LETTER ZHAR */
	K_Georgian_rae  Key = 0x10010e0 /* U+10E0 GEORGIAN LETTER RAE */
	K_Georgian_san  Key = 0x10010e1 /* U+10E1 GEORGIAN LETTER SAN */
	K_Georgian_tar  Key = 0x10010e2 /* U+10E2 GEORGIAN LETTER TAR */
	K_Georgian_un   Key = 0x10010e3 /* U+10E3 GEORGIAN LETTER UN */
	K_Georgian_phar Key = 0x10010e4 /* U+10E4 GEORGIAN LETTER PHAR */
	K_Georgian_khar Key = 0x10010e5 /* U+10E5 GEORGIAN LETTER KHAR */
	K_Georgian_ghan Key = 0x10010e6 /* U+10E6 GEORGIAN LETTER GHAN */
	K_Georgian_qar  Key = 0x10010e7 /* U+10E7 GEORGIAN LETTER QAR */
	K_Georgian_shin Key = 0x10010e8 /* U+10E8 GEORGIAN LETTER SHIN */
	K_Georgian_chin Key = 0x10010e9 /* U+10E9 GEORGIAN LETTER CHIN */
	K_Georgian_can  Key = 0x10010ea /* U+10EA GEORGIAN LETTER CAN */
	K_Georgian_jil  Key = 0x10010eb /* U+10EB GEORGIAN LETTER JIL */
	K_Georgian_cil  Key = 0x10010ec /* U+10EC GEORGIAN LETTER CIL */
	K_Georgian_char Key = 0x10010ed /* U+10ED GEORGIAN LETTER CHAR */
	K_Georgian_xan  Key = 0x10010ee /* U+10EE GEORGIAN LETTER XAN */
	K_Georgian_jhan Key = 0x10010ef /* U+10EF GEORGIAN LETTER JHAN */
	K_Georgian_hae  Key = 0x10010f0 /* U+10F0 GEORGIAN LETTER HAE */
	K_Georgian_he   Key = 0x10010f1 /* U+10F1 GEORGIAN LETTER HE */
	K_Georgian_hie  Key = 0x10010f2 /* U+10F2 GEORGIAN LETTER HIE */
	K_Georgian_we   Key = 0x10010f3 /* U+10F3 GEORGIAN LETTER WE */
	K_Georgian_har  Key = 0x10010f4 /* U+10F4 GEORGIAN LETTER HAR */
	K_Georgian_hoe  Key = 0x10010f5 /* U+10F5 GEORGIAN LETTER HOE */
	K_Georgian_fi   Key = 0x10010f6 /* U+10F6 GEORGIAN LETTER FI */

	/*
	 * Azeri (and other Turkic or Caucasian languages)
	 */

	/* latin */
	K_Xabovedot Key = 0x1001e8a /* U+1E8A LATIN CAPITAL LETTER X WITH DOT ABOVE */
	K_Ibreve    Key = 0x100012c /* U+012C LATIN CAPITAL LETTER I WITH BREVE */
	K_Zstroke   Key = 0x10001b5 /* U+01B5 LATIN CAPITAL LETTER Z WITH STROKE */
	K_Gcaron    Key = 0x10001e6 /* U+01E6 LATIN CAPITAL LETTER G WITH CARON */
	K_Ocaron    Key = 0x10001d1 /* U+01D1 LATIN CAPITAL LETTER O WITH CARON */
	K_Obarred   Key = 0x100019f /* U+019F LATIN CAPITAL LETTER O WITH MIDDLE TILDE */
	K_xabovedot Key = 0x1001e8b /* U+1E8B LATIN SMALL LETTER X WITH DOT ABOVE */
	K_ibreve    Key = 0x100012d /* U+012D LATIN SMALL LETTER I WITH BREVE */
	K_zstroke   Key = 0x10001b6 /* U+01B6 LATIN SMALL LETTER Z WITH STROKE */
	K_gcaron    Key = 0x10001e7 /* U+01E7 LATIN SMALL LETTER G WITH CARON */
	K_ocaron    Key = 0x10001d2 /* U+01D2 LATIN SMALL LETTER O WITH CARON */
	K_obarred   Key = 0x1000275 /* U+0275 LATIN SMALL LETTER BARRED O */
	K_SCHWA     Key = 0x100018f /* U+018F LATIN CAPITAL LETTER SCHWA */
	K_schwa     Key = 0x1000259 /* U+0259 LATIN SMALL LETTER SCHWA */
	K_EZH       Key = 0x10001b7 /* U+01B7 LATIN CAPITAL LETTER EZH */
	K_ezh       Key = 0x1000292 /* U+0292 LATIN SMALL LETTER EZH */
	/* those are not really Caucasus */
	/* For Inupiak */
	K_Lbelowdot Key = 0x1001e36 /* U+1E36 LATIN CAPITAL LETTER L WITH DOT BELOW */
	K_lbelowdot Key = 0x1001e37 /* U+1E37 LATIN SMALL LETTER L WITH DOT BELOW */

	/*
	 * Vietnamese
	 */

	K_Abelowdot           Key = 0x1001ea0 /* U+1EA0 LATIN CAPITAL LETTER A WITH DOT BELOW */
	K_abelowdot           Key = 0x1001ea1 /* U+1EA1 LATIN SMALL LETTER A WITH DOT BELOW */
	K_Ahook               Key = 0x1001ea2 /* U+1EA2 LATIN CAPITAL LETTER A WITH HOOK ABOVE */
	K_ahook               Key = 0x1001ea3 /* U+1EA3 LATIN SMALL LETTER A WITH HOOK ABOVE */
	K_Acircumflexacute    Key = 0x1001ea4 /* U+1EA4 LATIN CAPITAL LETTER A WITH CIRCUMFLEX AND ACUTE */
	K_acircumflexacute    Key = 0x1001ea5 /* U+1EA5 LATIN SMALL LETTER A WITH CIRCUMFLEX AND ACUTE */
	K_Acircumflexgrave    Key = 0x1001ea6 /* U+1EA6 LATIN CAPITAL LETTER A WITH CIRCUMFLEX AND GRAVE */
	K_acircumflexgrave    Key = 0x1001ea7 /* U+1EA7 LATIN SMALL LETTER A WITH CIRCUMFLEX AND GRAVE */
	K_Acircumflexhook     Key = 0x1001ea8 /* U+1EA8 LATIN CAPITAL LETTER A WITH CIRCUMFLEX AND HOOK ABOVE */
	K_acircumflexhook     Key = 0x1001ea9 /* U+1EA9 LATIN SMALL LETTER A WITH CIRCUMFLEX AND HOOK ABOVE */
	K_Acircumflextilde    Key = 0x1001eaa /* U+1EAA LATIN CAPITAL LETTER A WITH CIRCUMFLEX AND TILDE */
	K_acircumflextilde    Key = 0x1001eab /* U+1EAB LATIN SMALL LETTER A WITH CIRCUMFLEX AND TILDE */
	K_Acircumflexbelowdot Key = 0x1001eac /* U+1EAC LATIN CAPITAL LETTER A WITH CIRCUMFLEX AND DOT BELOW */
	K_acircumflexbelowdot Key = 0x1001ead /* U+1EAD LATIN SMALL LETTER A WITH CIRCUMFLEX AND DOT BELOW */
	K_Abreveacute         Key = 0x1001eae /* U+1EAE LATIN CAPITAL LETTER A WITH BREVE AND ACUTE */
	K_abreveacute         Key = 0x1001eaf /* U+1EAF LATIN SMALL LETTER A WITH BREVE AND ACUTE */
	K_Abrevegrave         Key = 0x1001eb0 /* U+1EB0 LATIN CAPITAL LETTER A WITH BREVE AND GRAVE */
	K_abrevegrave         Key = 0x1001eb1 /* U+1EB1 LATIN SMALL LETTER A WITH BREVE AND GRAVE */
	K_Abrevehook          Key = 0x1001eb2 /* U+1EB2 LATIN CAPITAL LETTER A WITH BREVE AND HOOK ABOVE */
	K_abrevehook          Key = 0x1001eb3 /* U+1EB3 LATIN SMALL LETTER A WITH BREVE AND HOOK ABOVE */
	K_Abrevetilde         Key = 0x1001eb4 /* U+1EB4 LATIN CAPITAL LETTER A WITH BREVE AND TILDE */
	K_abrevetilde         Key = 0x1001eb5 /* U+1EB5 LATIN SMALL LETTER A WITH BREVE AND TILDE */
	K_Abrevebelowdot      Key = 0x1001eb6 /* U+1EB6 LATIN CAPITAL LETTER A WITH BREVE AND DOT BELOW */
	K_abrevebelowdot      Key = 0x1001eb7 /* U+1EB7 LATIN SMALL LETTER A WITH BREVE AND DOT BELOW */
	K_Ebelowdot           Key = 0x1001eb8 /* U+1EB8 LATIN CAPITAL LETTER E WITH DOT BELOW */
	K_ebelowdot           Key = 0x1001eb9 /* U+1EB9 LATIN SMALL LETTER E WITH DOT BELOW */
	K_Ehook               Key = 0x1001eba /* U+1EBA LATIN CAPITAL LETTER E WITH HOOK ABOVE */
	K_ehook               Key = 0x1001ebb /* U+1EBB LATIN SMALL LETTER E WITH HOOK ABOVE */
	K_Etilde              Key = 0x1001ebc /* U+1EBC LATIN CAPITAL LETTER E WITH TILDE */
	K_etilde              Key = 0x1001ebd /* U+1EBD LATIN SMALL LETTER E WITH TILDE */
	K_Ecircumflexacute    Key = 0x1001ebe /* U+1EBE LATIN CAPITAL LETTER E WITH CIRCUMFLEX AND ACUTE */
	K_ecircumflexacute    Key = 0x1001ebf /* U+1EBF LATIN SMALL LETTER E WITH CIRCUMFLEX AND ACUTE */
	K_Ecircumflexgrave    Key = 0x1001ec0 /* U+1EC0 LATIN CAPITAL LETTER E WITH CIRCUMFLEX AND GRAVE */
	K_ecircumflexgrave    Key = 0x1001ec1 /* U+1EC1 LATIN SMALL LETTER E WITH CIRCUMFLEX AND GRAVE */
	K_Ecircumflexhook     Key = 0x1001ec2 /* U+1EC2 LATIN CAPITAL LETTER E WITH CIRCUMFLEX AND HOOK ABOVE */
	K_ecircumflexhook     Key = 0x1001ec3 /* U+1EC3 LATIN SMALL LETTER E WITH CIRCUMFLEX AND HOOK ABOVE */
	K_Ecircumflextilde    Key = 0x1001ec4 /* U+1EC4 LATIN CAPITAL LETTER E WITH CIRCUMFLEX AND TILDE */
	K_ecircumflextilde    Key = 0x1001ec5 /* U+1EC5 LATIN SMALL LETTER E WITH CIRCUMFLEX AND TILDE */
	K_Ecircumflexbelowdot Key = 0x1001ec6 /* U+1EC6 LATIN CAPITAL LETTER E WITH CIRCUMFLEX AND DOT BELOW */
	K_ecircumflexbelowdot Key = 0x1001ec7 /* U+1EC7 LATIN SMALL LETTER E WITH CIRCUMFLEX AND DOT BELOW */
	K_Ihook               Key = 0x1001ec8 /* U+1EC8 LATIN CAPITAL LETTER I WITH HOOK ABOVE */
	K_ihook               Key = 0x1001ec9 /* U+1EC9 LATIN SMALL LETTER I WITH HOOK ABOVE */
	K_Ibelowdot           Key = 0x1001eca /* U+1ECA LATIN CAPITAL LETTER I WITH DOT BELOW */
	K_ibelowdot           Key = 0x1001ecb /* U+1ECB LATIN SMALL LETTER I WITH DOT BELOW */
	K_Obelowdot           Key = 0x1001ecc /* U+1ECC LATIN CAPITAL LETTER O WITH DOT BELOW */
	K_obelowdot           Key = 0x1001ecd /* U+1ECD LATIN SMALL LETTER O WITH DOT BELOW */
	K_Ohook               Key = 0x1001ece /* U+1ECE LATIN CAPITAL LETTER O WITH HOOK ABOVE */
	K_ohook               Key = 0x1001ecf /* U+1ECF LATIN SMALL LETTER O WITH HOOK ABOVE */
	K_Ocircumflexacute    Key = 0x1001ed0 /* U+1ED0 LATIN CAPITAL LETTER O WITH CIRCUMFLEX AND ACUTE */
	K_ocircumflexacute    Key = 0x1001ed1 /* U+1ED1 LATIN SMALL LETTER O WITH CIRCUMFLEX AND ACUTE */
	K_Ocircumflexgrave    Key = 0x1001ed2 /* U+1ED2 LATIN CAPITAL LETTER O WITH CIRCUMFLEX AND GRAVE */
	K_ocircumflexgrave    Key = 0x1001ed3 /* U+1ED3 LATIN SMALL LETTER O WITH CIRCUMFLEX AND GRAVE */
	K_Ocircumflexhook     Key = 0x1001ed4 /* U+1ED4 LATIN CAPITAL LETTER O WITH CIRCUMFLEX AND HOOK ABOVE */
	K_ocircumflexhook     Key = 0x1001ed5 /* U+1ED5 LATIN SMALL LETTER O WITH CIRCUMFLEX AND HOOK ABOVE */
	K_Ocircumflextilde    Key = 0x1001ed6 /* U+1ED6 LATIN CAPITAL LETTER O WITH CIRCUMFLEX AND TILDE */
	K_ocircumflextilde    Key = 0x1001ed7 /* U+1ED7 LATIN SMALL LETTER O WITH CIRCUMFLEX AND TILDE */
	K_Ocircumflexbelowdot Key = 0x1001ed8 /* U+1ED8 LATIN CAPITAL LETTER O WITH CIRCUMFLEX AND DOT BELOW */
	K_ocircumflexbelowdot Key = 0x1001ed9 /* U+1ED9 LATIN SMALL LETTER O WITH CIRCUMFLEX AND DOT BELOW */
	K_Ohornacute          Key = 0x1001eda /* U+1EDA LATIN CAPITAL LETTER O WITH HORN AND ACUTE */
	K_ohornacute          Key = 0x1001edb /* U+1EDB LATIN SMALL LETTER O WITH HORN AND ACUTE */
	K_Ohorngrave          Key = 0x1001edc /* U+1EDC LATIN CAPITAL LETTER O WITH HORN AND GRAVE */
	K_ohorngrave          Key = 0x1001edd /* U+1EDD LATIN SMALL LETTER O WITH HORN AND GRAVE */
	K_Ohornhook           Key = 0x1001ede /* U+1EDE LATIN CAPITAL LETTER O WITH HORN AND HOOK ABOVE */
	K_ohornhook           Key = 0x1001edf /* U+1EDF LATIN SMALL LETTER O WITH HORN AND HOOK ABOVE */
	K_Ohorntilde          Key = 0x1001ee0 /* U+1EE0 LATIN CAPITAL LETTER O WITH HORN AND TILDE */
	K_ohorntilde          Key = 0x1001ee1 /* U+1EE1 LATIN SMALL LETTER O WITH HORN AND TILDE */
	K_Ohornbelowdot       Key = 0x1001ee2 /* U+1EE2 LATIN CAPITAL LETTER O WITH HORN AND DOT BELOW */
	K_ohornbelowdot       Key = 0x1001ee3 /* U+1EE3 LATIN SMALL LETTER O WITH HORN AND DOT BELOW */
	K_Ubelowdot           Key = 0x1001ee4 /* U+1EE4 LATIN CAPITAL LETTER U WITH DOT BELOW */
	K_ubelowdot           Key = 0x1001ee5 /* U+1EE5 LATIN SMALL LETTER U WITH DOT BELOW */
	K_Uhook               Key = 0x1001ee6 /* U+1EE6 LATIN CAPITAL LETTER U WITH HOOK ABOVE */
	K_uhook               Key = 0x1001ee7 /* U+1EE7 LATIN SMALL LETTER U WITH HOOK ABOVE */
	K_Uhornacute          Key = 0x1001ee8 /* U+1EE8 LATIN CAPITAL LETTER U WITH HORN AND ACUTE */
	K_uhornacute          Key = 0x1001ee9 /* U+1EE9 LATIN SMALL LETTER U WITH HORN AND ACUTE */
	K_Uhorngrave          Key = 0x1001eea /* U+1EEA LATIN CAPITAL LETTER U WITH HORN AND GRAVE */
	K_uhorngrave          Key = 0x1001eeb /* U+1EEB LATIN SMALL LETTER U WITH HORN AND GRAVE */
	K_Uhornhook           Key = 0x1001eec /* U+1EEC LATIN CAPITAL LETTER U WITH HORN AND HOOK ABOVE */
	K_uhornhook           Key = 0x1001eed /* U+1EED LATIN SMALL LETTER U WITH HORN AND HOOK ABOVE */
	K_Uhorntilde          Key = 0x1001eee /* U+1EEE LATIN CAPITAL LETTER U WITH HORN AND TILDE */
	K_uhorntilde          Key = 0x1001eef /* U+1EEF LATIN SMALL LETTER U WITH HORN AND TILDE */
	K_Uhornbelowdot       Key = 0x1001ef0 /* U+1EF0 LATIN CAPITAL LETTER U WITH HORN AND DOT BELOW */
	K_uhornbelowdot       Key = 0x1001ef1 /* U+1EF1 LATIN SMALL LETTER U WITH HORN AND DOT BELOW */
	K_Ybelowdot           Key = 0x1001ef4 /* U+1EF4 LATIN CAPITAL LETTER Y WITH DOT BELOW */
	K_ybelowdot           Key = 0x1001ef5 /* U+1EF5 LATIN SMALL LETTER Y WITH DOT BELOW */
	K_Yhook               Key = 0x1001ef6 /* U+1EF6 LATIN CAPITAL LETTER Y WITH HOOK ABOVE */
	K_yhook               Key = 0x1001ef7 /* U+1EF7 LATIN SMALL LETTER Y WITH HOOK ABOVE */
	K_Ytilde              Key = 0x1001ef8 /* U+1EF8 LATIN CAPITAL LETTER Y WITH TILDE */
	K_ytilde              Key = 0x1001ef9 /* U+1EF9 LATIN SMALL LETTER Y WITH TILDE */
	K_Ohorn               Key = 0x10001a0 /* U+01A0 LATIN CAPITAL LETTER O WITH HORN */
	K_ohorn               Key = 0x10001a1 /* U+01A1 LATIN SMALL LETTER O WITH HORN */
	K_Uhorn               Key = 0x10001af /* U+01AF LATIN CAPITAL LETTER U WITH HORN */
	K_uhorn               Key = 0x10001b0 /* U+01B0 LATIN SMALL LETTER U WITH HORN */
	K_combining_tilde     Key = 0x1000303 /* U+0303 COMBINING TILDE */
	K_combining_grave     Key = 0x1000300 /* U+0300 COMBINING GRAVE ACCENT */
	K_combining_acute     Key = 0x1000301 /* U+0301 COMBINING ACUTE ACCENT */
	K_combining_hook      Key = 0x1000309 /* U+0309 COMBINING HOOK ABOVE */
	K_combining_belowdot  Key = 0x1000323 /* U+0323 COMBINING DOT BELOW */

	K_EcuSign       Key = 0x10020a0 /* U+20A0 EURO-CURRENCY SIGN */
	K_ColonSign     Key = 0x10020a1 /* U+20A1 COLON SIGN */
	K_CruzeiroSign  Key = 0x10020a2 /* U+20A2 CRUZEIRO SIGN */
	K_FFrancSign    Key = 0x10020a3 /* U+20A3 FRENCH FRANC SIGN */
	K_LiraSign      Key = 0x10020a4 /* U+20A4 LIRA SIGN */
	K_MillSign      Key = 0x10020a5 /* U+20A5 MILL SIGN */
	K_NairaSign     Key = 0x10020a6 /* U+20A6 NAIRA SIGN */
	K_PesetaSign    Key = 0x10020a7 /* U+20A7 PESETA SIGN */
	K_RupeeSign     Key = 0x10020a8 /* U+20A8 RUPEE SIGN */
	K_WonSign       Key = 0x10020a9 /* U+20A9 WON SIGN */
	K_NewSheqelSign Key = 0x10020aa /* U+20AA NEW SHEQEL SIGN */
	K_DongSign      Key = 0x10020ab /* U+20AB DONG SIGN */
	K_EuroSign      Key = 0x20ac    /* U+20AC EURO SIGN */

	/* one, two and three are defined above. */
	K_zerosuperior     Key = 0x1002070 /* U+2070 SUPERSCRIPT ZERO */
	K_foursuperior     Key = 0x1002074 /* U+2074 SUPERSCRIPT FOUR */
	K_fivesuperior     Key = 0x1002075 /* U+2075 SUPERSCRIPT FIVE */
	K_sixsuperior      Key = 0x1002076 /* U+2076 SUPERSCRIPT SIX */
	K_sevensuperior    Key = 0x1002077 /* U+2077 SUPERSCRIPT SEVEN */
	K_eightsuperior    Key = 0x1002078 /* U+2078 SUPERSCRIPT EIGHT */
	K_ninesuperior     Key = 0x1002079 /* U+2079 SUPERSCRIPT NINE */
	K_zerosubscript    Key = 0x1002080 /* U+2080 SUBSCRIPT ZERO */
	K_onesubscript     Key = 0x1002081 /* U+2081 SUBSCRIPT ONE */
	K_twosubscript     Key = 0x1002082 /* U+2082 SUBSCRIPT TWO */
	K_threesubscript   Key = 0x1002083 /* U+2083 SUBSCRIPT THREE */
	K_foursubscript    Key = 0x1002084 /* U+2084 SUBSCRIPT FOUR */
	K_fivesubscript    Key = 0x1002085 /* U+2085 SUBSCRIPT FIVE */
	K_sixsubscript     Key = 0x1002086 /* U+2086 SUBSCRIPT SIX */
	K_sevensubscript   Key = 0x1002087 /* U+2087 SUBSCRIPT SEVEN */
	K_eightsubscript   Key = 0x1002088 /* U+2088 SUBSCRIPT EIGHT */
	K_ninesubscript    Key = 0x1002089 /* U+2089 SUBSCRIPT NINE */
	K_partdifferential Key = 0x1002202 /* U+2202 PARTIAL DIFFERENTIAL */
	K_emptyset         Key = 0x1002205 /* U+2205 EMPTY SET */
	K_elementof        Key = 0x1002208 /* U+2208 ELEMENT OF */
	K_notelementof     Key = 0x1002209 /* U+2209 NOT AN ELEMENT OF */
	K_containsas       Key = 0x100220b /* U+220B CONTAINS AS MEMBER */
	K_squareroot       Key = 0x100221a /* U+221A SQUARE ROOT */
	K_cuberoot         Key = 0x100221b /* U+221B CUBE ROOT */
	K_fourthroot       Key = 0x100221c /* U+221C FOURTH ROOT */
	K_dintegral        Key = 0x100222c /* U+222C DOUBLE INTEGRAL */
	K_tintegral        Key = 0x100222d /* U+222D TRIPLE INTEGRAL */
	K_because          Key = 0x1002235 /* U+2235 BECAUSE */
	K_approxeq         Key = 0x1002248 /*(U+2248 ALMOST EQUAL TO)*/
	K_notapproxeq      Key = 0x1002247 /*(U+2247 NEITHER APPROXIMATELY NOR ACTUALLY EQUAL TO)*/
	K_notidentical     Key = 0x1002262 /* U+2262 NOT IDENTICAL TO */
	K_stricteq         Key = 0x1002263 /* U+2263 STRICTLY EQUIVALENT TO */

	K_braille_dot_1         Key = 0xfff1
	K_braille_dot_2         Key = 0xfff2
	K_braille_dot_3         Key = 0xfff3
	K_braille_dot_4         Key = 0xfff4
	K_braille_dot_5         Key = 0xfff5
	K_braille_dot_6         Key = 0xfff6
	K_braille_dot_7         Key = 0xfff7
	K_braille_dot_8         Key = 0xfff8
	K_braille_dot_9         Key = 0xfff9
	K_braille_dot_10        Key = 0xfffa
	K_braille_blank         Key = 0x1002800 /* U+2800 BRAILLE PATTERN BLANK */
	K_braille_dots_1        Key = 0x1002801 /* U+2801 BRAILLE PATTERN DOTS-1 */
	K_braille_dots_2        Key = 0x1002802 /* U+2802 BRAILLE PATTERN DOTS-2 */
	K_braille_dots_12       Key = 0x1002803 /* U+2803 BRAILLE PATTERN DOTS-12 */
	K_braille_dots_3        Key = 0x1002804 /* U+2804 BRAILLE PATTERN DOTS-3 */
	K_braille_dots_13       Key = 0x1002805 /* U+2805 BRAILLE PATTERN DOTS-13 */
	K_braille_dots_23       Key = 0x1002806 /* U+2806 BRAILLE PATTERN DOTS-23 */
	K_braille_dots_123      Key = 0x1002807 /* U+2807 BRAILLE PATTERN DOTS-123 */
	K_braille_dots_4        Key = 0x1002808 /* U+2808 BRAILLE PATTERN DOTS-4 */
	K_braille_dots_14       Key = 0x1002809 /* U+2809 BRAILLE PATTERN DOTS-14 */
	K_braille_dots_24       Key = 0x100280a /* U+280A BRAILLE PATTERN DOTS-24 */
	K_braille_dots_124      Key = 0x100280b /* U+280B BRAILLE PATTERN DOTS-124 */
	K_braille_dots_34       Key = 0x100280c /* U+280C BRAILLE PATTERN DOTS-34 */
	K_braille_dots_134      Key = 0x100280d /* U+280D BRAILLE PATTERN DOTS-134 */
	K_braille_dots_234      Key = 0x100280e /* U+280E BRAILLE PATTERN DOTS-234 */
	K_braille_dots_1234     Key = 0x100280f /* U+280F BRAILLE PATTERN DOTS-1234 */
	K_braille_dots_5        Key = 0x1002810 /* U+2810 BRAILLE PATTERN DOTS-5 */
	K_braille_dots_15       Key = 0x1002811 /* U+2811 BRAILLE PATTERN DOTS-15 */
	K_braille_dots_25       Key = 0x1002812 /* U+2812 BRAILLE PATTERN DOTS-25 */
	K_braille_dots_125      Key = 0x1002813 /* U+2813 BRAILLE PATTERN DOTS-125 */
	K_braille_dots_35       Key = 0x1002814 /* U+2814 BRAILLE PATTERN DOTS-35 */
	K_braille_dots_135      Key = 0x1002815 /* U+2815 BRAILLE PATTERN DOTS-135 */
	K_braille_dots_235      Key = 0x1002816 /* U+2816 BRAILLE PATTERN DOTS-235 */
	K_braille_dots_1235     Key = 0x1002817 /* U+2817 BRAILLE PATTERN DOTS-1235 */
	K_braille_dots_45       Key = 0x1002818 /* U+2818 BRAILLE PATTERN DOTS-45 */
	K_braille_dots_145      Key = 0x1002819 /* U+2819 BRAILLE PATTERN DOTS-145 */
	K_braille_dots_245      Key = 0x100281a /* U+281A BRAILLE PATTERN DOTS-245 */
	K_braille_dots_1245     Key = 0x100281b /* U+281B BRAILLE PATTERN DOTS-1245 */
	K_braille_dots_345      Key = 0x100281c /* U+281C BRAILLE PATTERN DOTS-345 */
	K_braille_dots_1345     Key = 0x100281d /* U+281D BRAILLE PATTERN DOTS-1345 */
	K_braille_dots_2345     Key = 0x100281e /* U+281E BRAILLE PATTERN DOTS-2345 */
	K_braille_dots_12345    Key = 0x100281f /* U+281F BRAILLE PATTERN DOTS-12345 */
	K_braille_dots_6        Key = 0x1002820 /* U+2820 BRAILLE PATTERN DOTS-6 */
	K_braille_dots_16       Key = 0x1002821 /* U+2821 BRAILLE PATTERN DOTS-16 */
	K_braille_dots_26       Key = 0x1002822 /* U+2822 BRAILLE PATTERN DOTS-26 */
	K_braille_dots_126      Key = 0x1002823 /* U+2823 BRAILLE PATTERN DOTS-126 */
	K_braille_dots_36       Key = 0x1002824 /* U+2824 BRAILLE PATTERN DOTS-36 */
	K_braille_dots_136      Key = 0x1002825 /* U+2825 BRAILLE PATTERN DOTS-136 */
	K_braille_dots_236      Key = 0x1002826 /* U+2826 BRAILLE PATTERN DOTS-236 */
	K_braille_dots_1236     Key = 0x1002827 /* U+2827 BRAILLE PATTERN DOTS-1236 */
	K_braille_dots_46       Key = 0x1002828 /* U+2828 BRAILLE PATTERN DOTS-46 */
	K_braille_dots_146      Key = 0x1002829 /* U+2829 BRAILLE PATTERN DOTS-146 */
	K_braille_dots_246      Key = 0x100282a /* U+282A BRAILLE PATTERN DOTS-246 */
	K_braille_dots_1246     Key = 0x100282b /* U+282B BRAILLE PATTERN DOTS-1246 */
	K_braille_dots_346      Key = 0x100282c /* U+282C BRAILLE PATTERN DOTS-346 */
	K_braille_dots_1346     Key = 0x100282d /* U+282D BRAILLE PATTERN DOTS-1346 */
	K_braille_dots_2346     Key = 0x100282e /* U+282E BRAILLE PATTERN DOTS-2346 */
	K_braille_dots_12346    Key = 0x100282f /* U+282F BRAILLE PATTERN DOTS-12346 */
	K_braille_dots_56       Key = 0x1002830 /* U+2830 BRAILLE PATTERN DOTS-56 */
	K_braille_dots_156      Key = 0x1002831 /* U+2831 BRAILLE PATTERN DOTS-156 */
	K_braille_dots_256      Key = 0x1002832 /* U+2832 BRAILLE PATTERN DOTS-256 */
	K_braille_dots_1256     Key = 0x1002833 /* U+2833 BRAILLE PATTERN DOTS-1256 */
	K_braille_dots_356      Key = 0x1002834 /* U+2834 BRAILLE PATTERN DOTS-356 */
	K_braille_dots_1356     Key = 0x1002835 /* U+2835 BRAILLE PATTERN DOTS-1356 */
	K_braille_dots_2356     Key = 0x1002836 /* U+2836 BRAILLE PATTERN DOTS-2356 */
	K_braille_dots_12356    Key = 0x1002837 /* U+2837 BRAILLE PATTERN DOTS-12356 */
	K_braille_dots_456      Key = 0x1002838 /* U+2838 BRAILLE PATTERN DOTS-456 */
	K_braille_dots_1456     Key = 0x1002839 /* U+2839 BRAILLE PATTERN DOTS-1456 */
	K_braille_dots_2456     Key = 0x100283a /* U+283A BRAILLE PATTERN DOTS-2456 */
	K_braille_dots_12456    Key = 0x100283b /* U+283B BRAILLE PATTERN DOTS-12456 */
	K_braille_dots_3456     Key = 0x100283c /* U+283C BRAILLE PATTERN DOTS-3456 */
	K_braille_dots_13456    Key = 0x100283d /* U+283D BRAILLE PATTERN DOTS-13456 */
	K_braille_dots_23456    Key = 0x100283e /* U+283E BRAILLE PATTERN DOTS-23456 */
	K_braille_dots_123456   Key = 0x100283f /* U+283F BRAILLE PATTERN DOTS-123456 */
	K_braille_dots_7        Key = 0x1002840 /* U+2840 BRAILLE PATTERN DOTS-7 */
	K_braille_dots_17       Key = 0x1002841 /* U+2841 BRAILLE PATTERN DOTS-17 */
	K_braille_dots_27       Key = 0x1002842 /* U+2842 BRAILLE PATTERN DOTS-27 */
	K_braille_dots_127      Key = 0x1002843 /* U+2843 BRAILLE PATTERN DOTS-127 */
	K_braille_dots_37       Key = 0x1002844 /* U+2844 BRAILLE PATTERN DOTS-37 */
	K_braille_dots_137      Key = 0x1002845 /* U+2845 BRAILLE PATTERN DOTS-137 */
	K_braille_dots_237      Key = 0x1002846 /* U+2846 BRAILLE PATTERN DOTS-237 */
	K_braille_dots_1237     Key = 0x1002847 /* U+2847 BRAILLE PATTERN DOTS-1237 */
	K_braille_dots_47       Key = 0x1002848 /* U+2848 BRAILLE PATTERN DOTS-47 */
	K_braille_dots_147      Key = 0x1002849 /* U+2849 BRAILLE PATTERN DOTS-147 */
	K_braille_dots_247      Key = 0x100284a /* U+284A BRAILLE PATTERN DOTS-247 */
	K_braille_dots_1247     Key = 0x100284b /* U+284B BRAILLE PATTERN DOTS-1247 */
	K_braille_dots_347      Key = 0x100284c /* U+284C BRAILLE PATTERN DOTS-347 */
	K_braille_dots_1347     Key = 0x100284d /* U+284D BRAILLE PATTERN DOTS-1347 */
	K_braille_dots_2347     Key = 0x100284e /* U+284E BRAILLE PATTERN DOTS-2347 */
	K_braille_dots_12347    Key = 0x100284f /* U+284F BRAILLE PATTERN DOTS-12347 */
	K_braille_dots_57       Key = 0x1002850 /* U+2850 BRAILLE PATTERN DOTS-57 */
	K_braille_dots_157      Key = 0x1002851 /* U+2851 BRAILLE PATTERN DOTS-157 */
	K_braille_dots_257      Key = 0x1002852 /* U+2852 BRAILLE PATTERN DOTS-257 */
	K_braille_dots_1257     Key = 0x1002853 /* U+2853 BRAILLE PATTERN DOTS-1257 */
	K_braille_dots_357      Key = 0x1002854 /* U+2854 BRAILLE PATTERN DOTS-357 */
	K_braille_dots_1357     Key = 0x1002855 /* U+2855 BRAILLE PATTERN DOTS-1357 */
	K_braille_dots_2357     Key = 0x1002856 /* U+2856 BRAILLE PATTERN DOTS-2357 */
	K_braille_dots_12357    Key = 0x1002857 /* U+2857 BRAILLE PATTERN DOTS-12357 */
	K_braille_dots_457      Key = 0x1002858 /* U+2858 BRAILLE PATTERN DOTS-457 */
	K_braille_dots_1457     Key = 0x1002859 /* U+2859 BRAILLE PATTERN DOTS-1457 */
	K_braille_dots_2457     Key = 0x100285a /* U+285A BRAILLE PATTERN DOTS-2457 */
	K_braille_dots_12457    Key = 0x100285b /* U+285B BRAILLE PATTERN DOTS-12457 */
	K_braille_dots_3457     Key = 0x100285c /* U+285C BRAILLE PATTERN DOTS-3457 */
	K_braille_dots_13457    Key = 0x100285d /* U+285D BRAILLE PATTERN DOTS-13457 */
	K_braille_dots_23457    Key = 0x100285e /* U+285E BRAILLE PATTERN DOTS-23457 */
	K_braille_dots_123457   Key = 0x100285f /* U+285F BRAILLE PATTERN DOTS-123457 */
	K_braille_dots_67       Key = 0x1002860 /* U+2860 BRAILLE PATTERN DOTS-67 */
	K_braille_dots_167      Key = 0x1002861 /* U+2861 BRAILLE PATTERN DOTS-167 */
	K_braille_dots_267      Key = 0x1002862 /* U+2862 BRAILLE PATTERN DOTS-267 */
	K_braille_dots_1267     Key = 0x1002863 /* U+2863 BRAILLE PATTERN DOTS-1267 */
	K_braille_dots_367      Key = 0x1002864 /* U+2864 BRAILLE PATTERN DOTS-367 */
	K_braille_dots_1367     Key = 0x1002865 /* U+2865 BRAILLE PATTERN DOTS-1367 */
	K_braille_dots_2367     Key = 0x1002866 /* U+2866 BRAILLE PATTERN DOTS-2367 */
	K_braille_dots_12367    Key = 0x1002867 /* U+2867 BRAILLE PATTERN DOTS-12367 */
	K_braille_dots_467      Key = 0x1002868 /* U+2868 BRAILLE PATTERN DOTS-467 */
	K_braille_dots_1467     Key = 0x1002869 /* U+2869 BRAILLE PATTERN DOTS-1467 */
	K_braille_dots_2467     Key = 0x100286a /* U+286A BRAILLE PATTERN DOTS-2467 */
	K_braille_dots_12467    Key = 0x100286b /* U+286B BRAILLE PATTERN DOTS-12467 */
	K_braille_dots_3467     Key = 0x100286c /* U+286C BRAILLE PATTERN DOTS-3467 */
	K_braille_dots_13467    Key = 0x100286d /* U+286D BRAILLE PATTERN DOTS-13467 */
	K_braille_dots_23467    Key = 0x100286e /* U+286E BRAILLE PATTERN DOTS-23467 */
	K_braille_dots_123467   Key = 0x100286f /* U+286F BRAILLE PATTERN DOTS-123467 */
	K_braille_dots_567      Key = 0x1002870 /* U+2870 BRAILLE PATTERN DOTS-567 */
	K_braille_dots_1567     Key = 0x1002871 /* U+2871 BRAILLE PATTERN DOTS-1567 */
	K_braille_dots_2567     Key = 0x1002872 /* U+2872 BRAILLE PATTERN DOTS-2567 */
	K_braille_dots_12567    Key = 0x1002873 /* U+2873 BRAILLE PATTERN DOTS-12567 */
	K_braille_dots_3567     Key = 0x1002874 /* U+2874 BRAILLE PATTERN DOTS-3567 */
	K_braille_dots_13567    Key = 0x1002875 /* U+2875 BRAILLE PATTERN DOTS-13567 */
	K_braille_dots_23567    Key = 0x1002876 /* U+2876 BRAILLE PATTERN DOTS-23567 */
	K_braille_dots_123567   Key = 0x1002877 /* U+2877 BRAILLE PATTERN DOTS-123567 */
	K_braille_dots_4567     Key = 0x1002878 /* U+2878 BRAILLE PATTERN DOTS-4567 */
	K_braille_dots_14567    Key = 0x1002879 /* U+2879 BRAILLE PATTERN DOTS-14567 */
	K_braille_dots_24567    Key = 0x100287a /* U+287A BRAILLE PATTERN DOTS-24567 */
	K_braille_dots_124567   Key = 0x100287b /* U+287B BRAILLE PATTERN DOTS-124567 */
	K_braille_dots_34567    Key = 0x100287c /* U+287C BRAILLE PATTERN DOTS-34567 */
	K_braille_dots_134567   Key = 0x100287d /* U+287D BRAILLE PATTERN DOTS-134567 */
	K_braille_dots_234567   Key = 0x100287e /* U+287E BRAILLE PATTERN DOTS-234567 */
	K_braille_dots_1234567  Key = 0x100287f /* U+287F BRAILLE PATTERN DOTS-1234567 */
	K_braille_dots_8        Key = 0x1002880 /* U+2880 BRAILLE PATTERN DOTS-8 */
	K_braille_dots_18       Key = 0x1002881 /* U+2881 BRAILLE PATTERN DOTS-18 */
	K_braille_dots_28       Key = 0x1002882 /* U+2882 BRAILLE PATTERN DOTS-28 */
	K_braille_dots_128      Key = 0x1002883 /* U+2883 BRAILLE PATTERN DOTS-128 */
	K_braille_dots_38       Key = 0x1002884 /* U+2884 BRAILLE PATTERN DOTS-38 */
	K_braille_dots_138      Key = 0x1002885 /* U+2885 BRAILLE PATTERN DOTS-138 */
	K_braille_dots_238      Key = 0x1002886 /* U+2886 BRAILLE PATTERN DOTS-238 */
	K_braille_dots_1238     Key = 0x1002887 /* U+2887 BRAILLE PATTERN DOTS-1238 */
	K_braille_dots_48       Key = 0x1002888 /* U+2888 BRAILLE PATTERN DOTS-48 */
	K_braille_dots_148      Key = 0x1002889 /* U+2889 BRAILLE PATTERN DOTS-148 */
	K_braille_dots_248      Key = 0x100288a /* U+288A BRAILLE PATTERN DOTS-248 */
	K_braille_dots_1248     Key = 0x100288b /* U+288B BRAILLE PATTERN DOTS-1248 */
	K_braille_dots_348      Key = 0x100288c /* U+288C BRAILLE PATTERN DOTS-348 */
	K_braille_dots_1348     Key = 0x100288d /* U+288D BRAILLE PATTERN DOTS-1348 */
	K_braille_dots_2348     Key = 0x100288e /* U+288E BRAILLE PATTERN DOTS-2348 */
	K_braille_dots_12348    Key = 0x100288f /* U+288F BRAILLE PATTERN DOTS-12348 */
	K_braille_dots_58       Key = 0x1002890 /* U+2890 BRAILLE PATTERN DOTS-58 */
	K_braille_dots_158      Key = 0x1002891 /* U+2891 BRAILLE PATTERN DOTS-158 */
	K_braille_dots_258      Key = 0x1002892 /* U+2892 BRAILLE PATTERN DOTS-258 */
	K_braille_dots_1258     Key = 0x1002893 /* U+2893 BRAILLE PATTERN DOTS-1258 */
	K_braille_dots_358      Key = 0x1002894 /* U+2894 BRAILLE PATTERN DOTS-358 */
	K_braille_dots_1358     Key = 0x1002895 /* U+2895 BRAILLE PATTERN DOTS-1358 */
	K_braille_dots_2358     Key = 0x1002896 /* U+2896 BRAILLE PATTERN DOTS-2358 */
	K_braille_dots_12358    Key = 0x1002897 /* U+2897 BRAILLE PATTERN DOTS-12358 */
	K_braille_dots_458      Key = 0x1002898 /* U+2898 BRAILLE PATTERN DOTS-458 */
	K_braille_dots_1458     Key = 0x1002899 /* U+2899 BRAILLE PATTERN DOTS-1458 */
	K_braille_dots_2458     Key = 0x100289a /* U+289A BRAILLE PATTERN DOTS-2458 */
	K_braille_dots_12458    Key = 0x100289b /* U+289B BRAILLE PATTERN DOTS-12458 */
	K_braille_dots_3458     Key = 0x100289c /* U+289C BRAILLE PATTERN DOTS-3458 */
	K_braille_dots_13458    Key = 0x100289d /* U+289D BRAILLE PATTERN DOTS-13458 */
	K_braille_dots_23458    Key = 0x100289e /* U+289E BRAILLE PATTERN DOTS-23458 */
	K_braille_dots_123458   Key = 0x100289f /* U+289F BRAILLE PATTERN DOTS-123458 */
	K_braille_dots_68       Key = 0x10028a0 /* U+28A0 BRAILLE PATTERN DOTS-68 */
	K_braille_dots_168      Key = 0x10028a1 /* U+28A1 BRAILLE PATTERN DOTS-168 */
	K_braille_dots_268      Key = 0x10028a2 /* U+28A2 BRAILLE PATTERN DOTS-268 */
	K_braille_dots_1268     Key = 0x10028a3 /* U+28A3 BRAILLE PATTERN DOTS-1268 */
	K_braille_dots_368      Key = 0x10028a4 /* U+28A4 BRAILLE PATTERN DOTS-368 */
	K_braille_dots_1368     Key = 0x10028a5 /* U+28A5 BRAILLE PATTERN DOTS-1368 */
	K_braille_dots_2368     Key = 0x10028a6 /* U+28A6 BRAILLE PATTERN DOTS-2368 */
	K_braille_dots_12368    Key = 0x10028a7 /* U+28A7 BRAILLE PATTERN DOTS-12368 */
	K_braille_dots_468      Key = 0x10028a8 /* U+28A8 BRAILLE PATTERN DOTS-468 */
	K_braille_dots_1468     Key = 0x10028a9 /* U+28A9 BRAILLE PATTERN DOTS-1468 */
	K_braille_dots_2468     Key = 0x10028aa /* U+28AA BRAILLE PATTERN DOTS-2468 */
	K_braille_dots_12468    Key = 0x10028ab /* U+28AB BRAILLE PATTERN DOTS-12468 */
	K_braille_dots_3468     Key = 0x10028ac /* U+28AC BRAILLE PATTERN DOTS-3468 */
	K_braille_dots_13468    Key = 0x10028ad /* U+28AD BRAILLE PATTERN DOTS-13468 */
	K_braille_dots_23468    Key = 0x10028ae /* U+28AE BRAILLE PATTERN DOTS-23468 */
	K_braille_dots_123468   Key = 0x10028af /* U+28AF BRAILLE PATTERN DOTS-123468 */
	K_braille_dots_568      Key = 0x10028b0 /* U+28B0 BRAILLE PATTERN DOTS-568 */
	K_braille_dots_1568     Key = 0x10028b1 /* U+28B1 BRAILLE PATTERN DOTS-1568 */
	K_braille_dots_2568     Key = 0x10028b2 /* U+28B2 BRAILLE PATTERN DOTS-2568 */
	K_braille_dots_12568    Key = 0x10028b3 /* U+28B3 BRAILLE PATTERN DOTS-12568 */
	K_braille_dots_3568     Key = 0x10028b4 /* U+28B4 BRAILLE PATTERN DOTS-3568 */
	K_braille_dots_13568    Key = 0x10028b5 /* U+28B5 BRAILLE PATTERN DOTS-13568 */
	K_braille_dots_23568    Key = 0x10028b6 /* U+28B6 BRAILLE PATTERN DOTS-23568 */
	K_braille_dots_123568   Key = 0x10028b7 /* U+28B7 BRAILLE PATTERN DOTS-123568 */
	K_braille_dots_4568     Key = 0x10028b8 /* U+28B8 BRAILLE PATTERN DOTS-4568 */
	K_braille_dots_14568    Key = 0x10028b9 /* U+28B9 BRAILLE PATTERN DOTS-14568 */
	K_braille_dots_24568    Key = 0x10028ba /* U+28BA BRAILLE PATTERN DOTS-24568 */
	K_braille_dots_124568   Key = 0x10028bb /* U+28BB BRAILLE PATTERN DOTS-124568 */
	K_braille_dots_34568    Key = 0x10028bc /* U+28BC BRAILLE PATTERN DOTS-34568 */
	K_braille_dots_134568   Key = 0x10028bd /* U+28BD BRAILLE PATTERN DOTS-134568 */
	K_braille_dots_234568   Key = 0x10028be /* U+28BE BRAILLE PATTERN DOTS-234568 */
	K_braille_dots_1234568  Key = 0x10028bf /* U+28BF BRAILLE PATTERN DOTS-1234568 */
	K_braille_dots_78       Key = 0x10028c0 /* U+28C0 BRAILLE PATTERN DOTS-78 */
	K_braille_dots_178      Key = 0x10028c1 /* U+28C1 BRAILLE PATTERN DOTS-178 */
	K_braille_dots_278      Key = 0x10028c2 /* U+28C2 BRAILLE PATTERN DOTS-278 */
	K_braille_dots_1278     Key = 0x10028c3 /* U+28C3 BRAILLE PATTERN DOTS-1278 */
	K_braille_dots_378      Key = 0x10028c4 /* U+28C4 BRAILLE PATTERN DOTS-378 */
	K_braille_dots_1378     Key = 0x10028c5 /* U+28C5 BRAILLE PATTERN DOTS-1378 */
	K_braille_dots_2378     Key = 0x10028c6 /* U+28C6 BRAILLE PATTERN DOTS-2378 */
	K_braille_dots_12378    Key = 0x10028c7 /* U+28C7 BRAILLE PATTERN DOTS-12378 */
	K_braille_dots_478      Key = 0x10028c8 /* U+28C8 BRAILLE PATTERN DOTS-478 */
	K_braille_dots_1478     Key = 0x10028c9 /* U+28C9 BRAILLE PATTERN DOTS-1478 */
	K_braille_dots_2478     Key = 0x10028ca /* U+28CA BRAILLE PATTERN DOTS-2478 */
	K_braille_dots_12478    Key = 0x10028cb /* U+28CB BRAILLE PATTERN DOTS-12478 */
	K_braille_dots_3478     Key = 0x10028cc /* U+28CC BRAILLE PATTERN DOTS-3478 */
	K_braille_dots_13478    Key = 0x10028cd /* U+28CD BRAILLE PATTERN DOTS-13478 */
	K_braille_dots_23478    Key = 0x10028ce /* U+28CE BRAILLE PATTERN DOTS-23478 */
	K_braille_dots_123478   Key = 0x10028cf /* U+28CF BRAILLE PATTERN DOTS-123478 */
	K_braille_dots_578      Key = 0x10028d0 /* U+28D0 BRAILLE PATTERN DOTS-578 */
	K_braille_dots_1578     Key = 0x10028d1 /* U+28D1 BRAILLE PATTERN DOTS-1578 */
	K_braille_dots_2578     Key = 0x10028d2 /* U+28D2 BRAILLE PATTERN DOTS-2578 */
	K_braille_dots_12578    Key = 0x10028d3 /* U+28D3 BRAILLE PATTERN DOTS-12578 */
	K_braille_dots_3578     Key = 0x10028d4 /* U+28D4 BRAILLE PATTERN DOTS-3578 */
	K_braille_dots_13578    Key = 0x10028d5 /* U+28D5 BRAILLE PATTERN DOTS-13578 */
	K_braille_dots_23578    Key = 0x10028d6 /* U+28D6 BRAILLE PATTERN DOTS-23578 */
	K_braille_dots_123578   Key = 0x10028d7 /* U+28D7 BRAILLE PATTERN DOTS-123578 */
	K_braille_dots_4578     Key = 0x10028d8 /* U+28D8 BRAILLE PATTERN DOTS-4578 */
	K_braille_dots_14578    Key = 0x10028d9 /* U+28D9 BRAILLE PATTERN DOTS-14578 */
	K_braille_dots_24578    Key = 0x10028da /* U+28DA BRAILLE PATTERN DOTS-24578 */
	K_braille_dots_124578   Key = 0x10028db /* U+28DB BRAILLE PATTERN DOTS-124578 */
	K_braille_dots_34578    Key = 0x10028dc /* U+28DC BRAILLE PATTERN DOTS-34578 */
	K_braille_dots_134578   Key = 0x10028dd /* U+28DD BRAILLE PATTERN DOTS-134578 */
	K_braille_dots_234578   Key = 0x10028de /* U+28DE BRAILLE PATTERN DOTS-234578 */
	K_braille_dots_1234578  Key = 0x10028df /* U+28DF BRAILLE PATTERN DOTS-1234578 */
	K_braille_dots_678      Key = 0x10028e0 /* U+28E0 BRAILLE PATTERN DOTS-678 */
	K_braille_dots_1678     Key = 0x10028e1 /* U+28E1 BRAILLE PATTERN DOTS-1678 */
	K_braille_dots_2678     Key = 0x10028e2 /* U+28E2 BRAILLE PATTERN DOTS-2678 */
	K_braille_dots_12678    Key = 0x10028e3 /* U+28E3 BRAILLE PATTERN DOTS-12678 */
	K_braille_dots_3678     Key = 0x10028e4 /* U+28E4 BRAILLE PATTERN DOTS-3678 */
	K_braille_dots_13678    Key = 0x10028e5 /* U+28E5 BRAILLE PATTERN DOTS-13678 */
	K_braille_dots_23678    Key = 0x10028e6 /* U+28E6 BRAILLE PATTERN DOTS-23678 */
	K_braille_dots_123678   Key = 0x10028e7 /* U+28E7 BRAILLE PATTERN DOTS-123678 */
	K_braille_dots_4678     Key = 0x10028e8 /* U+28E8 BRAILLE PATTERN DOTS-4678 */
	K_braille_dots_14678    Key = 0x10028e9 /* U+28E9 BRAILLE PATTERN DOTS-14678 */
	K_braille_dots_24678    Key = 0x10028ea /* U+28EA BRAILLE PATTERN DOTS-24678 */
	K_braille_dots_124678   Key = 0x10028eb /* U+28EB BRAILLE PATTERN DOTS-124678 */
	K_braille_dots_34678    Key = 0x10028ec /* U+28EC BRAILLE PATTERN DOTS-34678 */
	K_braille_dots_134678   Key = 0x10028ed /* U+28ED BRAILLE PATTERN DOTS-134678 */
	K_braille_dots_234678   Key = 0x10028ee /* U+28EE BRAILLE PATTERN DOTS-234678 */
	K_braille_dots_1234678  Key = 0x10028ef /* U+28EF BRAILLE PATTERN DOTS-1234678 */
	K_braille_dots_5678     Key = 0x10028f0 /* U+28F0 BRAILLE PATTERN DOTS-5678 */
	K_braille_dots_15678    Key = 0x10028f1 /* U+28F1 BRAILLE PATTERN DOTS-15678 */
	K_braille_dots_25678    Key = 0x10028f2 /* U+28F2 BRAILLE PATTERN DOTS-25678 */
	K_braille_dots_125678   Key = 0x10028f3 /* U+28F3 BRAILLE PATTERN DOTS-125678 */
	K_braille_dots_35678    Key = 0x10028f4 /* U+28F4 BRAILLE PATTERN DOTS-35678 */
	K_braille_dots_135678   Key = 0x10028f5 /* U+28F5 BRAILLE PATTERN DOTS-135678 */
	K_braille_dots_235678   Key = 0x10028f6 /* U+28F6 BRAILLE PATTERN DOTS-235678 */
	K_braille_dots_1235678  Key = 0x10028f7 /* U+28F7 BRAILLE PATTERN DOTS-1235678 */
	K_braille_dots_45678    Key = 0x10028f8 /* U+28F8 BRAILLE PATTERN DOTS-45678 */
	K_braille_dots_145678   Key = 0x10028f9 /* U+28F9 BRAILLE PATTERN DOTS-145678 */
	K_braille_dots_245678   Key = 0x10028fa /* U+28FA BRAILLE PATTERN DOTS-245678 */
	K_braille_dots_1245678  Key = 0x10028fb /* U+28FB BRAILLE PATTERN DOTS-1245678 */
	K_braille_dots_345678   Key = 0x10028fc /* U+28FC BRAILLE PATTERN DOTS-345678 */
	K_braille_dots_1345678  Key = 0x10028fd /* U+28FD BRAILLE PATTERN DOTS-1345678 */
	K_braille_dots_2345678  Key = 0x10028fe /* U+28FE BRAILLE PATTERN DOTS-2345678 */
	K_braille_dots_12345678 Key = 0x10028ff /* U+28FF BRAILLE PATTERN DOTS-12345678 */

	/*
	 * Sinhala (http://unicode.org/charts/PDF/U0D80.pdf)
	 * http://www.nongnu.org/sinhala/doc/transliteration/sinhala-transliteration_6.html
	 */

	K_Sinh_ng         Key = 0x1000d82 /* U+0D82 SINHALA SIGN ANUSVARAYA */
	K_Sinh_h2         Key = 0x1000d83 /* U+0D83 SINHALA SIGN VISARGAYA */
	K_Sinh_a          Key = 0x1000d85 /* U+0D85 SINHALA LETTER AYANNA */
	K_Sinh_aa         Key = 0x1000d86 /* U+0D86 SINHALA LETTER AAYANNA */
	K_Sinh_ae         Key = 0x1000d87 /* U+0D87 SINHALA LETTER AEYANNA */
	K_Sinh_aee        Key = 0x1000d88 /* U+0D88 SINHALA LETTER AEEYANNA */
	K_Sinh_i          Key = 0x1000d89 /* U+0D89 SINHALA LETTER IYANNA */
	K_Sinh_ii         Key = 0x1000d8a /* U+0D8A SINHALA LETTER IIYANNA */
	K_Sinh_u          Key = 0x1000d8b /* U+0D8B SINHALA LETTER UYANNA */
	K_Sinh_uu         Key = 0x1000d8c /* U+0D8C SINHALA LETTER UUYANNA */
	K_Sinh_ri         Key = 0x1000d8d /* U+0D8D SINHALA LETTER IRUYANNA */
	K_Sinh_rii        Key = 0x1000d8e /* U+0D8E SINHALA LETTER IRUUYANNA */
	K_Sinh_lu         Key = 0x1000d8f /* U+0D8F SINHALA LETTER ILUYANNA */
	K_Sinh_luu        Key = 0x1000d90 /* U+0D90 SINHALA LETTER ILUUYANNA */
	K_Sinh_e          Key = 0x1000d91 /* U+0D91 SINHALA LETTER EYANNA */
	K_Sinh_ee         Key = 0x1000d92 /* U+0D92 SINHALA LETTER EEYANNA */
	K_Sinh_ai         Key = 0x1000d93 /* U+0D93 SINHALA LETTER AIYANNA */
	K_Sinh_o          Key = 0x1000d94 /* U+0D94 SINHALA LETTER OYANNA */
	K_Sinh_oo         Key = 0x1000d95 /* U+0D95 SINHALA LETTER OOYANNA */
	K_Sinh_au         Key = 0x1000d96 /* U+0D96 SINHALA LETTER AUYANNA */
	K_Sinh_ka         Key = 0x1000d9a /* U+0D9A SINHALA LETTER ALPAPRAANA KAYANNA */
	K_Sinh_kha        Key = 0x1000d9b /* U+0D9B SINHALA LETTER MAHAAPRAANA KAYANNA */
	K_Sinh_ga         Key = 0x1000d9c /* U+0D9C SINHALA LETTER ALPAPRAANA GAYANNA */
	K_Sinh_gha        Key = 0x1000d9d /* U+0D9D SINHALA LETTER MAHAAPRAANA GAYANNA */
	K_Sinh_ng2        Key = 0x1000d9e /* U+0D9E SINHALA LETTER KANTAJA NAASIKYAYA */
	K_Sinh_nga        Key = 0x1000d9f /* U+0D9F SINHALA LETTER SANYAKA GAYANNA */
	K_Sinh_ca         Key = 0x1000da0 /* U+0DA0 SINHALA LETTER ALPAPRAANA CAYANNA */
	K_Sinh_cha        Key = 0x1000da1 /* U+0DA1 SINHALA LETTER MAHAAPRAANA CAYANNA */
	K_Sinh_ja         Key = 0x1000da2 /* U+0DA2 SINHALA LETTER ALPAPRAANA JAYANNA */
	K_Sinh_jha        Key = 0x1000da3 /* U+0DA3 SINHALA LETTER MAHAAPRAANA JAYANNA */
	K_Sinh_nya        Key = 0x1000da4 /* U+0DA4 SINHALA LETTER TAALUJA NAASIKYAYA */
	K_Sinh_jnya       Key = 0x1000da5 /* U+0DA5 SINHALA LETTER TAALUJA SANYOOGA NAAKSIKYAYA */
	K_Sinh_nja        Key = 0x1000da6 /* U+0DA6 SINHALA LETTER SANYAKA JAYANNA */
	K_Sinh_tta        Key = 0x1000da7 /* U+0DA7 SINHALA LETTER ALPAPRAANA TTAYANNA */
	K_Sinh_ttha       Key = 0x1000da8 /* U+0DA8 SINHALA LETTER MAHAAPRAANA TTAYANNA */
	K_Sinh_dda        Key = 0x1000da9 /* U+0DA9 SINHALA LETTER ALPAPRAANA DDAYANNA */
	K_Sinh_ddha       Key = 0x1000daa /* U+0DAA SINHALA LETTER MAHAAPRAANA DDAYANNA */
	K_Sinh_nna        Key = 0x1000dab /* U+0DAB SINHALA LETTER MUURDHAJA NAYANNA */
	K_Sinh_ndda       Key = 0x1000dac /* U+0DAC SINHALA LETTER SANYAKA DDAYANNA */
	K_Sinh_tha        Key = 0x1000dad /* U+0DAD SINHALA LETTER ALPAPRAANA TAYANNA */
	K_Sinh_thha       Key = 0x1000dae /* U+0DAE SINHALA LETTER MAHAAPRAANA TAYANNA */
	K_Sinh_dha        Key = 0x1000daf /* U+0DAF SINHALA LETTER ALPAPRAANA DAYANNA */
	K_Sinh_dhha       Key = 0x1000db0 /* U+0DB0 SINHALA LETTER MAHAAPRAANA DAYANNA */
	K_Sinh_na         Key = 0x1000db1 /* U+0DB1 SINHALA LETTER DANTAJA NAYANNA */
	K_Sinh_ndha       Key = 0x1000db3 /* U+0DB3 SINHALA LETTER SANYAKA DAYANNA */
	K_Sinh_pa         Key = 0x1000db4 /* U+0DB4 SINHALA LETTER ALPAPRAANA PAYANNA */
	K_Sinh_pha        Key = 0x1000db5 /* U+0DB5 SINHALA LETTER MAHAAPRAANA PAYANNA */
	K_Sinh_ba         Key = 0x1000db6 /* U+0DB6 SINHALA LETTER ALPAPRAANA BAYANNA */
	K_Sinh_bha        Key = 0x1000db7 /* U+0DB7 SINHALA LETTER MAHAAPRAANA BAYANNA */
	K_Sinh_ma         Key = 0x1000db8 /* U+0DB8 SINHALA LETTER MAYANNA */
	K_Sinh_mba        Key = 0x1000db9 /* U+0DB9 SINHALA LETTER AMBA BAYANNA */
	K_Sinh_ya         Key = 0x1000dba /* U+0DBA SINHALA LETTER YAYANNA */
	K_Sinh_ra         Key = 0x1000dbb /* U+0DBB SINHALA LETTER RAYANNA */
	K_Sinh_la         Key = 0x1000dbd /* U+0DBD SINHALA LETTER DANTAJA LAYANNA */
	K_Sinh_va         Key = 0x1000dc0 /* U+0DC0 SINHALA LETTER VAYANNA */
	K_Sinh_sha        Key = 0x1000dc1 /* U+0DC1 SINHALA LETTER TAALUJA SAYANNA */
	K_Sinh_ssha       Key = 0x1000dc2 /* U+0DC2 SINHALA LETTER MUURDHAJA SAYANNA */
	K_Sinh_sa         Key = 0x1000dc3 /* U+0DC3 SINHALA LETTER DANTAJA SAYANNA */
	K_Sinh_ha         Key = 0x1000dc4 /* U+0DC4 SINHALA LETTER HAYANNA */
	K_Sinh_lla        Key = 0x1000dc5 /* U+0DC5 SINHALA LETTER MUURDHAJA LAYANNA */
	K_Sinh_fa         Key = 0x1000dc6 /* U+0DC6 SINHALA LETTER FAYANNA */
	K_Sinh_al         Key = 0x1000dca /* U+0DCA SINHALA SIGN AL-LAKUNA */
	K_Sinh_aa2        Key = 0x1000dcf /* U+0DCF SINHALA VOWEL SIGN AELA-PILLA */
	K_Sinh_ae2        Key = 0x1000dd0 /* U+0DD0 SINHALA VOWEL SIGN KETTI AEDA-PILLA */
	K_Sinh_aee2       Key = 0x1000dd1 /* U+0DD1 SINHALA VOWEL SIGN DIGA AEDA-PILLA */
	K_Sinh_i2         Key = 0x1000dd2 /* U+0DD2 SINHALA VOWEL SIGN KETTI IS-PILLA */
	K_Sinh_ii2        Key = 0x1000dd3 /* U+0DD3 SINHALA VOWEL SIGN DIGA IS-PILLA */
	K_Sinh_u2         Key = 0x1000dd4 /* U+0DD4 SINHALA VOWEL SIGN KETTI PAA-PILLA */
	K_Sinh_uu2        Key = 0x1000dd6 /* U+0DD6 SINHALA VOWEL SIGN DIGA PAA-PILLA */
	K_Sinh_ru2        Key = 0x1000dd8 /* U+0DD8 SINHALA VOWEL SIGN GAETTA-PILLA */
	K_Sinh_e2         Key = 0x1000dd9 /* U+0DD9 SINHALA VOWEL SIGN KOMBUVA */
	K_Sinh_ee2        Key = 0x1000dda /* U+0DDA SINHALA VOWEL SIGN DIGA KOMBUVA */
	K_Sinh_ai2        Key = 0x1000ddb /* U+0DDB SINHALA VOWEL SIGN KOMBU DEKA */
	K_Sinh_o2         Key = 0x1000ddc /* U+0DDC SINHALA VOWEL SIGN KOMBUVA HAA AELA-PILLA */
	K_Sinh_oo2        Key = 0x1000ddd /* U+0DDD SINHALA VOWEL SIGN KOMBUVA HAA DIGA AELA-PILLA */
	K_Sinh_au2        Key = 0x1000dde /* U+0DDE SINHALA VOWEL SIGN KOMBUVA HAA GAYANUKITTA */
	K_Sinh_lu2        Key = 0x1000ddf /* U+0DDF SINHALA VOWEL SIGN GAYANUKITTA */
	K_Sinh_ruu2       Key = 0x1000df2 /* U+0DF2 SINHALA VOWEL SIGN DIGA GAETTA-PILLA */
	K_Sinh_luu2       Key = 0x1000df3 /* U+0DF3 SINHALA VOWEL SIGN DIGA GAYANUKITTA */
	K_Sinh_kunddaliya Key = 0x1000df4 /* U+0DF4 SINHALA PUNCTUATION KUNDDALIYA */
	/*
	 * XFree86 vendor specific keysyms.
	 *
	 * The XFree86 keysym range is 0x10080001 - 0x1008ffff.
	 *
	 * The XF86 set of keysyms is a catch-all set of defines for keysyms found
	 * on various multimedia keyboards. Originally specific to XFree86 they have
	 * been been adopted over time and are considered a "standard" part of X
	 * keysym definitions.
	 * XFree86 never properly commented these keysyms, so we have done our
	 * best to explain the semantic meaning of these keys.
	 *
	 * XFree86 has removed their mail archives of the period, that might have
	 * shed more light on some of these definitions. Until/unless we resurrect
	 * these archives, these are from memory and usage.
	 */

	/*
	 * ModeLock
	 *
	 * This one is old, and not really used any more since XKB offers this
	 * functionality.
	 */

	K_XF86ModeLock Key = 0x1008ff01 /* Mode Switch Lock */

	/* Backlight controls. */
	K_XF86MonBrightnessUp    Key = 0x1008ff02 /* Monitor/panel brightness */
	K_XF86MonBrightnessDown  Key = 0x1008ff03 /* Monitor/panel brightness */
	K_XF86KbdLightOnOff      Key = 0x1008ff04 /* Keyboards may be lit     */
	K_XF86KbdBrightnessUp    Key = 0x1008ff05 /* Keyboards may be lit     */
	K_XF86KbdBrightnessDown  Key = 0x1008ff06 /* Keyboards may be lit     */
	K_XF86MonBrightnessCycle Key = 0x1008ff07 /* Monitor/panel brightness */

	/*
	 * Keys found on some "Internet" keyboards.
	 */
	K_XF86Standby          Key = 0x1008ff10 /* System into standby mode   */
	K_XF86AudioLowerVolume Key = 0x1008ff11 /* Volume control down        */
	K_XF86AudioMute        Key = 0x1008ff12 /* Mute sound from the system */
	K_XF86AudioRaiseVolume Key = 0x1008ff13 /* Volume control up          */
	K_XF86AudioPlay        Key = 0x1008ff14 /* Start playing of audio >   */
	K_XF86AudioStop        Key = 0x1008ff15 /* Stop playing audio         */
	K_XF86AudioPrev        Key = 0x1008ff16 /* Previous track             */
	K_XF86AudioNext        Key = 0x1008ff17 /* Next track                 */
	K_XF86HomePage         Key = 0x1008ff18 /* Display user's home page   */
	K_XF86Mail             Key = 0x1008ff19 /* Invoke user's mail program */
	K_XF86Start            Key = 0x1008ff1a /* Start application          */
	K_XF86Search           Key = 0x1008ff1b /* Search                     */
	K_XF86AudioRecord      Key = 0x1008ff1c /* Record audio application   */

	/* These are sometimes found on PDA's (e.g. Palm, PocketPC or elsewhere)   */
	K_XF86Calculator     Key = 0x1008ff1d /* Invoke calculator program  */
	K_XF86Memo           Key = 0x1008ff1e /* Invoke Memo taking program */
	K_XF86ToDoList       Key = 0x1008ff1f /* Invoke To Do List program  */
	K_XF86Calendar       Key = 0x1008ff20 /* Invoke Calendar program    */
	K_XF86PowerDown      Key = 0x1008ff21 /* Deep sleep the system      */
	K_XF86ContrastAdjust Key = 0x1008ff22 /* Adjust screen contrast     */
	K_XF86RockerUp       Key = 0x1008ff23 /* Rocker switches exist up   */
	K_XF86RockerDown     Key = 0x1008ff24 /* and down                   */
	K_XF86RockerEnter    Key = 0x1008ff25 /* and let you press them     */

	/* Some more "Internet" keyboard symbols */
	K_XF86Back             Key = 0x1008ff26 /* Like back on a browser     */
	K_XF86Forward          Key = 0x1008ff27 /* Like forward on a browser  */
	K_XF86Stop             Key = 0x1008ff28 /* Stop current operation     */
	K_XF86Refresh          Key = 0x1008ff29 /* Refresh the page           */
	K_XF86PowerOff         Key = 0x1008ff2a /* Power off system entirely  */
	K_XF86WakeUp           Key = 0x1008ff2b /* Wake up system from sleep  */
	K_XF86Eject            Key = 0x1008ff2c /* Eject device (e.g. DVD)    */
	K_XF86ScreenSaver      Key = 0x1008ff2d /* Invoke screensaver         */
	K_XF86WWW              Key = 0x1008ff2e /* Invoke web browser         */
	K_XF86Sleep            Key = 0x1008ff2f /* Put system to sleep        */
	K_XF86Favorites        Key = 0x1008ff30 /* Show favorite locations    */
	K_XF86AudioPause       Key = 0x1008ff31 /* Pause audio playing        */
	K_XF86AudioMedia       Key = 0x1008ff32 /* Launch media collection app */
	K_XF86MyComputer       Key = 0x1008ff33 /* Display "My Computer" window */
	K_XF86VendorHome       Key = 0x1008ff34 /* Display vendor home web site */
	K_XF86LightBulb        Key = 0x1008ff35 /* Light bulb keys exist       */
	K_XF86Shop             Key = 0x1008ff36 /* Display shopping web site   */
	K_XF86History          Key = 0x1008ff37 /* Show history of web surfing */
	K_XF86OpenURL          Key = 0x1008ff38 /* Open selected URL           */
	K_XF86AddFavorite      Key = 0x1008ff39 /* Add URL to favorites list   */
	K_XF86HotLinks         Key = 0x1008ff3a /* Show "hot" links            */
	K_XF86BrightnessAdjust Key = 0x1008ff3b /* Invoke brightness adj. UI   */
	K_XF86Finance          Key = 0x1008ff3c /* Display financial site      */
	K_XF86Community        Key = 0x1008ff3d /* Display user's community    */
	K_XF86AudioRewind      Key = 0x1008ff3e /* "rewind" audio track        */
	K_XF86BackForward      Key = 0x1008ff3f /* ??? */
	K_XF86Launch0          Key = 0x1008ff40 /* Launch Application          */
	K_XF86Launch1          Key = 0x1008ff41 /* Launch Application          */
	K_XF86Launch2          Key = 0x1008ff42 /* Launch Application          */
	K_XF86Launch3          Key = 0x1008ff43 /* Launch Application          */
	K_XF86Launch4          Key = 0x1008ff44 /* Launch Application          */
	K_XF86Launch5          Key = 0x1008ff45 /* Launch Application          */
	K_XF86Launch6          Key = 0x1008ff46 /* Launch Application          */
	K_XF86Launch7          Key = 0x1008ff47 /* Launch Application          */
	K_XF86Launch8          Key = 0x1008ff48 /* Launch Application          */
	K_XF86Launch9          Key = 0x1008ff49 /* Launch Application          */
	K_XF86LaunchA          Key = 0x1008ff4a /* Launch Application          */
	K_XF86LaunchB          Key = 0x1008ff4b /* Launch Application          */
	K_XF86LaunchC          Key = 0x1008ff4c /* Launch Application          */
	K_XF86LaunchD          Key = 0x1008ff4d /* Launch Application          */
	K_XF86LaunchE          Key = 0x1008ff4e /* Launch Application          */
	K_XF86LaunchF          Key = 0x1008ff4f /* Launch Application          */

	K_XF86ApplicationLeft  Key = 0x1008ff50 /* switch to application, left */
	K_XF86ApplicationRight Key = 0x1008ff51 /* switch to application, right*/
	K_XF86Book             Key = 0x1008ff52 /* Launch bookreader           */
	K_XF86CD               Key = 0x1008ff53 /* Launch CD/DVD player        */
	K_XF86MediaSelectCD    Key = 0x1008ff53 /* Alias for XF86CD            */
	K_XF86Calculater       Key = 0x1008ff54 /* Launch Calculater           */
	K_XF86Clear            Key = 0x1008ff55 /* Clear window, screen        */
	K_XF86Close            Key = 0x1008ff56 /* Close window                */
	K_XF86Copy             Key = 0x1008ff57 /* Copy selection              */
	K_XF86Cut              Key = 0x1008ff58 /* Cut selection               */
	K_XF86Display          Key = 0x1008ff59 /* Output switch key           */
	K_XF86DOS              Key = 0x1008ff5a /* Launch DOS (emulation)      */
	K_XF86Documents        Key = 0x1008ff5b /* Open documents window       */
	K_XF86Excel            Key = 0x1008ff5c /* Launch spread sheet         */
	K_XF86Explorer         Key = 0x1008ff5d /* Launch file explorer        */
	K_XF86Game             Key = 0x1008ff5e /* Launch game                 */
	K_XF86Go               Key = 0x1008ff5f /* Go to URL                   */
	K_XF86iTouch           Key = 0x1008ff60 /* Logitech iTouch- don't use  */
	K_XF86LogOff           Key = 0x1008ff61 /* Log off system              */
	K_XF86Market           Key = 0x1008ff62 /* ??                          */
	K_XF86Meeting          Key = 0x1008ff63 /* enter meeting in calendar   */
	K_XF86MenuKB           Key = 0x1008ff65 /* distinguish keyboard from PB */
	K_XF86MenuPB           Key = 0x1008ff66 /* distinguish PB from keyboard */
	K_XF86MySites          Key = 0x1008ff67 /* Favourites                  */
	K_XF86New              Key = 0x1008ff68 /* New (folder, document...    */
	K_XF86News             Key = 0x1008ff69 /* News                        */
	K_XF86OfficeHome       Key = 0x1008ff6a /* Office home (old Staroffice)*/
	K_XF86Open             Key = 0x1008ff6b /* Open                        */
	K_XF86Option           Key = 0x1008ff6c /* ?? */
	K_XF86Paste            Key = 0x1008ff6d /* Paste                       */
	K_XF86Phone            Key = 0x1008ff6e /* Launch phone; dial number   */
	K_XF86Q                Key = 0x1008ff70 /* Compaq's Q - don't use      */
	K_XF86Reply            Key = 0x1008ff72 /* Reply e.g., mail            */
	K_XF86Reload           Key = 0x1008ff73 /* Reload web page, file, etc. */
	K_XF86RotateWindows    Key = 0x1008ff74 /* Rotate windows e.g. xrandr  */
	K_XF86RotationPB       Key = 0x1008ff75 /* don't use                   */
	K_XF86RotationKB       Key = 0x1008ff76 /* don't use                   */
	K_XF86Save             Key = 0x1008ff77 /* Save (file, document, state */
	K_XF86ScrollUp         Key = 0x1008ff78 /* Scroll window/contents up   */
	K_XF86ScrollDown       Key = 0x1008ff79 /* Scrool window/contentd down */
	K_XF86ScrollClick      Key = 0x1008ff7a /* Use XKB mousekeys instead   */
	K_XF86Send             Key = 0x1008ff7b /* Send mail, file, object     */
	K_XF86Spell            Key = 0x1008ff7c /* Spell checker               */
	K_XF86SplitScreen      Key = 0x1008ff7d /* Split window or screen      */
	K_XF86Support          Key = 0x1008ff7e /* Get support (??)            */
	K_XF86TaskPane         Key = 0x1008ff7f /* Show tasks */
	K_XF86Terminal         Key = 0x1008ff80 /* Launch terminal emulator    */
	K_XF86Tools            Key = 0x1008ff81 /* toolbox of desktop/app.     */
	K_XF86Travel           Key = 0x1008ff82 /* ?? */
	K_XF86UserPB           Key = 0x1008ff84 /* ?? */
	K_XF86User1KB          Key = 0x1008ff85 /* ?? */
	K_XF86User2KB          Key = 0x1008ff86 /* ?? */
	K_XF86Video            Key = 0x1008ff87 /* Launch video player       */
	K_XF86WheelButton      Key = 0x1008ff88 /* button from a mouse wheel */
	K_XF86Word             Key = 0x1008ff89 /* Launch word processor     */
	K_XF86Xfer             Key = 0x1008ff8a
	K_XF86ZoomIn           Key = 0x1008ff8b /* zoom in view, map, etc.   */
	K_XF86ZoomOut          Key = 0x1008ff8c /* zoom out view, map, etc.  */

	K_XF86Away        Key = 0x1008ff8d /* mark yourself as away     */
	K_XF86Messenger   Key = 0x1008ff8e /* as in instant messaging   */
	K_XF86WebCam      Key = 0x1008ff8f /* Launch web camera app.    */
	K_XF86MailForward Key = 0x1008ff90 /* Forward in mail           */
	K_XF86Pictures    Key = 0x1008ff91 /* Show pictures             */
	K_XF86Music       Key = 0x1008ff92 /* Launch music application  */

	K_XF86Battery   Key = 0x1008ff93 /* Display battery information */
	K_XF86Bluetooth Key = 0x1008ff94 /* Enable/disable Bluetooth    */
	K_XF86WLAN      Key = 0x1008ff95 /* Enable/disable WLAN         */
	K_XF86UWB       Key = 0x1008ff96 /* Enable/disable UWB	    */

	K_XF86AudioForward    Key = 0x1008ff97 /* fast-forward audio track    */
	K_XF86AudioRepeat     Key = 0x1008ff98 /* toggle repeat mode          */
	K_XF86AudioRandomPlay Key = 0x1008ff99 /* toggle shuffle mode         */
	K_XF86Subtitle        Key = 0x1008ff9a /* cycle through subtitle      */
	K_XF86AudioCycleTrack Key = 0x1008ff9b /* cycle through audio tracks  */
	K_XF86CycleAngle      Key = 0x1008ff9c /* cycle through angles        */
	K_XF86FrameBack       Key = 0x1008ff9d /* video: go one frame back    */
	K_XF86FrameForward    Key = 0x1008ff9e /* video: go one frame forward */
	K_XF86Time            Key = 0x1008ff9f /* display, or shows an entry for time seeking */
	K_XF86Select          Key = 0x1008ffa0 /* Select button on joypads and remotes */
	K_XF86View            Key = 0x1008ffa1 /* Show a view options/properties */
	K_XF86TopMenu         Key = 0x1008ffa2 /* Go to a top-level menu in a video */

	K_XF86Red    Key = 0x1008ffa3 /* Red button                  */
	K_XF86Green  Key = 0x1008ffa4 /* Green button                */
	K_XF86Yellow Key = 0x1008ffa5 /* Yellow button               */
	K_XF86Blue   Key = 0x1008ffa6 /* Blue button                 */

	K_XF86Suspend        Key = 0x1008ffa7 /* Sleep to RAM                */
	K_XF86Hibernate      Key = 0x1008ffa8 /* Sleep to disk               */
	K_XF86TouchpadToggle Key = 0x1008ffa9 /* Toggle between touchpad/trackstick */
	K_XF86TouchpadOn     Key = 0x1008ffb0 /* The touchpad got switched on */
	K_XF86TouchpadOff    Key = 0x1008ffb1 /* The touchpad got switched off */

	K_XF86AudioMicMute Key = 0x1008ffb2 /* Mute the Mic from the system */

	K_XF86Keyboard Key = 0x1008ffb3 /* User defined keyboard related action */

	K_XF86WWAN   Key = 0x1008ffb4 /* Toggle WWAN (LTE, UMTS, etc.) radio */
	K_XF86RFKill Key = 0x1008ffb5 /* Toggle radios on/off */

	K_XF86AudioPreset Key = 0x1008ffb6 /* Select equalizer preset, e.g. theatre-mode */

	K_XF86RotationLockToggle Key = 0x1008ffb7 /* Toggle screen rotation lock on/off */

	K_XF86FullScreen Key = 0x1008ffb8 /* Toggle fullscreen */

	/* Keys for special action keys (hot keys) */
	/* Virtual terminals on some operating systems */
	K_XF86Switch_VT_1  Key = 0x1008fe01
	K_XF86Switch_VT_2  Key = 0x1008fe02
	K_XF86Switch_VT_3  Key = 0x1008fe03
	K_XF86Switch_VT_4  Key = 0x1008fe04
	K_XF86Switch_VT_5  Key = 0x1008fe05
	K_XF86Switch_VT_6  Key = 0x1008fe06
	K_XF86Switch_VT_7  Key = 0x1008fe07
	K_XF86Switch_VT_8  Key = 0x1008fe08
	K_XF86Switch_VT_9  Key = 0x1008fe09
	K_XF86Switch_VT_10 Key = 0x1008fe0a
	K_XF86Switch_VT_11 Key = 0x1008fe0b
	K_XF86Switch_VT_12 Key = 0x1008fe0c

	K_XF86Ungrab        Key = 0x1008fe20 /* force ungrab               */
	K_XF86ClearGrab     Key = 0x1008fe21 /* kill application with grab */
	K_XF86Next_VMode    Key = 0x1008fe22 /* next video mode available  */
	K_XF86Prev_VMode    Key = 0x1008fe23 /* prev. video mode available */
	K_XF86LogWindowTree Key = 0x1008fe24 /* print window tree to log   */
	K_XF86LogGrabInfo   Key = 0x1008fe25 /* print all active grabs to log */

	/*
	 * Reserved range for evdev symbols: 0x10081000-0x10081FFF
	 *
	 * Key symbols within this range are intended for a 1:1 mapping to the
	 * Linux kernel input-event-codes.h file:
	 * - Keysym value: `_EVDEVK(kernel value)`
	 * - Keysym name: it must be prefixed by `XKB_K_XF86`. The recommended *default*
	 *   name uses the following pattern: `XKB_K_XF86CamelCaseKernelName`. CamelCasing
	 *   is done with a human control as last authority, e.g. see VOD instead of Vod
	 *   for the Video on Demand key. In case that the kernel key name is too
	 *   ambiguous, it is recommended to create a more explicit name with the
	 *   following guidelines:
	 *   1. Names should be mnemonic.
	 *   2. Names should avoid acronyms, unless sufficiently common and documented
	 *      in a comment.
	 *   3. Names should identify a function.
	 *   4. Keysyms share a common namespace, so keys with a generic name should
	 *      denote a generic function, otherwise it should be renamed to add context.
	 *      E.g. `K_OK` has the associated keysym `XKB_K_XF86OK`, while `K_TITLE`
	 *      (used to open a menu to select a chapter of a media) is associated to
	 *      the keysym `XKB_K_XF86MediaTitleMenu` to avoid ambiguity.
	 *   5. Keysyms should support designing *portable* applications, therefore
	 *      their names should be self-explaining to facilitate finding them and
	 *      to avoid misuse.
	 *   6. The “HID usage tables for USB” can be used if there is an unambiguous
	 *      mapping. See:
	 *      - Reference document: https://usb.org/document-library/hid-usage-tables-16
	 *      - Mapping in the Linux source file: `drivers/hid/hid-input.c` as of 2025-07
	 *
	 * For example, the kernel
	 *     #defineK_MACRO_RECORD_START	0x2b0
	 * effectively ends up as:
	 *K_XF86MacroRecordStart Key = 0x100812b0
	 *
	 * For historical reasons, some keysyms within the reserved range will be
	 * missing, most notably all "normal" keys that are mapped through default
	 * XKB layouts (e.g.K_Q).
	 *
	 * The format for #defines is strict:
	 *
	 *     #define XKB_K_XF86Foo<spaces…>_EVDEVK(0xABC)<spaces…> |* kverK_FOO *|
	 *     #define XKB_K_XF86Foo<spaces…>_EVDEVK(0xABC)<spaces…> |* Alias for XF86Bar *|
	 *     #define XKB_K_XF86Foo<spaces…>_EVDEVK(0xABC)<spaces…> |* Deprecated alias for XF86Bar *|
	 *
	 * Where
	 * - alignment by spaces
	 * - the _EVDEVK macro must be used
	 * - the hex code must be in uppercase hex
	 * - the kernel version (kver) is in the form v5.10
	 * - kver and key name are within a slash-star comment (a pipe is used in
	 *   this example for technical reasons)
	 * These #defines are parsed by scripts. Do not stray from the given format.
	 *
	 * Where the evdev keycode is mapped to a different symbol, please add a
	 * comment line starting with Use: but otherwise the same format, e.g.
	 *  Use: XKB_K_XF86RotationLockToggle	_EVDEVK(0x231)		   v4.16K_ROTATE_LOCK_TOGGLE
	 *
	 */
	/* Use: XKB_K_XF86Eject                    _EVDEVK(0x0a2)K_EJECTCLOSECD */
	/* Use: XKB_K_XF86New                      _EVDEVK(0x0b5)     v2.6.14K_NEW */
	/* Use: XKB_K_Redo                         _EVDEVK(0x0b6)     v2.6.14K_REDO */
	/*K_DASHBOARD has been mapped to LaunchB in xkeyboard-config since 2011 */
	/* Use: XKB_K_XF86LaunchB                  _EVDEVK(0x0cc)     v2.6.28K_DASHBOARD */
	/* Use: XKB_K_XF86Display                  _EVDEVK(0x0e3)     v2.6.12K_SWITCHVIDEOMODE */
	/* Use: XKB_K_XF86KbdLightOnOff            _EVDEVK(0x0e4)     v2.6.12K_KBDILLUMTOGGLE */
	/* Use: XKB_K_XF86KbdBrightnessDown        _EVDEVK(0x0e5)     v2.6.12K_KBDILLUMDOWN */
	/* Use: XKB_K_XF86KbdBrightnessUp          _EVDEVK(0x0e6)     v2.6.12K_KBDILLUMUP */
	/* Use: XKB_K_XF86Send                     _EVDEVK(0x0e7)     v2.6.14K_SEND */
	/* Use: XKB_K_XF86Reply                    _EVDEVK(0x0e8)     v2.6.14K_REPLY */
	/* Use: XKB_K_XF86MailForward              _EVDEVK(0x0e9)     v2.6.14K_FORWARDMAIL */
	/* Use: XKB_K_XF86Save                     _EVDEVK(0x0ea)     v2.6.14K_SAVE */
	/* Use: XKB_K_XF86Documents                _EVDEVK(0x0eb)     v2.6.14K_DOCUMENTS */
	/* Use: XKB_K_XF86Battery                  _EVDEVK(0x0ec)     v2.6.17K_BATTERY */
	/* Use: XKB_K_XF86Bluetooth                _EVDEVK(0x0ed)     v2.6.19K_BLUETOOTH */
	/* Use: XKB_K_XF86WLAN                     _EVDEVK(0x0ee)     v2.6.19K_WLAN */
	/* Use: XKB_K_XF86UWB                      _EVDEVK(0x0ef)     v2.6.24K_UWB */
	/* Use: XKB_K_XF86Next_VMode               _EVDEVK(0x0f1)     v2.6.23K_VIDEO_NEXT */
	/* Use: XKB_K_XF86Prev_VMode               _EVDEVK(0x0f2)     v2.6.23K_VIDEO_PREV */
	/* Use: XKB_K_XF86MonBrightnessCycle       _EVDEVK(0x0f3)     v2.6.23K_BRIGHTNESS_CYCLE */
	K_XF86BrightnessAuto Key = 0x100810f4 /* v3.16K_BRIGHTNESS_AUTO */
	K_XF86DisplayOff     Key = 0x100810f5 /* v2.6.23K_DISPLAY_OFF */
	/* Use: XKB_K_XF86WWAN                     _EVDEVK(0x0f6)     v3.13K_WWAN */
	/* Use: XKB_K_XF86RFKill                   _EVDEVK(0x0f7)     v2.6.33K_RFKILL */
	/* Use: XKB_K_XF86AudioMicMute             _EVDEVK(0x0f8)     v3.1K_MICMUTE */
	K_XF86OK Key = 0x10081160 /* v2.5.26K_OK */
	/* Use: XKB_K_XF86Select                   _EVDEVK(0x161)     v2.5.26K_SELECT */
	K_XF86GoTo Key = 0x10081162 /* v2.5.26K_GOTO */
	/* Use: XKB_K_XF86Clear                    _EVDEVK(0x163)     v2.5.26K_CLEAR */
	/* TODO: Unclear function                    _EVDEVK(0x164)     v2.5.26K_POWER2 */
	/* Use: XKB_K_XF86Option                   _EVDEVK(0x165)     v2.5.26K_OPTION */
	K_XF86Info Key = 0x10081166 /* v2.5.26K_INFO */
	/* Use: XKB_K_XF86Time                     _EVDEVK(0x167)     v2.5.26K_TIME */
	K_XF86VendorLogo Key = 0x10081168 /* v2.5.26K_VENDOR */
	/* TODO: unclear function                    _EVDEVK(0x169)     v2.5.26K_ARCHIVE */
	K_XF86MediaSelectProgramGuide Key = 0x1008116a /* v2.5.26K_PROGRAM */
	/* Use: XKB_K_XF86NextFavorite             _EVDEVK(0x16b)     v2.5.26K_CHANNEL */
	/* Use: XKB_K_XF86Favorites                _EVDEVK(0x16c)     v2.5.26K_FAVORITES */
	/* Use: XKB_K_XF86MediaSelectProgramGuide  _EVDEVK(0x16d)     v2.5.26K_EPG */
	K_XF86MediaSelectHome Key = 0x1008116e /* v2.5.26K_PVR */
	/* TODO: Multimedia Home Platform            _EVDEVK(0x16f)     v2.5.26K_MHP */
	K_XF86MediaLanguageMenu Key = 0x10081170 /* v2.5.26K_LANGUAGE */
	K_XF86MediaTitleMenu    Key = 0x10081171 /* v2.5.26K_TITLE */
	/* Use: XKB_K_XF86Subtitle                 _EVDEVK(0x172)     v2.5.26K_SUBTITLE */
	/* Use: XKB_K_XF86CycleAngle               _EVDEVK(0x173)     v2.5.26K_ANGLE */
	/* Use: XKB_K_XF86FullScreen               _EVDEVK(0x174)     v5.1K_FULL_SCREEN */
	K_XF86AudioChannelMode Key = 0x10081175 /* v2.5.26K_MODE */
	/* Use: XKB_K_XF86Keyboard                 _EVDEVK(0x176)     v2.5.26K_KEYBOARD */
	K_XF86AspectRatio          Key = 0x10081177 /* v5.1K_ASPECT_RATIO */
	K_XF86MediaSelectPC        Key = 0x10081178 /* v2.5.26K_PC */
	K_XF86MediaSelectTV        Key = 0x10081179 /* v2.5.26K_TV */
	K_XF86MediaSelectCable     Key = 0x1008117a /* v2.5.26K_TV2 */
	K_XF86MediaSelectVCR       Key = 0x1008117b /* v2.5.26K_VCR */
	K_XF86MediaSelectVCRPlus   Key = 0x1008117c /* v2.5.26K_VCR2 */
	K_XF86MediaSelectSatellite Key = 0x1008117d /* v2.5.26K_SAT */
	/* TODO: unclear media selector              _EVDEVK(0x17e)     v2.5.26K_SAT2 */
	/* Use: XKB_K_XF86MediaSelectCD            _EVDEVK(0x17f)     v2.5.26K_CD */
	K_XF86MediaSelectTape      Key = 0x10081180 /* v2.5.26K_TAPE */
	K_XF86MediaSelectRadio     Key = 0x10081181 /* v2.5.26K_RADIO */
	K_XF86MediaSelectTuner     Key = 0x10081182 /* v2.5.26K_TUNER */
	K_XF86MediaPlayer          Key = 0x10081183 /* v2.5.26K_PLAYER */
	K_XF86MediaSelectTeletext  Key = 0x10081184 /* v2.5.26K_TEXT */
	K_XF86DVD                  Key = 0x10081185 /* v2.5.26K_DVD */
	K_XF86MediaSelectDVD       Key = 0x10081185 /* Alias for XF86DVD */
	K_XF86MediaSelectAuxiliary Key = 0x10081186 /* v2.5.26K_AUX */
	/* TODO: unclear media selector              _EVDEVK(0x187)     v2.5.26K_MP3 */
	K_XF86Audio Key = 0x10081188 /* v2.5.26K_AUDIO */
	/* Use: XKB_K_XF86Video                    _EVDEVK(0x189)     v2.5.26K_VIDEO */
	/* TODO: unclear function                    _EVDEVK(0x18a)     v2.5.26K_DIRECTORY */
	/* TODO: unclear function                    _EVDEVK(0x18b)     v2.5.26K_LIST */
	/* Use: XKB_K_XF86Memo                     _EVDEVK(0x18c)     v2.5.26K_MEMO */
	/* Use: XKB_K_XF86Calendar                 _EVDEVK(0x18d)     v2.5.26K_CALENDAR */
	/* Use: XKB_K_XF86Red                      _EVDEVK(0x18e)     v2.5.26K_RED */
	/* Use: XKB_K_XF86Green                    _EVDEVK(0x18f)     v2.5.26K_GREEN */
	/* Use: XKB_K_XF86Yellow                   _EVDEVK(0x190)     v2.5.26K_YELLOW */
	/* Use: XKB_K_XF86Blue                     _EVDEVK(0x191)     v2.5.26K_BLUE */
	K_XF86ChannelUp   Key = 0x10081192 /* v2.5.26K_CHANNELUP */
	K_XF86ChannelDown Key = 0x10081193 /* v2.5.26K_CHANNELDOWN */
	/* TODO: unclear function                    _EVDEVK(0x194)     v2.5.26K_FIRST */
	/* TODO: unclear function                    _EVDEVK(0x195)     v2.5.26K_LAST */
	/* TODO: unclear function                    _EVDEVK(0x196)     v2.5.26K_AB */
	/* TODO: unclear function                    _EVDEVK(0x197)     v2.5.26K_NEXT */
	/* TODO: unclear function                    _EVDEVK(0x198)     v2.5.26K_RESTART */
	K_XF86MediaPlaySlow Key = 0x10081199 /* v2.5.26K_SLOW */
	/* Use: XKB_K_XF86AudioRandomPlay          _EVDEVK(0x19a)     v2.5.26K_SHUFFLE */
	K_XF86Break Key = 0x1008119b /* v2.5.26K_BREAK */
	/* TODO: unclear function                    _EVDEVK(0x19c)     v2.5.26K_PREVIOUS */
	K_XF86NumberEntryMode Key = 0x1008119d /* v2.5.26K_DIGITS */
	/* TODO: unclear function                    _EVDEVK(0x19e)     v2.5.26K_TEEN */
	/* TODO: unclear function (Twenties?)        _EVDEVK(0x19f)     v2.5.26K_TWEN */
	K_XF86VideoPhone Key = 0x100811a0 /* v2.6.20K_VIDEOPHONE */
	/* Use: XKB_K_XF86Game                     _EVDEVK(0x1a1)     v2.6.20K_GAMES */
	/* Use: XKB_K_XF86ZoomIn                   _EVDEVK(0x1a2)     v2.6.20K_ZOOMIN */
	/* Use: XKB_K_XF86ZoomOut                  _EVDEVK(0x1a3)     v2.6.20K_ZOOMOUT */
	K_XF86ZoomReset Key = 0x100811a4 /* v2.6.20K_ZOOMRESET */
	/* Use: XKB_K_XF86Word                     _EVDEVK(0x1a5)     v2.6.20K_WORDPROCESSOR */
	K_XF86Editor Key = 0x100811a6 /* v2.6.20K_EDITOR */
	/* Use: XKB_K_XF86Excel                    _EVDEVK(0x1a7)     v2.6.20K_SPREADSHEET */
	K_XF86GraphicsEditor Key = 0x100811a8 /* v2.6.20K_GRAPHICSEDITOR */
	K_XF86Presentation   Key = 0x100811a9 /* v2.6.20K_PRESENTATION */
	K_XF86Database       Key = 0x100811aa /* v2.6.20K_DATABASE */
	/* Use: XKB_K_XF86News                     _EVDEVK(0x1ab)     v2.6.20K_NEWS */
	K_XF86Voicemail   Key = 0x100811ac /* v2.6.20K_VOICEMAIL */
	K_XF86Addressbook Key = 0x100811ad /* v2.6.20K_ADDRESSBOOK */
	/* Use: XKB_K_XF86Messenger                _EVDEVK(0x1ae)     v2.6.20K_MESSENGER */
	K_XF86DisplayToggle Key = 0x100811af /* v2.6.20K_DISPLAYTOGGLE */
	K_XF86SpellCheck    Key = 0x100811b0 /* v2.6.24K_SPELLCHECK */
	/* Use: XKB_K_XF86LogOff                   _EVDEVK(0x1b1)     v2.6.24K_LOGOFF */
	/* Use: XKB_K_dollar                       _EVDEVK(0x1b2)     v2.6.24K_DOLLAR */
	/* Use: XKB_K_EuroSign                     _EVDEVK(0x1b3)     v2.6.24K_EURO */
	/* Use: XKB_K_XF86FrameBack                _EVDEVK(0x1b4)     v2.6.24K_FRAMEBACK */
	/* Use: XKB_K_XF86FrameForward             _EVDEVK(0x1b5)     v2.6.24K_FRAMEFORWARD */
	K_XF86ContextMenu        Key = 0x100811b6 /* v2.6.24K_CONTEXT_MENU */
	K_XF86MediaRepeat        Key = 0x100811b7 /* v2.6.26K_MEDIA_REPEAT */
	K_XF8610ChannelsUp       Key = 0x100811b8 /* v2.6.38K_10CHANNELSUP */
	K_XF8610ChannelsDown     Key = 0x100811b9 /* v2.6.38K_10CHANNELSDOWN */
	K_XF86Images             Key = 0x100811ba /* v2.6.39K_IMAGES */
	K_XF86NotificationCenter Key = 0x100811bc /* v5.10K_NOTIFICATION_CENTER */
	K_XF86PickupPhone        Key = 0x100811bd /* v5.10K_PICKUP_PHONE */
	K_XF86HangupPhone        Key = 0x100811be /* v5.10K_HANGUP_PHONE */
	K_XF86Fn                 Key = 0x100811d0 /*K_FN */
	K_XF86Fn_Esc             Key = 0x100811d1 /*K_FN_ESC */
	K_XF86FnRightShift       Key = 0x100811e5 /* v5.10K_FN_RIGHT_SHIFT */
	/* Use: XKB_K_braille_dot_1                _EVDEVK(0x1f1)     v2.6.17K_BRL_DOT1 */
	/* Use: XKB_K_braille_dot_2                _EVDEVK(0x1f2)     v2.6.17K_BRL_DOT2 */
	/* Use: XKB_K_braille_dot_3                _EVDEVK(0x1f3)     v2.6.17K_BRL_DOT3 */
	/* Use: XKB_K_braille_dot_4                _EVDEVK(0x1f4)     v2.6.17K_BRL_DOT4 */
	/* Use: XKB_K_braille_dot_5                _EVDEVK(0x1f5)     v2.6.17K_BRL_DOT5 */
	/* Use: XKB_K_braille_dot_6                _EVDEVK(0x1f6)     v2.6.17K_BRL_DOT6 */
	/* Use: XKB_K_braille_dot_7                _EVDEVK(0x1f7)     v2.6.17K_BRL_DOT7 */
	/* Use: XKB_K_braille_dot_8                _EVDEVK(0x1f8)     v2.6.17K_BRL_DOT8 */
	/* Use: XKB_K_braille_dot_9                _EVDEVK(0x1f9)     v2.6.23K_BRL_DOT9 */
	/* Use: XKB_K_braille_dot_1                _EVDEVK(0x1fa)     v2.6.23K_BRL_DOT10 */
	K_XF86Numeric0     Key = 0x10081200 /* v2.6.28K_NUMERIC_0 */
	K_XF86Numeric1     Key = 0x10081201 /* v2.6.28K_NUMERIC_1 */
	K_XF86Numeric2     Key = 0x10081202 /* v2.6.28K_NUMERIC_2 */
	K_XF86Numeric3     Key = 0x10081203 /* v2.6.28K_NUMERIC_3 */
	K_XF86Numeric4     Key = 0x10081204 /* v2.6.28K_NUMERIC_4 */
	K_XF86Numeric5     Key = 0x10081205 /* v2.6.28K_NUMERIC_5 */
	K_XF86Numeric6     Key = 0x10081206 /* v2.6.28K_NUMERIC_6 */
	K_XF86Numeric7     Key = 0x10081207 /* v2.6.28K_NUMERIC_7 */
	K_XF86Numeric8     Key = 0x10081208 /* v2.6.28K_NUMERIC_8 */
	K_XF86Numeric9     Key = 0x10081209 /* v2.6.28K_NUMERIC_9 */
	K_XF86NumericStar  Key = 0x1008120a /* v2.6.28K_NUMERIC_STAR */
	K_XF86NumericPound Key = 0x1008120b /* v2.6.28K_NUMERIC_POUND */
	K_XF86NumericA     Key = 0x1008120c /* v4.1K_NUMERIC_A */
	K_XF86NumericB     Key = 0x1008120d /* v4.1K_NUMERIC_B */
	K_XF86NumericC     Key = 0x1008120e /* v4.1K_NUMERIC_C */
	K_XF86NumericD     Key = 0x1008120f /* v4.1K_NUMERIC_D */
	K_XF86CameraFocus  Key = 0x10081210 /* v2.6.33K_CAMERA_FOCUS */
	K_XF86WPSButton    Key = 0x10081211 /* v2.6.34K_WPS_BUTTON */
	/* Use: XKB_K_XF86TouchpadToggle           _EVDEVK(0x212)     v2.6.37K_TOUCHPAD_TOGGLE */
	/* Use: XKB_K_XF86TouchpadOn               _EVDEVK(0x213)     v2.6.37K_TOUCHPAD_ON */
	/* Use: XKB_K_XF86TouchpadOff              _EVDEVK(0x214)     v2.6.37K_TOUCHPAD_OFF */
	K_XF86CameraZoomIn    Key = 0x10081215 /* v2.6.39K_CAMERA_ZOOMIN */
	K_XF86CameraZoomOut   Key = 0x10081216 /* v2.6.39K_CAMERA_ZOOMOUT */
	K_XF86CameraUp        Key = 0x10081217 /* v2.6.39K_CAMERA_UP */
	K_XF86CameraDown      Key = 0x10081218 /* v2.6.39K_CAMERA_DOWN */
	K_XF86CameraLeft      Key = 0x10081219 /* v2.6.39K_CAMERA_LEFT */
	K_XF86CameraRight     Key = 0x1008121a /* v2.6.39K_CAMERA_RIGHT */
	K_XF86AttendantOn     Key = 0x1008121b /* v3.10K_ATTENDANT_ON */
	K_XF86AttendantOff    Key = 0x1008121c /* v3.10K_ATTENDANT_OFF */
	K_XF86AttendantToggle Key = 0x1008121d /* v3.10K_ATTENDANT_TOGGLE */
	K_XF86LightsToggle    Key = 0x1008121e /* v3.10K_LIGHTS_TOGGLE */
	K_XF86ALSToggle       Key = 0x10081230 /* v3.13K_ALS_TOGGLE */
	/* Use: XKB_K_XF86RotationLockToggle       _EVDEVK(0x231)     v4.16K_ROTATE_LOCK_TOGGLE */
	K_XF86RefreshRateToggle Key = 0x10081232 /* v6.9K_REFRESH_RATE_TOGGLE */
	K_XF86Buttonconfig      Key = 0x10081240 /* v3.16K_BUTTONCONFIG */
	K_XF86Taskmanager       Key = 0x10081241 /* v3.16K_TASKMANAGER */
	K_XF86Journal           Key = 0x10081242 /* v3.16K_JOURNAL */
	K_XF86ControlPanel      Key = 0x10081243 /* v3.16K_CONTROLPANEL */
	K_XF86AppSelect         Key = 0x10081244 /* v3.16K_APPSELECT */
	K_XF86Screensaver       Key = 0x10081245 /* v3.16K_SCREENSAVER */
	K_XF86VoiceCommand      Key = 0x10081246 /* v3.16K_VOICECOMMAND */
	K_XF86Assistant         Key = 0x10081247 /* v4.13K_ASSISTANT */
	/* Use: XKB_K_ISO_Next_Group               _EVDEVK(0x248)     v5.2K_KBD_LAYOUT_NEXT */
	K_XF86EmojiPicker             Key = 0x10081249 /* v5.13K_EMOJI_PICKER */
	K_XF86Dictate                 Key = 0x1008124a /* v5.17K_DICTATE */
	K_XF86CameraAccessEnable      Key = 0x1008124b /* v6.2K_CAMERA_ACCESS_ENABLE */
	K_XF86CameraAccessDisable     Key = 0x1008124c /* v6.2K_CAMERA_ACCESS_DISABLE */
	K_XF86CameraAccessToggle      Key = 0x1008124d /* v6.2K_CAMERA_ACCESS_TOGGLE */
	K_XF86Accessibility           Key = 0x1008124e /* v6.10K_ACCESSIBILITY */
	K_XF86DoNotDisturb            Key = 0x1008124f /* v6.10K_DO_NOT_DISTURB */
	K_XF86BrightnessMin           Key = 0x10081250 /* v3.16K_BRIGHTNESS_MIN */
	K_XF86BrightnessMax           Key = 0x10081251 /* v3.16K_BRIGHTNESS_MAX */
	K_XF86KbdInputAssistPrev      Key = 0x10081260 /* v3.18K_KBDINPUTASSIST_PREV */
	K_XF86KbdInputAssistNext      Key = 0x10081261 /* v3.18K_KBDINPUTASSIST_NEXT */
	K_XF86KbdInputAssistPrevgroup Key = 0x10081262 /* v3.18K_KBDINPUTASSIST_PREVGROUP */
	K_XF86KbdInputAssistNextgroup Key = 0x10081263 /* v3.18K_KBDINPUTASSIST_NEXTGROUP */
	K_XF86KbdInputAssistAccept    Key = 0x10081264 /* v3.18K_KBDINPUTASSIST_ACCEPT */
	K_XF86KbdInputAssistCancel    Key = 0x10081265 /* v3.18K_KBDINPUTASSIST_CANCEL */
	K_XF86RightUp                 Key = 0x10081266 /* v4.7K_RIGHT_UP */
	K_XF86RightDown               Key = 0x10081267 /* v4.7K_RIGHT_DOWN */
	K_XF86LeftUp                  Key = 0x10081268 /* v4.7K_LEFT_UP */
	K_XF86LeftDown                Key = 0x10081269 /* v4.7K_LEFT_DOWN */
	K_XF86RootMenu                Key = 0x1008126a /* v4.7K_ROOT_MENU */
	K_XF86MediaTopMenu            Key = 0x1008126b /* v4.7K_MEDIA_TOP_MENU */
	K_XF86Numeric11               Key = 0x1008126c /* v4.7K_NUMERIC_11 */
	K_XF86Numeric12               Key = 0x1008126d /* v4.7K_NUMERIC_12 */
	K_XF86AudioDesc               Key = 0x1008126e /* v4.7K_AUDIO_DESC */
	K_XF863DMode                  Key = 0x1008126f /* v4.7K_3D_MODE */
	K_XF86NextFavorite            Key = 0x10081270 /* v4.7K_NEXT_FAVORITE */
	K_XF86StopRecord              Key = 0x10081271 /* v4.7K_STOP_RECORD */
	K_XF86PauseRecord             Key = 0x10081272 /* v4.7K_PAUSE_RECORD */
	K_XF86VOD                     Key = 0x10081273 /* v4.7K_VOD */
	K_XF86Unmute                  Key = 0x10081274 /* v4.7K_UNMUTE */
	K_XF86FastReverse             Key = 0x10081275 /* v4.7K_FASTREVERSE */
	K_XF86SlowReverse             Key = 0x10081276 /* v4.7K_SLOWREVERSE */
	K_XF86Data                    Key = 0x10081277 /* v4.7K_DATA */
	K_XF86OnScreenKeyboard        Key = 0x10081278 /* v4.12K_ONSCREEN_KEYBOARD */
	K_XF86PrivacyScreenToggle     Key = 0x10081279 /* v5.5K_PRIVACY_SCREEN_TOGGLE */
	K_XF86SelectiveScreenshot     Key = 0x1008127a /* v5.6K_SELECTIVE_SCREENSHOT */
	K_XF86NextElement             Key = 0x1008127b /* v5.18K_NEXT_ELEMENT */
	K_XF86PreviousElement         Key = 0x1008127c /* v5.18K_PREVIOUS_ELEMENT */
	K_XF86AutopilotEngageToggle   Key = 0x1008127d /* v5.18K_AUTOPILOT_ENGAGE_TOGGLE */
	K_XF86MarkWaypoint            Key = 0x1008127e /* v5.18K_MARK_WAYPOINT */
	K_XF86Sos                     Key = 0x1008127f /* v5.18K_SOS */
	K_XF86NavChart                Key = 0x10081280 /* v5.18K_NAV_CHART */
	K_XF86FishingChart            Key = 0x10081281 /* v5.18K_FISHING_CHART */
	K_XF86SingleRangeRadar        Key = 0x10081282 /* v5.18K_SINGLE_RANGE_RADAR */
	K_XF86DualRangeRadar          Key = 0x10081283 /* v5.18K_DUAL_RANGE_RADAR */
	K_XF86RadarOverlay            Key = 0x10081284 /* v5.18K_RADAR_OVERLAY */
	K_XF86TraditionalSonar        Key = 0x10081285 /* v5.18K_TRADITIONAL_SONAR */
	K_XF86ClearvuSonar            Key = 0x10081286 /* v5.18K_CLEARVU_SONAR */
	K_XF86SidevuSonar             Key = 0x10081287 /* v5.18K_SIDEVU_SONAR */
	K_XF86NavInfo                 Key = 0x10081288 /* v5.18K_NAV_INFO */
	/* Use: XKB_K_XF86BrightnessAdjust         _EVDEVK(0x289)     v5.18K_BRIGHTNESS_MENU */
	K_XF86Macro1           Key = 0x10081290 /* v5.5K_MACRO1 */
	K_XF86Macro2           Key = 0x10081291 /* v5.5K_MACRO2 */
	K_XF86Macro3           Key = 0x10081292 /* v5.5K_MACRO3 */
	K_XF86Macro4           Key = 0x10081293 /* v5.5K_MACRO4 */
	K_XF86Macro5           Key = 0x10081294 /* v5.5K_MACRO5 */
	K_XF86Macro6           Key = 0x10081295 /* v5.5K_MACRO6 */
	K_XF86Macro7           Key = 0x10081296 /* v5.5K_MACRO7 */
	K_XF86Macro8           Key = 0x10081297 /* v5.5K_MACRO8 */
	K_XF86Macro9           Key = 0x10081298 /* v5.5K_MACRO9 */
	K_XF86Macro10          Key = 0x10081299 /* v5.5K_MACRO10 */
	K_XF86Macro11          Key = 0x1008129a /* v5.5K_MACRO11 */
	K_XF86Macro12          Key = 0x1008129b /* v5.5K_MACRO12 */
	K_XF86Macro13          Key = 0x1008129c /* v5.5K_MACRO13 */
	K_XF86Macro14          Key = 0x1008129d /* v5.5K_MACRO14 */
	K_XF86Macro15          Key = 0x1008129e /* v5.5K_MACRO15 */
	K_XF86Macro16          Key = 0x1008129f /* v5.5K_MACRO16 */
	K_XF86Macro17          Key = 0x100812a0 /* v5.5K_MACRO17 */
	K_XF86Macro18          Key = 0x100812a1 /* v5.5K_MACRO18 */
	K_XF86Macro19          Key = 0x100812a2 /* v5.5K_MACRO19 */
	K_XF86Macro20          Key = 0x100812a3 /* v5.5K_MACRO20 */
	K_XF86Macro21          Key = 0x100812a4 /* v5.5K_MACRO21 */
	K_XF86Macro22          Key = 0x100812a5 /* v5.5K_MACRO22 */
	K_XF86Macro23          Key = 0x100812a6 /* v5.5K_MACRO23 */
	K_XF86Macro24          Key = 0x100812a7 /* v5.5K_MACRO24 */
	K_XF86Macro25          Key = 0x100812a8 /* v5.5K_MACRO25 */
	K_XF86Macro26          Key = 0x100812a9 /* v5.5K_MACRO26 */
	K_XF86Macro27          Key = 0x100812aa /* v5.5K_MACRO27 */
	K_XF86Macro28          Key = 0x100812ab /* v5.5K_MACRO28 */
	K_XF86Macro29          Key = 0x100812ac /* v5.5K_MACRO29 */
	K_XF86Macro30          Key = 0x100812ad /* v5.5K_MACRO30 */
	K_XF86MacroRecordStart Key = 0x100812b0 /* v5.5K_MACRO_RECORD_START */
	K_XF86MacroRecordStop  Key = 0x100812b1 /* v5.5K_MACRO_RECORD_STOP */
	K_XF86MacroPresetCycle Key = 0x100812b2 /* v5.5K_MACRO_PRESET_CYCLE */
	K_XF86MacroPreset1     Key = 0x100812b3 /* v5.5K_MACRO_PRESET1 */
	K_XF86MacroPreset2     Key = 0x100812b4 /* v5.5K_MACRO_PRESET2 */
	K_XF86MacroPreset3     Key = 0x100812b5 /* v5.5K_MACRO_PRESET3 */
	K_XF86KbdLcdMenu1      Key = 0x100812b8 /* v5.5K_KBD_LCD_MENU1 */
	K_XF86KbdLcdMenu2      Key = 0x100812b9 /* v5.5K_KBD_LCD_MENU2 */
	K_XF86KbdLcdMenu3      Key = 0x100812ba /* v5.5K_KBD_LCD_MENU3 */
	K_XF86KbdLcdMenu4      Key = 0x100812bb /* v5.5K_KBD_LCD_MENU4 */
	K_XF86KbdLcdMenu5      Key = 0x100812bc /* v5.5K_KBD_LCD_MENU5 */
	/*
	 * Copyright (c) 1991, Oracle and/or its affiliates.
	 *
	 * Permission is hereby granted, free of charge, to any person obtaining a
	 * copy of this software and associated documentation files (the "Software"),
	 * to deal in the Software without restriction, including without limitation
	 * the rights to use, copy, modify, merge, publish, distribute, sublicense,
	 * and/or sell copies of the Software, and to permit persons to whom the
	 * Software is furnished to do so, subject to the following conditions:
	 *
	 * The above copyright notice and this permission notice (including the next
	 * paragraph) shall be included in all copies or substantial portions of the
	 * Software.
	 *
	 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
	 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
	 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL
	 * THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
	 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
	 * FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER
	 * DEALINGS IN THE SOFTWARE.
	 */
	/************************************************************

	  Copyright 1991, 1998  The Open Group

	  Permission to use, copy, modify, distribute, and sell this software and its
	  documentation for any purpose is hereby granted without fee, provided that
	  the above copyright notice appear in all copies and that both that
	  copyright notice and this permission notice appear in supporting
	  documentation.

	  The above copyright notice and this permission notice shall be included in
	  all copies or substantial portions of the Software.

	  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
	  IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
	  FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL THE
	  OPEN GROUP BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN
	  AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
	  CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

	  Except as contained in this notice, the name of The Open Group shall not be
	  used in advertising or otherwise to promote the sale, use or other dealings
	  in this Software without prior written authorization from The Open Group.

	  ***********************************************************/

	/*
	 * Floating Accent
	 */

	K_SunFA_Grave     Key = 0x1005ff00
	K_SunFA_Circum    Key = 0x1005ff01
	K_SunFA_Tilde     Key = 0x1005ff02
	K_SunFA_Acute     Key = 0x1005ff03
	K_SunFA_Diaeresis Key = 0x1005ff04
	K_SunFA_Cedilla   Key = 0x1005ff05

	/*
	 * Miscellaneous Functions
	 */

	K_SunF36 Key = 0x1005ff10 /* Labeled F11 */
	K_SunF37 Key = 0x1005ff11 /* Labeled F12 */

	K_SunSys_Req      Key = 0x1005ff60
	K_SunPrint_Screen Key = 0x0000ff61 /* Same as XKB_K_Print */

	/*
	 * International & Multi-Key Character Composition
	 */

	K_SunCompose  Key = 0x0000ff20 /* Same as XKB_K_Multi_key */
	K_SunAltGraph Key = 0x0000ff7e /* Same as XKB_K_Mode_switch */

	/*
	 * Cursor Control
	 */

	K_SunPageUp   Key = 0x0000ff55 /* Same as XKB_K_Prior */
	K_SunPageDown Key = 0x0000ff56 /* Same as XKB_K_Next */

	/*
	 * Open Look Functions
	 */

	K_SunUndo  Key = 0x0000ff65 /* Same as XKB_K_Undo */
	K_SunAgain Key = 0x0000ff66 /* Same as XKB_K_Redo */
	K_SunFind  Key = 0x0000ff68 /* Same as XKB_K_Find */
	K_SunStop  Key = 0x0000ff69 /* Same as XKB_K_Cancel */
	K_SunProps Key = 0x1005ff70
	K_SunFront Key = 0x1005ff71
	K_SunCopy  Key = 0x1005ff72
	K_SunOpen  Key = 0x1005ff73
	K_SunPaste Key = 0x1005ff74
	K_SunCut   Key = 0x1005ff75

	K_SunPowerSwitch          Key = 0x1005ff76
	K_SunAudioLowerVolume     Key = 0x1005ff77
	K_SunAudioMute            Key = 0x1005ff78
	K_SunAudioRaiseVolume     Key = 0x1005ff79
	K_SunVideoDegauss         Key = 0x1005ff7a
	K_SunVideoLowerBrightness Key = 0x1005ff7b
	K_SunVideoRaiseBrightness Key = 0x1005ff7c
	K_SunPowerSwitchShift     Key = 0x1005ff7d
	/***********************************************************

	  Copyright 1988, 1998  The Open Group

	  Permission to use, copy, modify, distribute, and sell this software and its
	  documentation for any purpose is hereby granted without fee, provided that
	  the above copyright notice appear in all copies and that both that
	  copyright notice and this permission notice appear in supporting
	  documentation.

	  The above copyright notice and this permission notice shall be included in
	  all copies or substantial portions of the Software.

	  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
	  IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
	  FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL THE
	  OPEN GROUP BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN
	  AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
	  CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

	  Except as contained in this notice, the name of The Open Group shall not be
	  used in advertising or otherwise to promote the sale, use or other dealings
	  in this Software without prior written authorization from The Open Group.


	  Copyright 1988 by Digital Equipment Corporation, Maynard, Massachusetts.

	                          All Rights Reserved

	  Permission to use, copy, modify, and distribute this software and its
	  documentation for any purpose and without fee is hereby granted,
	  provided that the above copyright notice appear in all copies and that
	  both that copyright notice and this permission notice appear in
	  supporting documentation, and that the name of Digital not be
	  used in advertising or publicity pertaining to distribution of the
	  software without specific, written prior permission.

	  DIGITAL DISCLAIMS ALL WARRANTIES WITH REGARD TO THIS SOFTWARE, INCLUDING
	  ALL IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS, IN NO EVENT SHALL
	  DIGITAL BE LIABLE FOR ANY SPECIAL, INDIRECT OR CONSEQUENTIAL DAMAGES OR
	  ANY DAMAGES WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS,
	  WHETHER IN AN ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION,
	  ARISING OUT OF OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS
	  SOFTWARE.

	  ******************************************************************/

	/*
	 * DEC private keysyms
	 * (29th bit set)
	 */

	/* two-key compose sequence initiators, chosen to map to Latin1 characters */

	K_Dring_accent       Key = 0x1000feb0
	K_Dcircumflex_accent Key = 0x1000fe5e
	K_Dcedilla_accent    Key = 0x1000fe2c
	K_Dacute_accent      Key = 0x1000fe27
	K_Dgrave_accent      Key = 0x1000fe60
	K_Dtilde             Key = 0x1000fe7e
	K_Ddiaeresis         Key = 0x1000fe22

	/* special keysym for LK2** "Remove" key on editing keypad */

	K_DRemove Key = 0x1000ff00 /* Remove */
	/*

	   Copyright 1987, 1998  The Open Group

	   Permission to use, copy, modify, distribute, and sell this software and its
	   documentation for any purpose is hereby granted without fee, provided that
	   the above copyright notice appear in all copies and that both that
	   copyright notice and this permission notice appear in supporting
	   documentation.

	   The above copyright notice and this permission notice shall be included
	   in all copies or substantial portions of the Software.

	   THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS
	   OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
	   MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
	   IN NO EVENT SHALL THE OPEN GROUP BE LIABLE FOR ANY CLAIM, DAMAGES OR
	   OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE,
	   ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
	   OTHER DEALINGS IN THE SOFTWARE.

	   Except as contained in this notice, the name of The Open Group shall
	   not be used in advertising or otherwise to promote the sale, use or
	   other dealings in this Software without prior written authorization
	   from The Open Group.

	   Copyright 1987 by Digital Equipment Corporation, Maynard, Massachusetts,

	                           All Rights Reserved

	   Permission to use, copy, modify, and distribute this software and its
	   documentation for any purpose and without fee is hereby granted,
	   provided that the above copyright notice appear in all copies and that
	   both that copyright notice and this permission notice appear in
	   supporting documentation, and that the names of Hewlett Packard
	   or Digital not be
	   used in advertising or publicity pertaining to distribution of the
	   software without specific, written prior permission.

	   DIGITAL DISCLAIMS ALL WARRANTIES WITH REGARD TO THIS SOFTWARE, INCLUDING
	   ALL IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS, IN NO EVENT SHALL
	   DIGITAL BE LIABLE FOR ANY SPECIAL, INDIRECT OR CONSEQUENTIAL DAMAGES OR
	   ANY DAMAGES WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS,
	   WHETHER IN AN ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION,
	   ARISING OUT OF OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS
	   SOFTWARE.

	   HEWLETT-PACKARD MAKES NO WARRANTY OF ANY KIND WITH REGARD
	   TO THIS SOFTWARE, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
	   WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR
	   PURPOSE.  Hewlett-Packard shall not be liable for errors
	   contained herein or direct, indirect, special, incidental or
	   consequential damages in connection with the furnishing,
	   performance, or use of this material.

	*/

	K_hpClearLine        Key = 0x1000ff6f
	K_hpInsertLine       Key = 0x1000ff70
	K_hpDeleteLine       Key = 0x1000ff71
	K_hpInsertChar       Key = 0x1000ff72
	K_hpDeleteChar       Key = 0x1000ff73
	K_hpBackTab          Key = 0x1000ff74
	K_hpKP_BackTab       Key = 0x1000ff75
	K_hpModelock1        Key = 0x1000ff48
	K_hpModelock2        Key = 0x1000ff49
	K_hpReset            Key = 0x1000ff6c
	K_hpSystem           Key = 0x1000ff6d
	K_hpUser             Key = 0x1000ff6e
	K_hpmute_acute       Key = 0x100000a8
	K_hpmute_grave       Key = 0x100000a9
	K_hpmute_asciicircum Key = 0x100000aa
	K_hpmute_diaeresis   Key = 0x100000ab
	K_hpmute_asciitilde  Key = 0x100000ac
	K_hplira             Key = 0x100000af
	K_hpguilder          Key = 0x100000be
	K_hpYdiaeresis       Key = 0x100000ee
	K_hpIO               Key = 0x100000ee /* deprecated alias for hpYdiaeresis */
	K_hplongminus        Key = 0x100000f6
	K_hpblock            Key = 0x100000fc

	K_osfCopy         Key = 0x1004ff02
	K_osfCut          Key = 0x1004ff03
	K_osfPaste        Key = 0x1004ff04
	K_osfBackTab      Key = 0x1004ff07
	K_osfBackSpace    Key = 0x1004ff08
	K_osfClear        Key = 0x1004ff0b
	K_osfEscape       Key = 0x1004ff1b
	K_osfAddMode      Key = 0x1004ff31
	K_osfPrimaryPaste Key = 0x1004ff32
	K_osfQuickPaste   Key = 0x1004ff33
	K_osfPageLeft     Key = 0x1004ff40
	K_osfPageUp       Key = 0x1004ff41
	K_osfPageDown     Key = 0x1004ff42
	K_osfPageRight    Key = 0x1004ff43
	K_osfActivate     Key = 0x1004ff44
	K_osfMenuBar      Key = 0x1004ff45
	K_osfLeft         Key = 0x1004ff51
	K_osfUp           Key = 0x1004ff52
	K_osfRight        Key = 0x1004ff53
	K_osfDown         Key = 0x1004ff54
	K_osfEndLine      Key = 0x1004ff57
	K_osfBeginLine    Key = 0x1004ff58
	K_osfEndData      Key = 0x1004ff59
	K_osfBeginData    Key = 0x1004ff5a
	K_osfPrevMenu     Key = 0x1004ff5b
	K_osfNextMenu     Key = 0x1004ff5c
	K_osfPrevField    Key = 0x1004ff5d
	K_osfNextField    Key = 0x1004ff5e
	K_osfSelect       Key = 0x1004ff60
	K_osfInsert       Key = 0x1004ff63
	K_osfUndo         Key = 0x1004ff65
	K_osfMenu         Key = 0x1004ff67
	K_osfCancel       Key = 0x1004ff69
	K_osfHelp         Key = 0x1004ff6a
	K_osfSelectAll    Key = 0x1004ff71
	K_osfDeselectAll  Key = 0x1004ff72
	K_osfReselect     Key = 0x1004ff73
	K_osfExtend       Key = 0x1004ff74
	K_osfRestore      Key = 0x1004ff78
	K_osfDelete       Key = 0x1004ffff
)
