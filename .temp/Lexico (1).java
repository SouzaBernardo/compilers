import java.awt.Color;
import java.awt.Component;
import java.awt.Font;
import java.io.*;

import javax.swing.JFileChooser;
import javax.swing.JOptionPane;
import javax.swing.JTextArea;
import javax.swing.filechooser.FileFilter;

/**
 * Lexico - Classe para implementa��o do L�xico para nossa gram�tica
 * 
 * @author Ricardo Ferreira de Oliveira
 *
 *
 * Gram�tica:
 * 
 * 
<G> ::= 'PROGRAMA' <LISTA> <CMDS> 'FIM'
<LISTA> ::= 'VARIAVEIS' <VARS>
<VARS> ::= <VAR> , <VARS>
<VARS> ::= <VAR> 
<CMDS> ::= <CMD> ; <CMDS>
<CMDS> ::= <CMD>
<CMD> ::= <CMD_SE>
<CMD> ::= <CMD_ENQUANTO>
<CMD> ::= <CMD_PARA>
<CMD> ::= <CMD_ATRIBUICAO>
<CMD> ::= <CMD_LER>
<CMD> ::= <CMD_ESCREVER>
<CMD_SE> ::= 'SE' <CONDICAO> <CMDS> 'FIM_SE' 
<CMD_SE> ::= 'SE' <CONDICAO> <CMDS> 'SENAO' <CMDS> 'FIM_SE' 
<CMD_ENQUANTO> ::= 'ENQUANTO' <CONDICAO> <CMDS> 'FIM_ENQUANTO'
<CMD_PARA> ::= 'PARA' <VAR> '<-' <E> 'ATE' <E> <CMDS> 'FIM_PARA' 
<CMD_ATRIBUICAO> ::= <VAR> '<-' <E>
<CMD_LER> ::= 'LER' '(' <VAR> ')' 
<CMD_ESCREVER> ::= 'ESCREVER' '(' <E> ')'
<CONDICAO> ::= <E> '>' <E> 
<CONDICAO> ::= <E> '>=' <E> 
<CONDICAO> ::= <E> '<>' <E> 
<CONDICAO> ::= <E> '<=' <E> 
<CONDICAO> ::= <E> '<' <E> 
<CONDICAO> ::= <E> '==' <E>
<E> ::= <E> + <T>
<E> ::= <E> - <T>
<E> ::= <T>
<T> ::= <T> * <F>
<T> ::= <T> / <F>
<T> ::= <T> % <F>
<T> ::= <F>
<F> ::= -<F>
<F> ::= <X> ** <F>
<F> ::= <X>
<X> ::= '(' <E> ')'
<X> ::= [0-9]+('.'[0-9]+)
<X> ::= <VAR>
 *
 *
 */

public class Lexico {

  // Lista de tokens	
  static final int T_PROGRAMA        =   1;
  static final int T_FIM             =   2;
  static final int T_VARIAVEIS       =   3;
  static final int T_VIRGULA         =   4;
  static final int T_PONTO_VIRGULA   =   5;
  static final int T_SE              =   6;
  static final int T_SENAO           =   7;
  static final int T_FIM_SE          =   8;
  static final int T_ENQUANTO        =   9;
  static final int T_FIM_ENQUANTO    =  10;
  static final int T_PARA            =  11;
  static final int T_SETA            =  12;
  static final int T_ATE             =  13;
  static final int T_FIM_PARA        =  14;
  static final int T_LER             =  15;
  static final int T_ABRE_PAR        =  16;
  static final int T_FECHA_PAR       =  17;
  static final int T_ESCREVER        =  18;
  static final int T_MAIOR           =  19;
  static final int T_MENOR           =  20;
  static final int T_MAIOR_IGUAL     =  21;
  static final int T_MENOR_IGUAL     =  22;
  static final int T_IGUAL           =  23;
  static final int T_DIFERENTE       =  24;
  static final int T_MAIS            =  25;
  static final int T_MENOS           =  26;
  static final int T_VEZES           =  27;
  static final int T_DIVIDIDO        =  28;
  static final int T_RESTO           =  29;
  static final int T_ELEVADO         =  30;
  static final int T_NUMERO          =  31;
  static final int T_ID              =  32;
  
  static final int T_FIM_FONTE       =  90;
  static final int T_ERRO_LEX        =  98;
  static final int T_NULO            =  99;

  static final int FIM_ARQUIVO       =  26;

