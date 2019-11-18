extern printf
section .data
acc: dd 1
fmt: db "out: %d",10,0
dbgfmt: db "ip=%d [%d, %d, %d, %d, %d, %d]",10,0
section .text
global main
main:
mov ebx,0
mov ecx,0
mov edx,0
mov esi,0
mov edi,0
add0: mov ecx, 0
call dbg
add ecx, 16
jmp add17
add1: mov ecx, 1
call dbg
mov ebx, 1
add2: mov ecx, 2
call dbg
mov edx, 1
add3: mov ecx, 3
call dbg
mov eax, ebx
imul eax, edx
mov edi, eax
add4: mov ecx, 4
call dbg
mov eax, edi
cmp eax, esi
je add7
add6: mov ecx, 6
call dbg
add ecx, 1
jmp add8
add7: mov ecx, 7
call dbg
mov eax, ebx
add eax, [acc]
mov [acc], eax
add8: mov ecx, 8
call dbg
add edx, 1
add9: mov ecx, 9
call dbg
mov eax, edx
cmp eax, esi
jg add12
add11: mov ecx, 11
call dbg
mov ecx, 2
jmp add3
add12: mov ecx, 12
call dbg
add ebx, 1
add13: mov ecx, 13
call dbg
mov eax, ebx
cmp eax, esi
jg add16
add15: mov ecx, 15
call dbg
mov ecx, 1
jmp add2
add16: mov ecx, 16
call dbg
mov eax, ecx
imul eax, ecx
mov ecx, eax
jmp finish
add17: mov ecx, 17
call dbg
add esi, 2
add18: mov ecx, 18
call dbg
mov eax, esi
imul eax, esi
mov esi, eax
add19: mov ecx, 19
call dbg
mov eax, ecx
imul eax, esi
mov esi, eax
add20: mov ecx, 20
call dbg
mov eax, esi
imul eax, 11
mov esi, eax
add21: mov ecx, 21
call dbg
add edi, 6
add22: mov ecx, 22
call dbg
mov eax, edi
imul eax, ecx
mov edi, eax
add23: mov ecx, 23
call dbg
add edi, 19
add24: mov ecx, 24
call dbg
mov eax, esi
add eax, edi
mov esi, eax
add25: mov ecx, 25
call dbg
cmp byte [acc],1
add26: mov ecx, 26
call dbg
je add29
add27: mov ecx, 27
call dbg
mov eax, ecx
add eax, [acc]
mov ecx, eax
add28: mov ecx, 28
call dbg
mov ecx, 0
jmp add1
add29: mov ecx, 29
call dbg
mov edi, ecx
add30: mov ecx, 30
call dbg
mov eax, edi
imul eax, ecx
mov edi, eax
add31: mov ecx, 31
call dbg
mov eax, ecx
add eax, edi
mov edi, eax
add32: mov ecx, 32
call dbg
mov eax, ecx
imul eax, edi
mov edi, eax
add33: mov ecx, 33
call dbg
mov eax, edi
imul eax, 14
mov edi, eax
add34: mov ecx, 34
call dbg
mov eax, edi
imul eax, ecx
mov edi, eax
add35: mov ecx, 35
call dbg
mov eax, esi
add eax, edi
mov esi, eax
add36: mov ecx, 36
call dbg
mov BYTE [acc], 0
add37: mov ecx, 37
call dbg
mov ecx, 0
jmp add1
dbg: ret
push dword edi
push dword esi
push dword edx
push dword ecx
push dword ebx
push dword [acc]
push dword edi
push dword esi
push dword edx
push dword ecx
push dword ebx
push dword [acc]
push dword ecx
push dword dbgfmt
call printf
add esp, 32
pop dword [acc]
pop dword ebx
pop dword ecx
pop dword edx
pop dword esi
pop dword edi
ret
finish: push dword [acc]
push dword fmt
call printf
call dbg
add esp, 8
retn
