// Code generated by glow (https://github.com/go-gl/glow). DO NOT EDIT.

// Copyright (c) 2010 Khronos Group.
// This material may be distributed subject to the terms and conditions
// set forth in the Open Publication License, v 1.0, 8 June 1999.
// http://opencontent.org/openpub/.
//
// Copyright (c) 1991-2006 Silicon Graphics, Inc.
// This document is licensed under the SGI Free Software B License.
// For details, see http://oss.sgi.com/projects/FreeB.

// Package wgl implements Go bindings to OpenGL.
//
// This package was automatically generated using Glow:
//  https://github.com/go-gl/glow
//
package wgl

const (
	ERROR_INCOMPATIBLE_AFFINITY_MASKS_NV       = 0x20D0
	ERROR_INCOMPATIBLE_DEVICE_CONTEXTS_ARB     = 0x2054
	ERROR_INVALID_PIXEL_TYPE_ARB               = 0x2043
	ERROR_INVALID_PIXEL_TYPE_EXT               = 0x2043
	ERROR_INVALID_PROFILE_ARB                  = 0x2096
	ERROR_INVALID_VERSION_ARB                  = 0x2095
	ERROR_MISSING_AFFINITY_MASK_NV             = 0x20D1
	ACCELERATION_ARB                           = 0x2003
	ACCELERATION_EXT                           = 0x2003
	ACCESS_READ_ONLY_NV                        = 0x00000000
	ACCESS_READ_WRITE_NV                       = 0x00000001
	ACCESS_WRITE_DISCARD_NV                    = 0x00000002
	ACCUM_ALPHA_BITS_ARB                       = 0x2021
	ACCUM_ALPHA_BITS_EXT                       = 0x2021
	ACCUM_BITS_ARB                             = 0x201D
	ACCUM_BITS_EXT                             = 0x201D
	ACCUM_BLUE_BITS_ARB                        = 0x2020
	ACCUM_BLUE_BITS_EXT                        = 0x2020
	ACCUM_GREEN_BITS_ARB                       = 0x201F
	ACCUM_GREEN_BITS_EXT                       = 0x201F
	ACCUM_RED_BITS_ARB                         = 0x201E
	ACCUM_RED_BITS_EXT                         = 0x201E
	ALPHA_BITS_ARB                             = 0x201B
	ALPHA_BITS_EXT                             = 0x201B
	ALPHA_SHIFT_ARB                            = 0x201C
	ALPHA_SHIFT_EXT                            = 0x201C
	AUX0_ARB                                   = 0x2087
	AUX1_ARB                                   = 0x2088
	AUX2_ARB                                   = 0x2089
	AUX3_ARB                                   = 0x208A
	AUX4_ARB                                   = 0x208B
	AUX5_ARB                                   = 0x208C
	AUX6_ARB                                   = 0x208D
	AUX7_ARB                                   = 0x208E
	AUX8_ARB                                   = 0x208F
	AUX9_ARB                                   = 0x2090
	AUX_BUFFERS_ARB                            = 0x2024
	AUX_BUFFERS_EXT                            = 0x2024
	BACK_COLOR_BUFFER_BIT_ARB                  = 0x00000002
	BACK_LEFT_ARB                              = 0x2085
	BACK_RIGHT_ARB                             = 0x2086
	BIND_TO_TEXTURE_DEPTH_NV                   = 0x20A3
	BIND_TO_TEXTURE_RECTANGLE_DEPTH_NV         = 0x20A4
	BIND_TO_TEXTURE_RECTANGLE_FLOAT_RGBA_NV    = 0x20B4
	BIND_TO_TEXTURE_RECTANGLE_FLOAT_RGB_NV     = 0x20B3
	BIND_TO_TEXTURE_RECTANGLE_FLOAT_RG_NV      = 0x20B2
	BIND_TO_TEXTURE_RECTANGLE_FLOAT_R_NV       = 0x20B1
	BIND_TO_TEXTURE_RECTANGLE_RGBA_NV          = 0x20A1
	BIND_TO_TEXTURE_RECTANGLE_RGB_NV           = 0x20A0
	BIND_TO_TEXTURE_RGBA_ARB                   = 0x2071
	BIND_TO_TEXTURE_RGB_ARB                    = 0x2070
	BIND_TO_VIDEO_RGBA_NV                      = 0x20C1
	BIND_TO_VIDEO_RGB_AND_DEPTH_NV             = 0x20C2
	BIND_TO_VIDEO_RGB_NV                       = 0x20C0
	BLUE_BITS_ARB                              = 0x2019
	BLUE_BITS_EXT                              = 0x2019
	BLUE_SHIFT_ARB                             = 0x201A
	BLUE_SHIFT_EXT                             = 0x201A
	COLORSPACE_EXT                             = 0x3087
	COLORSPACE_LINEAR_EXT                      = 0x308A
	COLORSPACE_SRGB_EXT                        = 0x3089
	COLOR_BITS_ARB                             = 0x2014
	COLOR_BITS_EXT                             = 0x2014
	COLOR_SAMPLES_NV                           = 0x20B9
	CONTEXT_COMPATIBILITY_PROFILE_BIT_ARB      = 0x00000002
	CONTEXT_CORE_PROFILE_BIT_ARB               = 0x00000001
	CONTEXT_DEBUG_BIT_ARB                      = 0x00000001
	CONTEXT_ES2_PROFILE_BIT_EXT                = 0x00000004
	CONTEXT_ES_PROFILE_BIT_EXT                 = 0x00000004
	CONTEXT_FLAGS_ARB                          = 0x2094
	CONTEXT_FORWARD_COMPATIBLE_BIT_ARB         = 0x00000002
	CONTEXT_LAYER_PLANE_ARB                    = 0x2093
	CONTEXT_MAJOR_VERSION_ARB                  = 0x2091
	CONTEXT_MINOR_VERSION_ARB                  = 0x2092
	CONTEXT_OPENGL_NO_ERROR_ARB                = 0x31B3
	CONTEXT_PROFILE_MASK_ARB                   = 0x9126
	CONTEXT_RELEASE_BEHAVIOR_ARB               = 0x2097
	CONTEXT_RELEASE_BEHAVIOR_FLUSH_ARB         = 0x2098
	CONTEXT_RELEASE_BEHAVIOR_NONE_ARB          = 0
	CONTEXT_RESET_ISOLATION_BIT_ARB            = 0x00000008
	CONTEXT_RESET_NOTIFICATION_STRATEGY_ARB    = 0x8256
	CONTEXT_ROBUST_ACCESS_BIT_ARB              = 0x00000004
	COVERAGE_SAMPLES_NV                        = 0x2042
	CUBE_MAP_FACE_ARB                          = 0x207C
	DEPTH_BITS_ARB                             = 0x2022
	DEPTH_BITS_EXT                             = 0x2022
	DEPTH_BUFFER_BIT_ARB                       = 0x00000004
	DEPTH_COMPONENT_NV                         = 0x20A7
	DEPTH_FLOAT_EXT                            = 0x2040
	DEPTH_TEXTURE_FORMAT_NV                    = 0x20A5
	DIGITAL_VIDEO_CURSOR_ALPHA_FRAMEBUFFER_I3D = 0x2050
	DIGITAL_VIDEO_CURSOR_ALPHA_VALUE_I3D       = 0x2051
	DIGITAL_VIDEO_CURSOR_INCLUDED_I3D          = 0x2052
	DIGITAL_VIDEO_GAMMA_CORRECTED_I3D          = 0x2053
	DOUBLE_BUFFER_ARB                          = 0x2011
	DOUBLE_BUFFER_EXT                          = 0x2011
	DRAW_TO_BITMAP_ARB                         = 0x2002
	DRAW_TO_BITMAP_EXT                         = 0x2002
	DRAW_TO_PBUFFER_ARB                        = 0x202D
	DRAW_TO_PBUFFER_EXT                        = 0x202D
	DRAW_TO_WINDOW_ARB                         = 0x2001
	DRAW_TO_WINDOW_EXT                         = 0x2001
	FLOAT_COMPONENTS_NV                        = 0x20B0
	FONT_LINES                                 = 0
	FONT_POLYGONS                              = 1
	FRAMEBUFFER_SRGB_CAPABLE_ARB               = 0x20A9
	FRAMEBUFFER_SRGB_CAPABLE_EXT               = 0x20A9
	FRONT_COLOR_BUFFER_BIT_ARB                 = 0x00000001
	FRONT_LEFT_ARB                             = 0x2083
	FRONT_RIGHT_ARB                            = 0x2084
	FULL_ACCELERATION_ARB                      = 0x2027
	FULL_ACCELERATION_EXT                      = 0x2027
	GAMMA_EXCLUDE_DESKTOP_I3D                  = 0x204F
	GAMMA_TABLE_SIZE_I3D                       = 0x204E
	GENERIC_ACCELERATION_ARB                   = 0x2026
	GENERIC_ACCELERATION_EXT                   = 0x2026
	GENLOCK_SOURCE_DIGITAL_FIELD_I3D           = 0x2049
	GENLOCK_SOURCE_DIGITAL_SYNC_I3D            = 0x2048
	GENLOCK_SOURCE_EDGE_BOTH_I3D               = 0x204C
	GENLOCK_SOURCE_EDGE_FALLING_I3D            = 0x204A
	GENLOCK_SOURCE_EDGE_RISING_I3D             = 0x204B
	GENLOCK_SOURCE_EXTERNAL_FIELD_I3D          = 0x2046
	GENLOCK_SOURCE_EXTERNAL_SYNC_I3D           = 0x2045
	GENLOCK_SOURCE_EXTERNAL_TTL_I3D            = 0x2047
	GENLOCK_SOURCE_MULTIVIEW_I3D               = 0x2044
	GPU_CLOCK_AMD                              = 0x21A4
	GPU_FASTEST_TARGET_GPUS_AMD                = 0x21A2
	GPU_NUM_PIPES_AMD                          = 0x21A5
	GPU_NUM_RB_AMD                             = 0x21A7
	GPU_NUM_SIMD_AMD                           = 0x21A6
	GPU_NUM_SPI_AMD                            = 0x21A8
	GPU_OPENGL_VERSION_STRING_AMD              = 0x1F02
	GPU_RAM_AMD                                = 0x21A3
	GPU_RENDERER_STRING_AMD                    = 0x1F01
	GPU_VENDOR_AMD                             = 0x1F00
	GREEN_BITS_ARB                             = 0x2017
	GREEN_BITS_EXT                             = 0x2017
	GREEN_SHIFT_ARB                            = 0x2018
	GREEN_SHIFT_EXT                            = 0x2018
	IMAGE_BUFFER_LOCK_I3D                      = 0x00000002
	IMAGE_BUFFER_MIN_ACCESS_I3D                = 0x00000001
	LOSE_CONTEXT_ON_RESET_ARB                  = 0x8252
	MAX_PBUFFER_HEIGHT_ARB                     = 0x2030
	MAX_PBUFFER_HEIGHT_EXT                     = 0x2030
	MAX_PBUFFER_PIXELS_ARB                     = 0x202E
	MAX_PBUFFER_PIXELS_EXT                     = 0x202E
	MAX_PBUFFER_WIDTH_ARB                      = 0x202F
	MAX_PBUFFER_WIDTH_EXT                      = 0x202F
	MIPMAP_LEVEL_ARB                           = 0x207B
	MIPMAP_TEXTURE_ARB                         = 0x2074
	NEED_PALETTE_ARB                           = 0x2004
	NEED_PALETTE_EXT                           = 0x2004
	NEED_SYSTEM_PALETTE_ARB                    = 0x2005
	NEED_SYSTEM_PALETTE_EXT                    = 0x2005
	NO_ACCELERATION_ARB                        = 0x2025
	NO_ACCELERATION_EXT                        = 0x2025
	NO_RESET_NOTIFICATION_ARB                  = 0x8261
	NO_TEXTURE_ARB                             = 0x2077
	NUMBER_OVERLAYS_ARB                        = 0x2008
	NUMBER_OVERLAYS_EXT                        = 0x2008
	NUMBER_PIXEL_FORMATS_ARB                   = 0x2000
	NUMBER_PIXEL_FORMATS_EXT                   = 0x2000
	NUMBER_UNDERLAYS_ARB                       = 0x2009
	NUMBER_UNDERLAYS_EXT                       = 0x2009
	NUM_VIDEO_CAPTURE_SLOTS_NV                 = 0x20CF
	NUM_VIDEO_SLOTS_NV                         = 0x20F0
	OPTIMAL_PBUFFER_HEIGHT_EXT                 = 0x2032
	OPTIMAL_PBUFFER_WIDTH_EXT                  = 0x2031
	PBUFFER_HEIGHT_ARB                         = 0x2035
	PBUFFER_HEIGHT_EXT                         = 0x2035
	PBUFFER_LARGEST_ARB                        = 0x2033
	PBUFFER_LARGEST_EXT                        = 0x2033
	PBUFFER_LOST_ARB                           = 0x2036
	PBUFFER_WIDTH_ARB                          = 0x2034
	PBUFFER_WIDTH_EXT                          = 0x2034
	PIXEL_TYPE_ARB                             = 0x2013
	PIXEL_TYPE_EXT                             = 0x2013
	RED_BITS_ARB                               = 0x2015
	RED_BITS_EXT                               = 0x2015
	RED_SHIFT_ARB                              = 0x2016
	RED_SHIFT_EXT                              = 0x2016
	SAMPLES_3DFX                               = 0x2061
	SAMPLES_ARB                                = 0x2042
	SAMPLES_EXT                                = 0x2042
	SAMPLE_BUFFERS_3DFX                        = 0x2060
	SAMPLE_BUFFERS_ARB                         = 0x2041
	SAMPLE_BUFFERS_EXT                         = 0x2041
	SHARE_ACCUM_ARB                            = 0x200E
	SHARE_ACCUM_EXT                            = 0x200E
	SHARE_DEPTH_ARB                            = 0x200C
	SHARE_DEPTH_EXT                            = 0x200C
	SHARE_STENCIL_ARB                          = 0x200D
	SHARE_STENCIL_EXT                          = 0x200D
	STENCIL_BITS_ARB                           = 0x2023
	STENCIL_BITS_EXT                           = 0x2023
	STENCIL_BUFFER_BIT_ARB                     = 0x00000008
	STEREO_ARB                                 = 0x2012
	STEREO_EMITTER_DISABLE_3DL                 = 0x2056
	STEREO_EMITTER_ENABLE_3DL                  = 0x2055
	STEREO_EXT                                 = 0x2012
	STEREO_POLARITY_INVERT_3DL                 = 0x2058
	STEREO_POLARITY_NORMAL_3DL                 = 0x2057
	SUPPORT_GDI_ARB                            = 0x200F
	SUPPORT_GDI_EXT                            = 0x200F
	SUPPORT_OPENGL_ARB                         = 0x2010
	SUPPORT_OPENGL_EXT                         = 0x2010
	SWAP_COPY_ARB                              = 0x2029
	SWAP_COPY_EXT                              = 0x2029
	SWAP_EXCHANGE_ARB                          = 0x2028
	SWAP_EXCHANGE_EXT                          = 0x2028
	SWAP_LAYER_BUFFERS_ARB                     = 0x2006
	SWAP_LAYER_BUFFERS_EXT                     = 0x2006
	SWAP_MAIN_PLANE                            = 0x00000001
	SWAP_METHOD_ARB                            = 0x2007
	SWAP_METHOD_EXT                            = 0x2007
	SWAP_OVERLAY1                              = 0x00000002
	SWAP_OVERLAY10                             = 0x00000400
	SWAP_OVERLAY11                             = 0x00000800
	SWAP_OVERLAY12                             = 0x00001000
	SWAP_OVERLAY13                             = 0x00002000
	SWAP_OVERLAY14                             = 0x00004000
	SWAP_OVERLAY15                             = 0x00008000
	SWAP_OVERLAY2                              = 0x00000004
	SWAP_OVERLAY3                              = 0x00000008
	SWAP_OVERLAY4                              = 0x00000010
	SWAP_OVERLAY5                              = 0x00000020
	SWAP_OVERLAY6                              = 0x00000040
	SWAP_OVERLAY7                              = 0x00000080
	SWAP_OVERLAY8                              = 0x00000100
	SWAP_OVERLAY9                              = 0x00000200
	SWAP_UNDEFINED_ARB                         = 0x202A
	SWAP_UNDEFINED_EXT                         = 0x202A
	SWAP_UNDERLAY1                             = 0x00010000
	SWAP_UNDERLAY10                            = 0x02000000
	SWAP_UNDERLAY11                            = 0x04000000
	SWAP_UNDERLAY12                            = 0x08000000
	SWAP_UNDERLAY13                            = 0x10000000
	SWAP_UNDERLAY14                            = 0x20000000
	SWAP_UNDERLAY15                            = 0x40000000
	SWAP_UNDERLAY2                             = 0x00020000
	SWAP_UNDERLAY3                             = 0x00040000
	SWAP_UNDERLAY4                             = 0x00080000
	SWAP_UNDERLAY5                             = 0x00100000
	SWAP_UNDERLAY6                             = 0x00200000
	SWAP_UNDERLAY7                             = 0x00400000
	SWAP_UNDERLAY8                             = 0x00800000
	SWAP_UNDERLAY9                             = 0x01000000
	TEXTURE_1D_ARB                             = 0x2079
	TEXTURE_2D_ARB                             = 0x207A
	TEXTURE_CUBE_MAP_ARB                       = 0x2078
	TEXTURE_CUBE_MAP_NEGATIVE_X_ARB            = 0x207E
	TEXTURE_CUBE_MAP_NEGATIVE_Y_ARB            = 0x2080
	TEXTURE_CUBE_MAP_NEGATIVE_Z_ARB            = 0x2082
	TEXTURE_CUBE_MAP_POSITIVE_X_ARB            = 0x207D
	TEXTURE_CUBE_MAP_POSITIVE_Y_ARB            = 0x207F
	TEXTURE_CUBE_MAP_POSITIVE_Z_ARB            = 0x2081
	TEXTURE_DEPTH_COMPONENT_NV                 = 0x20A6
	TEXTURE_FLOAT_RGBA_NV                      = 0x20B8
	TEXTURE_FLOAT_RGB_NV                       = 0x20B7
	TEXTURE_FLOAT_RG_NV                        = 0x20B6
	TEXTURE_FLOAT_R_NV                         = 0x20B5
	TEXTURE_FORMAT_ARB                         = 0x2072
	TEXTURE_RECTANGLE_NV                       = 0x20A2
	TEXTURE_RGBA_ARB                           = 0x2076
	TEXTURE_RGB_ARB                            = 0x2075
	TEXTURE_TARGET_ARB                         = 0x2073
	TRANSPARENT_ALPHA_VALUE_ARB                = 0x203A
	TRANSPARENT_ARB                            = 0x200A
	TRANSPARENT_BLUE_VALUE_ARB                 = 0x2039
	TRANSPARENT_EXT                            = 0x200A
	TRANSPARENT_GREEN_VALUE_ARB                = 0x2038
	TRANSPARENT_INDEX_VALUE_ARB                = 0x203B
	TRANSPARENT_RED_VALUE_ARB                  = 0x2037
	TRANSPARENT_VALUE_EXT                      = 0x200B
	TYPE_COLORINDEX_ARB                        = 0x202C
	TYPE_COLORINDEX_EXT                        = 0x202C
	TYPE_RGBA_ARB                              = 0x202B
	TYPE_RGBA_EXT                              = 0x202B
	TYPE_RGBA_FLOAT_ARB                        = 0x21A0
	TYPE_RGBA_FLOAT_ATI                        = 0x21A0
	TYPE_RGBA_UNSIGNED_FLOAT_EXT               = 0x20A8
	UNIQUE_ID_NV                               = 0x20CE
	VIDEO_OUT_ALPHA_NV                         = 0x20C4
	VIDEO_OUT_COLOR_AND_ALPHA_NV               = 0x20C6
	VIDEO_OUT_COLOR_AND_DEPTH_NV               = 0x20C7
	VIDEO_OUT_COLOR_NV                         = 0x20C3
	VIDEO_OUT_DEPTH_NV                         = 0x20C5
	VIDEO_OUT_FIELD_1                          = 0x20C9
	VIDEO_OUT_FIELD_2                          = 0x20CA
	VIDEO_OUT_FRAME                            = 0x20C8
	VIDEO_OUT_STACKED_FIELDS_1_2               = 0x20CB
	VIDEO_OUT_STACKED_FIELDS_2_1               = 0x20CC
)

// Init initializes the OpenGL bindings by loading the function pointers (for
// each OpenGL function) from the active OpenGL context.
//
// It must be called under the presence of an active OpenGL context, e.g.,
// always after calling window.MakeContextCurrent() and always before calling
// any OpenGL functions exported by this package.
//
// On Windows, Init loads pointers that are context-specific (and hence you
// must re-init if switching between OpenGL contexts, although not calling Init
// again after switching between OpenGL contexts may work if the contexts belong
// to the same graphics driver/device).
//
// On macOS and the other POSIX systems, the behavior is different, but code
// written compatible with the Windows behavior is compatible with macOS and the
// other POSIX systems. That is, always Init under an active OpenGL context, and
// always re-init after switching graphics contexts.
//
// For information about caveats of Init, you should read the "Platform Specific
// Function Retrieval" section of https://www.opengl.org/wiki/Load_OpenGL_Functions.
func Init() error {
	return InitWithProcAddrFunc(getProcAddress)
}
