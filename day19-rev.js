#ip 2
00: jmp 16
01: r1 = 1
02: r3 = 1
03: r5 = r1 * r3 // r5 === 1
04: r5 = r5 === r4 ? 1 : 0
05: r2 = r5 + r2 // if(r4 === r5) skip next
06: jmp 7
07: r0 = r1 + r0 // r0++
08: r3 = r3 + 1 // r3 = 4
09: r5 = r3 > r4 ? 1 : 0
10: jmp 10 + r5
11: jmp 2 // r2 = 2
12: r1 = r1 + 1
13: r5 = r1 > r4 ? 1 : 0
14: r2 = r5 + r2 
15: jmp 1 //r2 = 1
16: r2 = r2 * r2 // jmp 256 halt????
17: r4 = r4 + 2  // 
18: r4 = r4 * r4 //
19: r4 = r2 * r4 //
20: r4 = r4 * 11 // r4 = (((r4 + 2) ^ 2) * 19) * 11
21: r5 = r5 + 6
22: r5 = r5 * r2
23: r5 = r5 + 19
24: r4 = r4 + r5 // r4 += (((r5 + 6) * 22) + 19)
25: r2 = r2 + r0 // if(r0 === 1) j27()
26: jmp 0 //r2 = 0
27: r4 += 10550400
28: r0 = 0
29: jmp 0 //r2 = 0
