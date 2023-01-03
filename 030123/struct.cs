using System;

struct C {
	public int foo;
}

class M {
public  static void u (C x) {
	x.foo=4;
	x=new C{foo=5};
	}
}

class Program
{
    static void Main(string[] args) {
		C y = new C {foo=3};                    // y Objekt von Struct C mit foo = 3
		C z = y;                                // z Objekt von Struct C mit foo = 3
		M.u(y);
		// Aufruf der Methode u der Klasse M, mit Übergabe des Objektes y
		/*
		public  static void u (ref C y){
        	y.foo=4;                            // foo des Objektes y wird auf 4 gesetzt
        	y=new C{foo=5};                     // Objekt y wird neu initialisiert mit foo = 5
        }
        */
		Console.WriteLine("y.foo: " + y.foo);   // Ausgabe: y.foo: 3
        Console.WriteLine("z.foo: " + z.foo);   // Ausgabe: z.foo: 3
    }
}