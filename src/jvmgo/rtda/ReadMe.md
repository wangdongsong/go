运行时数据区——实现线程私有的运行时数据区

Run-time data area

运行时数据包括：线程私有和堆内存

<pre>
线程私有：由PC寄存器和Java虚拟机栈组成
    PC寄存器：保存当前正在执行的Java虚拟机指令地址
    JVM栈：由栈帧组成，而栈帧保存方法执行的状态，包括局部变量表和操作数栈
        局部变量表
        操作数桢
        
共享内存（堆和方法区，都属于堆）：保存类数据和对象（类实例）
    对象：对象数据存入堆中。
    方法区：保存类数据
</pre>