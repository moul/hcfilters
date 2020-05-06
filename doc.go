// Copyright © 2020 Manfred Touron <manfred.life>
// SPDX-License-Identifier: Apache-2.0 OR MIT

// Package hcfilters provides filtering middlewares for github.com/gregjones/httpcache.
//
// message from the author:
//	+--------------------------------------------------------------+
//	| * * * ░░░░░░░░░░░░░░░░░░░░  Hello  ░░░░░░░░░░░░░░░░░░░░░░░░░░|
//	+--------------------------------------------------------------+
//	|                                                              |
//	|     ++              ______________________________________   |
//	|     ++++           /                                      \  |
//	|      ++++          |                                      |  |
//	|    ++++++++++      |   Feel free to contribute to this    |  |
//	|   +++       |      |       project or contact me on       |  |
//	|   ++         |     |    manfred.life if you like this     |  |
//	|   +  -==   ==|     |               project!               |  |
//	|  (   <*>   <*>     |                                      |  |
//	|   |          |    /|                  :)                  |  |
//	|   |         _)   / |                                      |  |
//	|   |      +++    /  \______________________________________/  |
//	|    \      =+   /                                             |
//	|     \      +                                                 |
//	|     |\++++++                                                 |
//	|     |  ++++      ||//                                        |
//	|  ___|   |___    _||/__                                     __|
//	| /    ---    \   \|  |||                   __ _  ___  __ __/ /|
//	|/  |       |  \    \ /                    /  ' \/ _ \/ // / / |
//	||  |       |  |    | |                   /_/_/_/\___/\_,_/_/  |
//	+--------------------------------------------------------------+
//
// This is how to wrap an existing backend:
//
// before:
//
//     client := &http.Client{
//         Transport: httpcache.NewTransport(
//	       diskcache.New(diskcachePath),
//         ),
//     }
//
// after:
//
//     client := &http.Client{
//         Transport: httpcache.NewTransport(
//             hcfilters.MaxSize( // skip caching results > 2Mb
//                 diskcache.New(diskcachePath),
//                 2*1024*1024,
//             ),
//         ),
//     }
package hcfilters
