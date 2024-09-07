format ELF64 executable 3

dump:
    mov     r8, -3689348814741910323
    sub     rsp, 40
    mov     BYTE [rsp+31], 10
    lea     rcx, [rsp+30]
.L2:
    mov     rax, rdi
    mul     r8
    mov     rax, rdi
    shr     rdx, 3
    lea     rsi, [rdx+rdx*4]
    add     rsi, rsi
    sub     rax, rsi
    mov     rsi, rcx
    sub     rcx, 1
    add     eax, 48
    mov     BYTE [rcx+1], al
    mov     rax, rdi
    mov     rdi, rdx
    cmp     rax, 9
    ja      .L2
    lea     rdx, [rsp+32]
    mov     edi, 1
    sub     rdx, rsi
    mov     rax, 1
    syscall
    add     rsp, 40
    ret
entry _start
_start:
    push 34
    push 35
    pop rax
    pop rbx
    add rax, rbx
    push rax
    pop rdi
    call dump
    push 500
    push 80
    pop rax
    pop rbx
    sub rbx, rax
    push rbx
    pop rdi
    call dump
    mov rax, 60
    mov rdi, 0
    syscall
