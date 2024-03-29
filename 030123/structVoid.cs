using System;

struct C {
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
		C y = new C {foo=3};                    // y Objekt von Struct C mit foo = 3
		C z = y;                                // z Objekt von Struct C mit foo = 3
		M.u(ref y);
		// Aufruf der Methode u der Klasse M, mit Übergabe der Speicheradresse des Objektes y
		/*
		public  static void u (ref C x){
        	x.foo=4;                            // foo der Speicheradresse von struct y wird auf 4 gesetzt
        	x=new C{foo=5};                     // Speicheradresse wird neu initialisiert mit foo = 5
        }
        */
		Console.WriteLine("y.foo: " + y.foo);   // Ausgabe: y.foo: 5
        Console.WriteLine("z.foo: " + z.foo);   // Ausgabe: z.foo: 3
    }
}