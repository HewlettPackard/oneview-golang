
The Go Playground
Imports
1
2
3
4
5
6
7
8
9
10
11
12
13
14
15
16
17
18
19
20
21
22
23
24
25
26
27
28
29
30
31
32
33
34
35
36
37
38
39
40
41
42
43
44
45
46
47
48
49
50
51
52
53
54
55
56
57
58
59
60
61
62
63
64
65
66
67
68
69
70
71
72
73
74
75
76
77
78
79
80
81
82
83
84
85
86
87
88
89
90
91
92
93
94
95
96
97
98
99
100
101
102
103
104
105
106
107
108
109
110
111
112
113
114
115
116
117
118
119
120
121
122
123
124
125
126
127
128
129
130
131
132
133
134
135
136
137
138
139
140
141
142
143
144
145
146
147
148
149
150
151
152
153
154
155
156
157
158
159
160
161
162
163
164
165
166
167
168
169
170
171
172
173
174
175
176
177
178
179
180
181
182
183
184
185
186
187
188
189
190
191
192
193
194
195
196
197
198
199
200
201
202
203
204
205
206
207
208
209
210
211
212
213
214
215
216
217
218
219
220
221
222
223
224
225
226
227
228
229
230
231
232
233
234
235
236
237
238
239
240
241
242
243
244
245
246
247
248
249
250
251
252
253
254
255
256
257
258
259
260
261
262
263
264
265
266
267
268
269
270
271
272
273
274
275
276
277
278
279
280
281
282
283
284
285
286
287
288
289
290
291
292
293
294
295
296
297
298
299
300
301
302
303
304
305
306
307
308
309
310
311
312
313
314
315
316
317
318
319
320
321
322
323
324
325
326
327
328
329
330
331
332
333
334
335
336
337
338
339
340
341
342
package main

import (
	"fmt"
	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/HewlettPackard/oneview-golang/utils"
	"os"
)

func main() {

	var (
		ClientOV       *ov.OVClient
		new_volume     = "TestVolume"
		name_to_update = "UpdatedName"
	)

	ovc := ClientOV.NewOVClient(
		os.Getenv("ONEVIEW_OV_USER"),
		os.Getenv("ONEVIEW_OV_PASSWORD"),
		os.Getenv("ONEVIEW_OV_DOMAIN"),
		os.Getenv("ONEVIEW_OV_ENDPOINT"),
		false,
		1000,
		"*")

	// Create storage volume with name <new_volume>
	properties := &ov.Properties{
		Name:                new_volume,
		Storagepool:         utils.NewNstring("/rest/storage-pools/F693B0B6-AD80-40C0-935D-AA99009ED046"),
		Size:                107374741824,
		ProvisioningType:    "Thin",
		DataProtectionLevel: "NetworkRaid10Mirror2Way",
	}

	storageVolume := ov.StorageVolume{TemplateURI: utils.NewNstring("/rest/storage-volume-templates/292c2ff4-3e9d-4936-9b24-aa99009f91a3"), Properties: properties, IsPermanent: true}

	err := ovc.CreateStorageVolume(storageVolume)
	if err != nil {
		fmt.Println("Could not create the volume", err)
	}

	// Update the given storage volume
	update_vol, _ := ovc.GetStorageVolumeByName(new_volume)

	updated_storage_volume := ov.StorageVolume{
		ProvisioningTypeForUpdate: update_vol.ProvisioningTypeForUpdate,
		IsPermanent:               update_vol.IsPermanent,
		IsShareable:               update_vol.IsShareable,
		Name:                      name_to_update,
		ProvisionedCapacity:       "107374741824",
		DeviceSpecificAttributes:  update_vol.DeviceSpecificAttributes,
		URI:             update_vol.URI,
		ETAG:            update_vol.ETAG,
		Description:     "empty",
		TemplateVersion: "1.1",
	}

	err = ovc.UpdateStorageVolume(updated_storage_volume)
	if err != nil {
		fmt.Println("Could not update the volume", err)
	}

	// Get All the volumes present
	fmt.Println("\nGetting all the storage volumes present in the system: \n")
	sort := "name:desc"
	vol_list, err := ovc.GetStorageVolumes("", sort)
	if err != nil {
		fmt.Println("Error Getting the storage volumes ", err)
	}
	for i := 0; i < len(vol_list.Members); i++ {
		fmt.Println(vol_list.Members[i].Name)
	}

	// Get volume by name
	fmt.Println("\nGetting details of volume with name: ", name_to_update)
	vol_by_name, _ := ovc.GetStorageVolumeByName(name_to_update)
	fmt.Println(vol_by_name)

	// Delete the created volume
	fmt.Println("\nDeleting the volume with name : UpdatedName")
	err = ovc.DeleteStorageVolume(name_to_update)
	if err != nil {
		fmt.Println("Delete Unsuccessful", err)
	}
}
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")
}