  static File arqFonte;
  static BufferedReader rdFonte;
  static File arqDestino;

  static char   lookAhead;
  static int    token;
  static String lexema;
  static int    ponteiro;
  static String linhaFonte;
  static int    linhaAtual;
  static int    colunaAtual;
  static String mensagemDeErro;
  static StringBuffer tokensIdentificados = new StringBuffer();

  public static void main( String s[] ) throws java.io.IOException
  {
      try {
    	  abreArquivo();
    	  abreDestino();
          linhaAtual     = 0;
          colunaAtual    = 0;
          ponteiro       = 0;
          linhaFonte     = "";
          token          = T_NULO;
          mensagemDeErro = "";

          movelookAhead();

          while ( ( token != T_FIM_FONTE ) && ( token != T_ERRO_LEX ) ) {
                  buscaProximoToken();
                  mostraToken();
          }
          if ( token == T_ERRO_LEX ) {
              JOptionPane.showMessageDialog( null, mensagemDeErro, "Erro L�xico!", JOptionPane.ERROR_MESSAGE );
          } else {
              JOptionPane.showMessageDialog( null, "An�lise L�xica terminada sem erros l�xicos", "An�lise L�xica terminada!", JOptionPane.INFORMATION_MESSAGE );
          }
          exibeTokens();
          gravaSaida( arqDestino );
          fechaFonte();
      } catch( FileNotFoundException fnfe ) {
          JOptionPane.showMessageDialog( null, "Arquivo nao existe!", "FileNotFoundException!", JOptionPane.ERROR_MESSAGE );
      } catch( UnsupportedEncodingException uee ) {
          JOptionPane.showMessageDialog( null, "Erro desconhecido", "UnsupportedEncodingException!", JOptionPane.ERROR_MESSAGE );
      } catch( IOException ioe ) {
          JOptionPane.showMessageDialog( null, "Erro de io: " + ioe.getMessage(), "IOException!", JOptionPane.ERROR_MESSAGE );
      } finally {
          System.out.println( "Execucao terminada!" );
      }
  }

  static void fechaFonte() throws IOException
  {
      rdFonte.close();
  }

  static void movelookAhead() throws IOException
  {
    if ( ( ponteiro + 1 ) > linhaFonte.length() ) { // parei aqui
    	linhaAtual++;
        ponteiro = 0;
        if ( ( linhaFonte = rdFonte.readLine() ) == null ) {
            lookAhead = FIM_ARQUIVO;
        } else {
        	StringBuffer sbLinhaFonte = new StringBuffer( linhaFonte );
        	sbLinhaFonte.append( '\13' ).append( '\10' );
        	linhaFonte = sbLinhaFonte.toString();
            lookAhead = linhaFonte.charAt( ponteiro );
        }
    } else {
        lookAhead = linhaFonte.charAt( ponteiro );
    }
    if ( ( lookAhead >= 'a' ) &&
         ( lookAhead <= 'z' ) ) {
        lookAhead = (char) ( lookAhead - 'a' + 'A' );
    }
    ponteiro++;
    colunaAtual = ponteiro + 1;
  }

