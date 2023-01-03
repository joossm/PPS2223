using System;

class C {
	public int foo;
}

class M {
public  static void u (ref C x){
	x.foo=4;
	x=new C{foo=5};
	}
}

class Program
{
    static void Main(string[] args) {
		C y = new C {foo=3};
		C z = y;
		M.u(y);
		Console.WriteLine (y.foo);
		Console.WriteLine (z.foo);
    }
}