//stats.go
//reads all the files in the systems directory and counts
//the total number of planets, confirmed planets, planetary systems
//and binary systems

package main

import (
   "io/ioutil"
   "fmt"
   "encoding/xml"	
   "os"
)


type System struct{
    XMLName xml.Name `xml:"system"`
    OtherConfPlanets  []string `xml:"planet>list"`
    OrphanPlanets []string `xml:"planet>discoveryyear"`
    ConfOrphanBinary []string `xml:"binary>planet>list"` 
    OtherBinPlanets []string `xml:"binary>planet>discoveryyear"`
    SingleStarPlanets []string `xml:"star>planet>list"`
    ConfirmedBinary []string `xml:"binary>star>planet>list"`
    BinaryPlanets []string `xml:"binary>star>planet>discoveryyear"`
    NestedBin []string `xml:"binary>binary>star>planet>list"`
}


func main() {

	//variables used for counting
    totalConfirmed := 0
    total := 0
    bin := 0

    //read all the xml files in the systems directory
    files, _ := ioutil.ReadDir("./")	  

    //for each file in the directory (except for stats.go, this file)
    //open the file and get its contents
    for _, f := range files {
        c := System {}
		if f.Name() == "stats.go"{ 
		}else {
		xmlFile, err := os.Open(f.Name())
        if err != nil{
	   		fmt.Println("error opening file: ", err)
	   		return
		}

		//close the file
		
     
        //read the file, if it reads EOF, exit with
        //error message
		XMLdata, _ := ioutil.ReadAll(xmlFile)
       	er := xml.Unmarshal(XMLdata, &c)
		if er != nil {
	    	fmt.Printf("error: %v", er)
	    	return
		}
		
		//count all planets within a binary system
        for key := range c.OtherBinPlanets {
            if c.OtherBinPlanets[key] != "" {
	        	total += 1
	    	}
	   	}

	   	//count all confirmed orphan plantets (no star)
	   	for key := range c.OtherConfPlanets {
	        if c.OtherConfPlanets[key] == "Confirmed planets" {
                totalConfirmed += 1
	       	}
        }

        //count all orphan planets (no star)
        for key := range c.OrphanPlanets {
            if c.OrphanPlanets[key] != "" {
                total += 1
	       	}
	   	}

	   	//count all confirmed orphan planets in a binary system 
	   	//(orbiting the system)
	   	for key := range c.ConfOrphanBinary {
	       	if c.ConfOrphanBinary[key] == "Confirmed planets" {
	           totalConfirmed += 1
	       	}
	   	}	

	   	//count all confirmed planets in a binary system 
	   	//(orbiting one star in the system)
	   	//because this is a binary, add to binary total (bin)
	   	for key := range c.ConfirmedBinary{
	       	if c.ConfirmedBinary[key] == "Confirmed planets" {
	        	totalConfirmed += 1
		   		bin += 1
	       	}
	   	}

	   	//count all planets in a binary system
	   	for key := range c.BinaryPlanets{
	      	if c.BinaryPlanets[key] != ""{
	          	total += 1
	      	}
	   	}

	   	//count all confirmed planets in a nested binary system
	   	//count the total number of planets in a nested binary system
	   	//add nested binary to binary total (bin)
	   	for key := range c.NestedBin{
            if c.NestedBin[key] == "Confirmed planets"{
                totalConfirmed += 1
		   		bin += 1
            }
	       	if c.NestedBin[key] != "" {
		   		total += 1
	       	}
        } 

        //count all confirmed planets in a single star system and
        //count the total number of planets in a single star system
	   	for key := range c.SingleStarPlanets{
            if c.SingleStarPlanets[key] == "Confirmed planets" {
				totalConfirmed += 1
            }
	       	if c.SingleStarPlanets[key] != "" {
	            total += 1       
			}
        }   
    	}	
    }

    fmt.Println("Number of confirmed planets: ", totalConfirmed)
    
    fmt.Println("Number of total planets: ", total)

    //each file is a planetary system to print out the total number of files
    //in the directory minus this program file
    fmt.Println("Number of planetary systems: ", len(files)-1)

    fmt.Println("Number of binary systems: ", bin)
    
}