  static void buscaProximoToken() throws IOException
  {
	//int i, j;
        
    StringBuffer sbLexema = new StringBuffer( "" );

    // Salto espa�oes enters e tabs at� o inicio do proximo token
  	while ( ( lookAhead == 9 ) ||
	        ( lookAhead == '\n' ) ||
	        ( lookAhead == 8 ) ||
	        ( lookAhead == 11 ) ||
	        ( lookAhead == 12 ) ||
	        ( lookAhead == '\r' ) ||
	        ( lookAhead == 32 ) )
    {
        movelookAhead();
    }

    /*--------------------------------------------------------------*
     * Caso o primeiro caracter seja alfabetico, procuro capturar a *
     * sequencia de caracteres que se segue a ele e classifica-la   *
     *--------------------------------------------------------------*/
    if ( ( lookAhead >= 'A' ) && ( lookAhead <= 'Z' ) ) {   
        sbLexema.append( lookAhead );
        movelookAhead();

        while ( ( ( lookAhead >= 'A' ) && ( lookAhead <= 'Z' ) ) ||
        		( ( lookAhead >= '0' ) && ( lookAhead <= '9' ) ) || ( lookAhead == '_' ) )
        {
            sbLexema.append( lookAhead );
            movelookAhead();
        }

        lexema = sbLexema.toString();  

        /* Classifico o meu token como palavra reservada ou id */
        if ( lexema.equals( "PROGRAMA" ) )
            token = T_PROGRAMA;
        else if ( lexema.equals( "FIM" ) )
            token = T_FIM;
        else if ( lexema.equals( "VARIAVEIS" ) )
            token = T_VARIAVEIS;
        else if ( lexema.equals( "SE" ) )
            token = T_SE;
        else if ( lexema.equals( "SENAO" ) )
            token = T_SENAO;
        else if ( lexema.equals( "FIM_SE" ) )
            token = T_FIM_SE;
        else if ( lexema.equals( "ENQUANTO" ) )
            token = T_ENQUANTO;
        else if ( lexema.equals( "FIM_ENQUANTO" ) )
            token = T_FIM_ENQUANTO;
        else if ( lexema.equals( "PARA" ) )
            token = T_PARA;
        else if ( lexema.equals( "ATE" ) )
            token = T_ATE;
        else if ( lexema.equals( "FIM_PARA" ) )
            token = T_FIM_PARA;
        else if ( lexema.equals( "LER" ) )
            token = T_LER;
        else if ( lexema.equals( "ESCREVER" ) )
            token = T_ESCREVER;
        else {
        	token = T_ID;
        }
    } else if ( ( lookAhead >= '0' ) && ( lookAhead <= '9' ) ) {
        sbLexema.append( lookAhead );
        movelookAhead();
        while ( ( lookAhead >= '0' ) && ( lookAhead <= '9' ) )
        {
            sbLexema.append( lookAhead );
            movelookAhead();
        }
        token = T_NUMERO;    	
    } else if ( lookAhead == '(' ){
        sbLexema.append( lookAhead );
        token = T_ABRE_PAR;    	
        movelookAhead();
    } else if ( lookAhead == ')' ){
        sbLexema.append( lookAhead );
        token = T_FECHA_PAR;    	
        movelookAhead();
    } else if ( lookAhead == ';' ){
        sbLexema.append( lookAhead );
        token = T_PONTO_VIRGULA;    	
        movelookAhead();
    } else if ( lookAhead == ',' ){
        sbLexema.append( lookAhead );
        token = T_VIRGULA;    	
        movelookAhead();
    } else if ( lookAhead == '+' ){
        sbLexema.append( lookAhead );
        token = T_MAIS;    	
        movelookAhead();
    } else if ( lookAhead == '-' ){
        sbLexema.append( lookAhead );
        token = T_MENOS;    	
        movelookAhead();
    } else if ( lookAhead == '*' ){
        sbLexema.append( lookAhead );
        movelookAhead();
        if ( lookAhead == '*' ) {
            sbLexema.append( lookAhead );
            movelookAhead();
            token = T_ELEVADO;    	
        } else {
            token = T_VEZES;    	
        }
    } else if ( lookAhead == '/' ){
        sbLexema.append( lookAhead );
        token = T_DIVIDIDO;    	
        movelookAhead();
    } else if ( lookAhead == '%' ){
        sbLexema.append( lookAhead );
        token = T_RESTO;    	
        movelookAhead();
    } else if ( lookAhead == '<' ){
        sbLexema.append( lookAhead );
        movelookAhead();
        if ( lookAhead == '>' ) {
            sbLexema.append( lookAhead );
            movelookAhead();
            token = T_DIFERENTE;
        } else if ( lookAhead == '-' ) {  
            sbLexema.append( lookAhead );
            movelookAhead();
            token = T_SETA;
        } else if ( lookAhead == '=' ) {  
            sbLexema.append( lookAhead );
            movelookAhead();
            token = T_MENOR_IGUAL;
        } else {
            token = T_MENOR;    	
        }
    } else if ( lookAhead == '>' ){
        sbLexema.append( lookAhead );
        movelookAhead();
        if ( lookAhead == '=' ) {
            sbLexema.append( lookAhead );
            movelookAhead();
            token = T_MAIOR_IGUAL;
        } else {
            token = T_MAIOR;    	
        }        
    } else if ( lookAhead == FIM_ARQUIVO ){
         token = T_FIM_FONTE;    	
    } else {
    	token = T_ERRO_LEX;
    	mensagemDeErro = "Erro L�xico na linha: " + linhaAtual + "\nReconhecido ao atingir a coluna: " + colunaAtual + "\nLinha do Erro: <" + linhaFonte + ">\nToken desconhecido: " + lookAhead;
    	sbLexema.append( lookAhead );
    }
        
    lexema = sbLexema.toString();  
  }

  private static void abreArquivo() {

		JFileChooser fileChooser = new JFileChooser();
		
		fileChooser.setFileSelectionMode( JFileChooser.FILES_ONLY );

		FiltroJoz3 filtro = new FiltroJoz3();
	    
		fileChooser.addChoosableFileFilter( filtro );
		int result = fileChooser.showOpenDialog( null );
		
		if( result == JFileChooser.CANCEL_OPTION ) {
			return;
		}
		
		arqFonte = fileChooser.getSelectedFile();
		abreFonte( arqFonte ); 	

	}


	private static boolean abreFonte( File fileName ) {

		if( arqFonte == null || fileName.getName().trim().equals( "" ) ) {
			JOptionPane.showMessageDialog( null, "Nome de Arquivo Inv�lido", "Nome de Arquivo Inv�lido", JOptionPane.ERROR_MESSAGE );
			return false;
		} else {
			linhaAtual = 1;
	        try {
				FileReader fr = new FileReader( arqFonte );
				rdFonte = new BufferedReader( fr );
			} catch (FileNotFoundException e) {
				// TODO Auto-generated catch block
				e.printStackTrace();
			} 
			return true;
		}
	}

	private static void abreDestino() {

		JFileChooser fileChooser = new JFileChooser();
			
		fileChooser.setFileSelectionMode( JFileChooser.FILES_ONLY );

		FiltroJoz3 filtro = new FiltroJoz3();
		    
		fileChooser.addChoosableFileFilter( filtro );
		int result = fileChooser.showSaveDialog( null );
			
		if( result == JFileChooser.CANCEL_OPTION ) {
			return;
		}
			
		arqDestino = fileChooser.getSelectedFile();
	}	
	

	private static boolean gravaSaida( File fileName ) {

		if( arqDestino == null || fileName.getName().trim().equals( "" ) ) {
			JOptionPane.showMessageDialog( null, "Nome de Arquivo Inv�lido", "Nome de Arquivo Inv�lido", JOptionPane.ERROR_MESSAGE );
			return false;
		} else {
			FileWriter fw;
			try {
				System.out.println( arqDestino.toString() );
				System.out.println( tokensIdentificados.toString() );
				fw = new FileWriter( arqDestino );
				BufferedWriter bfw = new BufferedWriter( fw ); 
				bfw.write( tokensIdentificados.toString() );
				bfw.close();
				JOptionPane.showMessageDialog( null, "Arquivo Salvo: " + arqDestino, "Salvando Arquivo", JOptionPane.INFORMATION_MESSAGE );
			} catch (IOException e) {
				JOptionPane.showMessageDialog( null, e.getMessage(), "Erro de Entrada/Sa�da", JOptionPane.ERROR_MESSAGE );
			} 
			return true;
		}
	}
	
	public static void exibeTokens() {
		
		JTextArea texto = new JTextArea();
		texto.append( tokensIdentificados.toString() );
		JOptionPane.showMessageDialog(null, texto, "Tokens Identificados (token/lexema)", JOptionPane.INFORMATION_MESSAGE );
	}
	
	public static void acumulaToken( String tokenIdentificado ) {

		tokensIdentificados.append( tokenIdentificado );
		tokensIdentificados.append( "\n" );
		
	}
		
}

/**
 * Classe Interna para cria��o de filtro de sele��o
 */
class FiltroJoz3 extends FileFilter {

	public boolean accept(File arg0) {
	   	 if(arg0 != null) {
	         if(arg0.isDirectory()) {
	       	  return true;
	         }
	         if( getExtensao(arg0) != null) {
	        	 if ( getExtensao(arg0).equalsIgnoreCase( "grm" ) ) {
		        	 return true;
	        	 }
	         };
	   	 }
	     return false;
	}

	/**
	 * Retorna quais extens�es poder�o ser escolhidas
	 */
	public String getDescription() {
		return "*.grm";
	}
	
	/**
	 * Retorna a parte com a extens�o de um arquivo
	 */
	public String getExtensao(File arq) {
	if(arq != null) {
		String filename = arq.getName();
	    int i = filename.lastIndexOf('.');
	    if(i>0 && i<filename.length()-1) {
	    	return filename.substring(i+1).toLowerCase();
	    };
	}
		return null;
	}
}
