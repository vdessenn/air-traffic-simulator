# Résumé du Manuel AIXM 4.5

## Structure des fichiers XML AIXM 4.5
Les fichiers AIXM 4.5 sont structurés en utilisant le format XML (eXtensible Markup Language). Ils contiennent des éléments et des attributs spécifiques définis par le schéma AIXM 4.5 pour représenter diverses entités aéronautiques telles que les aérodromes, les espaces aériens, les points de navigation, etc.

### Enveloppe racine
L'enveloppe racine d'un fichier AIXM 4.5 est l'élément `<AIXM-Snapshot>`. Elle définit le contexte global du document AIXM. L'éditeur du fichier, la version AIXM, la date de création et la date d'effet sont spécifiées dans les attributs de cet élément.
Elle encapsule toutes les données AIXM. Voici un exemple de l'enveloppe racine :

```xml
<AIXM-Snapshot created="2025-11-14T09:56:39.325+01:00" origin="Sia-France" version="4.5" effective="2025-12-25T00:00:00.000+01:00">
      <!-- Contenu AIXM ici -->
</AIXM-Snapshot>
```

### Liste des sections présentes dans le fichier AIXM et nombre d'occurrences
Les sections principales du fichier AIXM incluent :
   - Bordures d'Espaces Aériens (Abd): 4422
   - Géométrie dérivée d'Espaces Aériens (Adg): 585
   - Lumière d'Aérodrome au Sol (Agl): 79
   - Adresses Aérodrome/Héliport (Aha): 977
   - Obstacles Aérodrome/Héliport (Aho): 119
   - Aérodrome/Héliport (Ahp): 877
   - Services au sol Aérodrome/Héliport (Ahs): 2429
   - Utilisation de l'Aérodrome/Héliport (Ahu): 536
   - Rampe/Apron (Apn): 51
   - Espace Aérien (Ase): 5024
   - DME (Distance Measuring Equipment) (Dme): 190
   - Points Significatifs (Dpn): 4270
   - Fréquences (Fqy): 1730
   - FATO (Final Approach and Takeoff Area) (Fto): 275
   - Frontières Géographiques (Gbr): 40
   - Gate / Porte d'Embarquement(Gsd): 1588
   - Système ILS (Instrument Landing System) (Ils): 100
   - Balises d'approche (Mkr): 23
   - NDB (Non-Directional Beacon) (Ndb): 96
   - Checkpoint - Holding position ou Parking spot (Nsc): 1561
   - Obstacles (Obs): 13733
   - Organisations (Org): 31
   - Installations Passagers (Pfy): 881
   - Tableau des Niveaux de Croisière (Plb): 2
   - Colonne des Niveaux de Croisière (Plc): 4
   - Axe Piste (Rcp): 2070
   - Éclairage d'Approche Piste (Rda): 161
   - Distances Déclarées Piste (Rdd): 6969
   - Désignation Piste (Rdn): 1551
   - Éclairage Piste (Rls): 1558
   - Zone Protection Piste (Rpa): 268
   - Segments de Route (Rsg): 1494
   - Utilisation Segments Route (Rsu): 2344
   - Route (Rte): 429
   - Piste (Rwy): 778
   - Service Espace Aérien (Sae): 812
   - Service d'Aérodrome/Héliport (Sah): 808
   - Service (Ser): 891
   - Point Significatif Espace Aérien (Spa): 755
   - Date Spéciale (Spd): 46
   - Voie d'Arrêt (Swy): 65
   - TACAN (Tactical Air Navigation) (Tcn): 19
   - Piste d'Hélicoptère TLOF (Tla): 232
   - Voie de Circulation (Twy): 289
   - Adresse Contact Unité (Uas): 10
   - Unité (Uni): 400
   - VOR (VHF Omnidirectional Range) (Vor): 96

### Description des sections principales avec exemples

#### Aspect Géométrique de l'AICM 4.5
L'aspect géométrique de l'AICM 4.5 est représenté à travers divers éléments XML qui définissent les positions géographiques, les dimensions et les caractéristiques spatiales des entités aéronautiques. Les coordonnées géographiques sont généralement exprimées en latitude et longitude, tandis que les dimensions peuvent inclure des longueurs, des largeurs et des altitudes. Les éléments géométriques sont essentiels pour la navigation aérienne, la planification des vols et la gestion de l'espace aérien.

#### Aérodrome et Pistes
Les sections relatives aux aérodromes, pistes et voies de circulation comprennent :
   - **Aérodrome/Héliport (Ahp)**: Informations principales sur les aérodromes et héliports
   - **Adresses Aérodrome/Héliport (Aha)**: Adresses
   - **Obstacles Aérodrome/Héliport (Aho)**: Obstacles situés dans les aérodromes et héliports
   - **Lumière d'Aérodrome au Sol (Agl)**: Informations
   - **Services au sol Aérodrome/Héliport (Ahs)**: Services disponibles
   - **Utilisation de l'Aérodrome/Héliport (Ahu)**:
   - **Piste (Rwy)**: Informations principales sur les pistes
   - **Axe Piste (Rcp)**: Axes des pistes
   - **Éclairage d'Approche Piste (Rda)**: Éclairage des approches
   - **Distances Déclarées Piste (Rdd)**: Distances déclarées
   - **Désignation Piste (Rdn)**: Désignations des pistes
   - **Éclairage Piste (Rls)**: Éclairage des pistes
   - **Zone Protection Piste (Rpa)**: Zones de protection
   - **Piste d'Hélicoptère TLOF (Tla)**: Pistes pour hélicoptères
   - **Voie de Circulation (Twy)**: Voies de circulation
   - **Voie d'Arrêt (Swy)**: Voies d'arrêt
   - **Rampe/Apron (Apn)**: Ramps
   - **Gate / Porte d'Embarquement (Gsd)**: Portes d'embarquement
   - **Installations Passagers (Pfy)**: Installations pour passagers
   - **Service au sol Aérodrome/Héliport (Sah)**: Services au sol

##### Section relative aux Aérodromes
###### Exemple de structure d'un Aérodrome (Ahp)
```xml
<Ahp>
   <AhpUid mid="1520984">
      <codeId>LFPG</codeId>
   </AhpUid>
   <OrgUid mid="1520800">
      <txtName>FRANCE</txtName>
   </OrgUid>
   <txtName>PARIS CHARLES DE GAULLE</txtName>
   <codeIcao>LFPG</codeIcao>
   <codeIata>CDG</codeIata>
   <codeType>AD</codeType>
   <txtDescrRefPt>TWR Centrale</txtDescrRefPt>
   <geoLat>490035.09N</geoLat>
   <geoLong>0023252.15E</geoLong>
   <codeDatum>WGE</codeDatum>
   <valElev>392</valElev>
   <valGeoidUndulation>143</valGeoidUndulation>
   <uomDistVer>FT</uomDistVer>
   <txtNameCitySer>PARIS</txtNameCitySer>
   <txtDescrSite>13.5 NM de PARIS</txtDescrSite>
   <valMagVar>1.8548</valMagVar>
   <dateMagVar>2025</dateMagVar>
   <valMagVarChg>0.1344</valMagVarChg>
   <valRefT>24.7</valRefT>
   <uomRefT>C</uomRefT>
   <txtNameAdmin>DIRECTION DE L'AEROPORT DE PARIS CHARLES DE GAULLE</txtNameAdmin>
   <txtDescrAcl>Sur satellites et postes de stationnement.\\On satellites and apron stands.</txtDescrAcl>
   <valTransitionAlt>5000</valTransitionAlt>
   <uomTransitionAlt>FT</uomTransitionAlt>
   <Aht>
      <codeWorkHr>H24</codeWorkHr>
      <txtRmkWorkHr>H24</txtRmkWorkHr>
   </Aht>
</Ahp>
```
###### Exemple de structure d'Adresses Aérodrome/Héliport (Aha)
```xml
<Aha>
   <AhaUid mid="1574340">
      <AhpUid mid="1520984">
         <codeId>LFPG</codeId>
      </AhpUid>
      <codeType>TLX</codeType>
      <noSeq>4</noSeq>
   </AhaUid>
   <txtAddress>Non disponible/Not available</txtAddress>
</Aha>
```
###### Exemple de structure d'Obstacles Aérodrome/Héliport (Aho)
```xml
<Aho>
   <AhoUid mid="55729433">
      <ObsUid mid="54790906">
         <geoLat>211834S</geoLat>
         <geoLong>0552443E</geoLong>
      </ObsUid>
      <AhpUid mid="1521824">
         <codeId>FMEP</codeId>
      </AhpUid>
   </AhoUid>
</Aho>
```
###### Exemple de structure de la Lumière d'Aérodrome au Sol (Agl)
```xml
<Agl>
   <AglUid mid="1573892">
      <txtName>103</txtName>
      <codeType>BCN</codeType>
   </AglUid>
   <geoLat>415644.88N</geoLat>
   <geoLong>0084726.02E</geoLong>
   <codeDatum>U</codeDatum>
   <Agt>
      <codeWorkHr>HX</codeWorkHr>
      <txtRmkWorkHr>HN (allumage par cellules)</txtRmkWorkHr>
   </Agt>
</Agl>
```
###### Exemple de structure de Services au sol d'Aérodrome/Héliport (Ahs)
```xml
<Ahs>
   <AhsUid mid="1582700">
      <AhpUid mid="1520984">
         <codeId>LFPG</codeId>
      </AhpUid>
      <codeType>OTHER</codeType>
   </AhsUid>
   <Ast>
      <codeWorkHr>OTHER</codeWorkHr>
      <txtRmkWorkHr>NIL</txtRmkWorkHr>
   </Ast>
</Ahs>
```
###### Exemple de structure d'Utilisation Aérodrome/Héliport (Ahu)
```xml
<Ahu>
   <AhuUid mid="1522822">
      <AhpUid mid="1520984">
         <codeId>LFPG</codeId>
      </AhpUid>
   </AhuUid>
   <UsageLimitation>
      <codeUsageLimitation>OTHER</codeUsageLimitation>
      <UsageCondition>
         <AircraftClass/>
         <FlightClass>
            <codeType>OAT</codeType>
            <codeRule>IV</codeRule>
            <codeOrigin>ANY</codeOrigin>
            <codePurpose>OTHER</codePurpose>
         </FlightClass>
      </UsageCondition>
   </UsageLimitation>
   <UsageLimitation>
      <codeUsageLimitation>PERMIT</codeUsageLimitation>
      <UsageCondition>
         <AircraftClass/>
         <FlightClass>
            <codeType>GAT</codeType>
            <codeRule>IV</codeRule>
            <codeOrigin>ANY</codeOrigin>
            <codePurpose>S</codePurpose>
         </FlightClass>
      </UsageCondition>
   </UsageLimitation>
   <UsageLimitation>
      <codeUsageLimitation>PERMIT</codeUsageLimitation>
      <UsageCondition>
         <AircraftClass/>
         <FlightClass>
            <codeType>GAT</codeType>
            <codeRule>IV</codeRule>
            <codeOrigin>ANY</codeOrigin>
            <codePurpose>NS</codePurpose>
         </FlightClass>
      </UsageCondition>
   </UsageLimitation>
</Ahu>
```
###### Exemple de structure du Service au sol Aérodrome/Héliport (Sah)
```xml
<Sah>
	<SahUid mid="56300579">
		<AhpUid mid="1520972">
			<codeId>LFPO</codeId>
		</AhpUid>
		<SerUid mid="56300571">
			<UniUid mid="1524818">
				<txtName>LFPO ORLY</txtName>
			</UniUid>
			<codeType>VDF</codeType>
			<noSeq>40</noSeq>
		</SerUid>
	</SahUid>
</Sah>
```
###### Exemple de structure de Voie de Circulation (Twy)
```xml
   <Twy>
      <TwyUid mid="42231178">
         <AhpUid mid="1520856">
            <codeId>LFMU</codeId>
         </AhpUid>
         <txtDesig>A</txtDesig>
      </TwyUid>
      <codeType>OTHER</codeType>
      <valWid>23</valWid>
      <uomWid>M</uomWid>
      <codeComposition>BITUM</codeComposition>
      <valPcnClass>48</valPcnClass>
      <codePcnPavementType>F</codePcnPavementType>
      <codePcnPavementSubgrade>C</codePcnPavementSubgrade>
      <codePcnMaxTirePressure>X</codePcnMaxTirePressure>
      <codePcnEvalMethod>T</codePcnEvalMethod>
   </Twy>
```
###### Exemple de structure de Voie d'Arrêt (Swy)
```xml
   <Swy>
      <SwyUid mid="43221651">
         <RdnUid mid="1606866">
            <RwyUid mid="1606856">
               <AhpUid mid="1520912">
                  <codeId>LFSG</codeId>
               </AhpUid>
               <txtDesig>08/26</txtDesig>
            </RwyUid>
            <txtDesig>26</txtDesig>
         </RdnUid>
      </SwyUid>
      <valLen>120</valLen>
      <uomDim>M</uomDim>
      <codeSts>OTHER</codeSts>
   </Swy>
```
###### Exemple de structure de Rampe/Apron (Apn)
```xml
   <Apn>
      <ApnUid mid="16474961">
         <AhpUid mid="1521672">
            <codeId>LFMQ</codeId>
         </AhpUid>
         <txtName>LE CASTELLET APRON</txtName>
      </ApnUid>
      <codeSts>OTHER</codeSts>
      <txtMarking>NIL</txtMarking>
   </Apn>
```
###### Exemple de structure de Gate / Porte d'Embarquement (Gsd)
```xml
   <Gsd>
      <GsdUid mid="36669990">
         <ApnUid mid="36669986">
            <AhpUid mid="1520944">
               <codeId>LFRH</codeId>
            </AhpUid>
            <txtName>LFRH-APRON</txtName>
         </ApnUid>
         <txtDesig>INS</txtDesig>
      </GsdUid>
      <codeType>OTHER</codeType>
      <txtDescrRestrUse>Points de vérification INS : Seuil 20 voir AD2 LFRH12\INS checkpoints : THR 20 see AD2 LFRH12</txtDescrRestrUse>
      <geoLat>474614.32N</geoLat>
      <geoLong>0032608.20W</geoLong>
      <codeDatum>WGE</codeDatum>
      <uomDistVer>FT</uomDistVer>
      <valCrc>3AF79E22</valCrc>
   </Gsd>
```
###### Exemple de structure d'Installations Passagers (Pfy)
```xml
   <Pfy>
      <PfyUid mid="23967005">
         <AhpUid mid="1520852">
            <codeId>LFOB</codeId>
         </AhpUid>
         <codeType>OTHER</codeType>
         <noSeq>7</noSeq>
      </PfyUid>
      <txtDescr>NIL</txtDescr>
      <txtRmk>NIL</txtRmk>
   </Pfy>
```


##### Section relative aux Pistes 
###### Exemple de structure d'une Piste (Rwy)
```xml
	<Rwy>
		<RwyUid mid="1528919">
			<AhpUid mid="1521716">
				<codeId>LFVM</codeId>
			</AhpUid>
			<txtDesig>12/30</txtDesig>
		</RwyUid>
		<valLen>1000</valLen>
		<valWid>20</valWid>
		<uomDimRwy>M</uomDimRwy>
		<codeComposition>ASPH</codeComposition>
		<valPcnClass>23</valPcnClass>
		<codePcnPavementType>F</codePcnPavementType>
		<codePcnPavementSubgrade>B</codePcnPavementSubgrade>
		<codePcnMaxTirePressure>W</codePcnMaxTirePressure>
		<codePcnEvalMethod>T</codePcnEvalMethod>
		<txtPcnNote>23 F/B/W/T</txtPcnNote>
		<valLenStrip>1120</valLenStrip>
		<valWidStrip>80</valWidStrip>
		<uomDimStrip>M</uomDimStrip>
   </Rwy>
```
###### Exemple de structure d'Axe Piste (Rcp)
```xml
   <Rcp>
      <RcpUid mid="1554562">
         <RwyUid mid="1549780">
            <codeId>LFPG-08L/26R</codeId>
         </RwyUid>
         <noEnd>1</noEnd>
      </RcpUid>
      <geoLat>490036.21N</geoLat>
      <geoLong>0022423.11E</geoLong>
      <codeDatum>WGE</codeDatum>
      <valTrueBrg>084.0</valTrueBrg>
      <uomTrueBrg>DEG</uomTrueBrg>
      <valMagBrg>082.1</valMagBrg>
      <uomMagBrg>DEG</uomMagBrg>
   </Rcp>
```
###### Exemple de structure d'Éclairage d'Approche Piste (Rda)
```xml
   <Rda>
      <RdaUid mid="1561234">
         <RwyUid mid="1549780">
            <codeId>LFPG-08L/26R</codeId>
         </RwyUid>
         <noEnd>1</noEnd>
      </RdaUid>
      <codeType>ALS</codeType>
      <Rdt>
         <codeWorkHr>H24</codeWorkHr>
         <txtRmkWorkHr>H24</txtRmkWorkHr>
      </Rdt>
   </Rda>
```
###### Exemple de structure de Distances Déclarées Piste (Rdd)
```xml
   <Rdd>
      <RddUid mid="1557890">
         <RwyUid mid="1549780">
            <codeId>LFPG-08L/26R</codeId>
         </RwyUid>
         <noEnd>1</noEnd>
      </RddUid>
      <valTora>4000</valTora>
      <uomTora>M</uomTora>
      <valAsda>4000</valAsda>
      <uomAsda>M</uomAsda>
      <valLda>3800</valLda>
      <uomLda>M</uomLda>
      <valToda>4200</valToda>
      <uomToda>M</uomToda>
   </Rdd>
```
###### Exemple de structure de Désignation Piste (Rdn)
```xml
   <Rdn>
      <RdnUid mid="1551234">
         <RwyUid mid="1549780">
            <codeId>LFPG-08L/26R</codeId>
         </RwyUid>
         <noEnd>1</noEnd>
      </RdnUid>
      <codeDesignator>08L</codeDesignator>
      <valMagHeading>082.1</valMagHeading>
      <uomMagHeading>DEG</uomMagHeading>
      <valTrueHeading>084.0</valTrueHeading>
      <uomTrueHeading>DEG</uomTrueHeading>
   </Rdn>
```
###### Exemple de structure d'Éclairage Piste (Rls)
```xml
   <Rls>
      <RlsUid mid="1567890">
         <RwyUid mid="1549780">
            <codeId>LFPG-08L/26R</codeId>
         </RwyUid>
         <noEnd>1</noEnd>
      </RlsUid>
      <codeType>RCL</codeType>
      <Rlt>
         <codeWorkHr>H24</codeWorkHr>
         <txtRmkWorkHr>H24</txtRmkWorkHr>
      </Rlt>
   </Rls>
```
###### Exemple de structure de Zone Protection Piste (Rpa)  
```xml
<Rpa>
   <RpaUid mid="1573456">
      <RwyUid mid="1549780">
         <codeId>LFPG-08L/26R</codeId>
      </RwyUid>
      <noEnd>1</noEnd>
   </RpaUid>
   <valWidth>300</valWidth>
   <uomWidth>M</uomWidth>
   <valLength>4000</valLength>
   <uomLength>M</uomLength>
</Rpa>
```
###### Exemple de structure de Piste d'Hélicoptère TLOF (Tla)
```xml
   <Tla>
      <TlaUid mid="50462623">
         <AhpUid mid="1522334">
            <codeId>LF221</codeId>
         </AhpUid>
         <txtDesig>LF221-TLOF</txtDesig>
      </TlaUid>
      <FtoUid mid="39700038">
         <AhpUid mid="1522334">
            <codeId>LF221</codeId>
         </AhpUid>
         <txtDesig>LF221-FATO</txtDesig>
      </FtoUid>
      <geoLat>445450.10N</geoLat>
      <geoLong>0045416.00E</geoLong>
      <codeDatum>WGE</codeDatum>
      <valElev>535</valElev>
      <valGeoidUndulation>164</valGeoidUndulation>
      <uomDistVer>FT</uomDistVer>
      <valCrc>A0D9C0F4</valCrc>
      <valLen>20</valLen>
      <valWid>20</valWid>
      <uomDim>M</uomDim>
      <codeComposition>OTHER</codeComposition>
      <valSiwlWeight>5</valSiwlWeight>
      <uomSiwlWeight>T</uomSiwlWeight>
      <txtRmk>Aluminium</txtRmk>
   </Tla>
```
###### Exemple de structure de Checkpoint - Holding position ou Parking spot (Nsc)
```xml
   <Nsc>
      <NscUid mid="25229568">
         <GsdUid mid="25229563">
            <ApnUid mid="1574114">
               <AhpUid mid="1521708">
                  <codeId>TFFR</codeId>
               </AhpUid>
               <txtName>TFFR-APRON</txtName>
            </ApnUid>
            <txtDesig>TWY F</txtDesig>
         </GsdUid>
         <codeType>VOR</codeType>
      </NscUid>
      <txtRmk>FREQ 112.9 MHz - Distance DME / Azimuth : 0.7 NM / 076°</txtRmk>
   </Nsc>
```

#### Espace Aérien
Les sections relatives aux espaces aériens comprennent :
   - **Espace Aérien (Ase)**: Informations principales sur les espaces aériens
   - **Bordures d'Espaces Aériens (Abd)**: Définition des limites des espaces aériens
   - **Géométries Dérivées d'Espaces Aériens (Adg)**: Représentation géométrique des espaces aériens
   - **FATO (Final Approach and Takeoff Area) (Fto)**: Zones d'approche finale et de décollage
   - **Frontières Géographiques (Gbr)**: Délimitations géographiques des espaces aériens
   - **Service Espace Aérien (Sae)**: Services associés aux espaces aériens

##### Section relative aux Espaces Aériens
###### Exemple de structure d'un Espace Aérien (Ase)
```xml
	<Ase>
		<AseUid mid="1566285">
			<codeType>R</codeType>
			<codeId>LFR27</codeId>
		</AseUid>
		<txtName>27</txtName>
		<codeDistVerUpper>ALT</codeDistVerUpper>
		<valDistVerUpper>6500</valDistVerUpper>
		<uomDistVerUpper>FT</uomDistVerUpper>
		<codeDistVerLower>HEI</codeDistVerLower>
		<valDistVerLower>0</valDistVerLower>
		<uomDistVerLower>FT</uomDistVerLower>
		<Att>
			<codeWorkHr>HX</codeWorkHr>
			<Timsh>
				<codeTimeRef>UTCW</codeTimeRef>
				<dateValidWef>01-01</dateValidWef>
				<dateValidTil>31-12</dateValidTil>
				<codeDay>FRI</codeDay>
				<codeDayTil>FRI</codeDayTil>
				<timeWef>19:00</timeWef>
				<timeTil>23:00</timeTil>
			</Timsh>
			<Timsh>
				<codeTimeRef>UTCW</codeTimeRef>
				<dateValidWef>01-01</dateValidWef>
				<dateValidTil>31-12</dateValidTil>
				<codeDay>TUE</codeDay>
				<codeDayTil>TUE</codeDayTil>
				<timeWef>05:00</timeWef>
				<timeTil>11:00</timeTil>
			</Timsh>
			<Timsh>
				<codeTimeRef>UTCW</codeTimeRef>
				<dateValidWef>01-01</dateValidWef>
				<dateValidTil>31-12</dateValidTil>
				<codeDay>TUE</codeDay>
				<codeDayTil>TUE</codeDayTil>
				<timeWef>19:00</timeWef>
				<timeTil>23:00</timeTil>
			</Timsh>
			<txtRmkWorkHr>Possible activation MON, TUE, THU, FRI except public HOL#0500- 1100#1200- 1700#1900- 2300 #(SUM - 1HR).</txtRmkWorkHr>
		</Att>
		<txtRmk>IFR/VFR/OAT: avoidance mandatory during activity.</txtRmk>
	</Ase>
```
###### Exemple de structure de Bordures d'Espaces Aériens (Abd)
```xml
	<Abd>
		<AbdUid mid="46931039">
			<AseUid mid="46931037">
				<codeType>RAS</codeType>
				<codeId>LFFRASE</codeId>
			</AseUid>
		</AbdUid>
		<Avx>
			<codeType>GRC</codeType>
			<geoLat>434910N</geoLat>
			<geoLong>0074519.00E</geoLong>
			<codeDatum>WGE</codeDatum>
		</Avx>
		<Avx>
			<codeType>GRC</codeType>
			<geoLat>434644N</geoLat>
			<geoLong>0074701.00E</geoLong>
			<codeDatum>WGE</codeDatum>
		</Avx>
		<Avx>
			<codeType>GRC</codeType>
			<geoLat>434207N</geoLat>
			<geoLong>0075015.00E</geoLong>
			<codeDatum>WGE</codeDatum>
		</Avx>
	</Abd>
```
###### Exemple de structure de Géométries Dérivées d'Espaces Aériens (Adg)
```xml
	<Adg>
		<AdgUid>
			<AseUid mid="13835794">
				<codeType>SECTOR-C</codeType>
				<codeId>LFBB-BDX</codeId>
			</AseUid>
		</AdgUid>
		<AseUidBase mid="1566697">
			<codeType>SECTOR</codeType>
			<codeId>LFBBBN</codeId>
		</AseUidBase>
		<codeOpr>UNION</codeOpr>
		<AseUidComponent mid="1566683">
			<codeType>SECTOR</codeType>
			<codeId>LFBBLM</codeId>
		</AseUidComponent>
		<codeOpr>UNION</codeOpr>
		<AseUidComponent mid="1561607">
			<codeType>SECTOR</codeType>
			<codeId>LFBBTG</codeId>
		</AseUidComponent>
	</Adg>
```
###### Exemple de structure de FATO (Final Approach and Takeoff Area) (Fto)
```xml
	<Fto>
		<FtoUid mid="39700110">
			<AhpUid mid="20579145">
				<codeId>LF028</codeId>
			</AhpUid>
			<txtDesig>LF028-FATO</txtDesig>
		</FtoUid>
		<valLen>20.6</valLen>
		<valWid>20.6</valWid>
		<uomDim>M</uomDim>
		<codeComposition>OTHER</codeComposition>
		<valSiwlWeight>4</valSiwlWeight>
		<uomSiwlWeight>T</uomSiwlWeight>
		<txtRmk>20.6 x 20.6 - Résine époxy - 4 t</txtRmk>
	</Fto>
```
###### Exemple de structure de Frontières Géographiques (Gbr)
```xml
	<Gbr>
		<GbrUid mid="17263941">
			<txtName>REUNION</txtName>
		</GbrUid>
		<codeType>CS</codeType>
		<Gbv>
			<codeType>GRC</codeType>
			<geoLat>211455.91S</geoLat>
			<geoLong>0551915.01E</geoLong>
			<codeDatum>WGE</codeDatum>
		</Gbv>
		<Gbv>
			<codeType>GRC</codeType>
			<geoLat>211457.06S</geoLat>
			<geoLong>0551915.90E</geoLong>
			<codeDatum>WGE</codeDatum>
		</Gbv>
		<Gbv>
			<codeType>GRC</codeType>
			<geoLat>211457.54S</geoLat>
			<geoLong>0551915.90E</geoLong>
			<codeDatum>WGE</codeDatum>
		</Gbv>
	</Gbr>
```
###### Exemple de structure de Service Espace Aérien (Sae)
```xml
	<Sae>
		<SaeUid mid="1599736">
			<SerUid mid="1526010">
				<UniUid mid="1525116">
					<txtName>ARMOR CCM</txtName>
				</UniUid>
				<codeType>OTHER</codeType>
				<noSeq>10</noSeq>
			</SerUid>
			<AseUid mid="1560201">
				<codeType>D-OTHER</codeType>
				<codeId>ACC44</codeId>
			</AseUid>
		</SaeUid>
	</Sae>
```

#### Points significatifs
Les sections relatives aux points significatifs comprennent :
   - **Points Significatifs (Dpn)**: Informations principales sur les points significatifs
   - **Point Significatif Espace Aérien (Spa)**: Points significatifs spécifiques aux espaces aériens

##### Section relative aux Points Significatifs
###### Exemple de structure d'un Point Significatif (Dpn)
```xml
	<Dpn>
		<DpnUid mid="8397835">
			<codeId>ADN</codeId>
			<geoLat>492659N</geoLat>
			<geoLong>0024714E</geoLong>
		</DpnUid>
		<AhpUidAssoc mid="1521124">
			<codeId>LFAD</codeId>
		</AhpUidAssoc>
		<codeDatum>WGE</codeDatum>
		<valCrc>61E106E5</valCrc>
		<codeType>OTHER</codeType>
		<txtName>AD-N</txtName>
		<txtRmk>VRP-Junction between RD588 and D935</txtRmk>
	</Dpn>
```
###### Exemple de structure d'un Point Significatif Espace Aérien (Spa)
```xml
	<Spa>
		<SpaUid mid="38348725">
			<DpnUidSpn mid="1542355">
				<codeId>NASEP</codeId>
				<geoLat>443054.00N</geoLat>
				<geoLong>0021356.00E</geoLong>
			</DpnUidSpn>
			<AseUid mid="38348497">
				<codeType>RAS</codeType>
				<codeId>LFFRASW</codeId>
			</AseUid>
		</SpaUid>
		<codeType>IN</codeType>
		<txtRmk>FRA(CI): LFFRASW</txtRmk>
	</Spa>
```

#### Aides à la navigation aérienne (Navaids)
Les sections relatives aux aides à la navigation aérienne comprennent :
   - **VOR (VHF Omnidirectional Range) (Vor)**: Informations sur les VOR
   - **NDB (Non-Directional Beacon) (Ndb)**: Informations sur les NDB
   - **DME (Distance Measuring Equipment) (Dme)**: Informations sur les DME
   - **TACAN (Tactical Air Navigation) (Tcn)**: Informations sur les TACAN
   - **Système ILS (Instrument Landing System) (Ils)**: Informations sur les systèmes ILS
   - **Balises d'approche (Mkr)**: Informations sur les balises d'approche

##### Section relative aux Aides à la navigation aérienne (Navaids)
###### Exemple de structure d'un VOR (Vor)
```xml
	<Vor>
		<VorUid mid="53882756">
			<codeId>BZH</codeId>
			<geoLat>482631.4N</geoLat>
			<geoLong>0042639.1W</geoLong>
		</VorUid>
		<OrgUid mid="1520800">
			<txtName>FRANCE</txtName>
		</OrgUid>
		<txtName>BREST BRETAGNE</txtName>
		<codeType>VOR</codeType>
		<valFreq>110.65</valFreq>
		<uomFreq>MHZ</uomFreq>
		<codeTypeNorth>MAG</codeTypeNorth>
		<codeDatum>WGE</codeDatum>
		<valElev>313</valElev>
		<uomDistVer>FT</uomDistVer>
		<valCrc>42D9EDDF</valCrc>
		<Vtt>
			<codeWorkHr>H24</codeWorkHr>
			<txtRmkWorkHr>H24</txtRmkWorkHr>
		</Vtt>
	</Vor>
```
###### Exemple de structure d'un NDB (Ndb)
```xml
	<Ndb>
		<NdbUid mid="55135713">
			<codeId>OA</codeId>
			<geoLat>145721.3S</geoLat>
			<geoLong>1473937.6W</geoLong>
		</NdbUid>
		<OrgUid mid="1520776">
			<txtName>POLYNESIE FRANCAISE</txtName>
		</OrgUid>
		<txtName>RANGIROA</txtName>
		<valFreq>358</valFreq>
		<uomFreq>KHZ</uomFreq>
		<codeClass>B</codeClass>
		<codeEm>A2</codeEm>
		<codeDatum>WGE</codeDatum>
		<valElev>54</valElev>
		<uomDistVer>FT</uomDistVer>
		<valCrc>5FC258AD</valCrc>
		<Ntt>
			<codeWorkHr>H24</codeWorkHr>
			<txtRmkWorkHr>H24</txtRmkWorkHr>
		</Ntt>
	</Ndb>
```
###### Exemple de structure d'un DME (Dme)
```xml
	<Dme>
		<DmeUid mid="4423731">
			<codeId>BCY</codeId>
			<geoLat>475907.1N</geoLat>
			<geoLong>0014520.0E</geoLong>
		</DmeUid>
		<OrgUid mid="1520800">
			<txtName>FRANCE</txtName>
		</OrgUid>
		<VorUid mid="4423788">
			<codeId>BCY</codeId>
			<geoLat>475907.1N</geoLat>
			<geoLong>0014520.0E</geoLong>
		</VorUid>
		<txtName>ORLEANS BRICY</txtName>
		<codeChannel>120Y</codeChannel>
		<valGhostFreq>117.35</valGhostFreq>
		<uomGhostFreq>MHZ</uomGhostFreq>
		<codeDatum>WGE</codeDatum>
		<valElev>435</valElev>
		<uomDistVer>FT</uomDistVer>
		<valCrc>71F6F83A</valCrc>
		<Dtt>
			<codeWorkHr>H24</codeWorkHr>
		</Dtt>
	</Dme>
```
###### Exemple de structure d'un TACAN (Tcn)
```xml
	<Tcn>
		<TcnUid mid="1527240">
			<codeId>ELB</codeId>
			<geoLat>424348.00N</geoLat>
			<geoLong>0102346.00E</geoLong>
		</TcnUid>
		<OrgUid mid="1520794">
			<txtName>ITALIE</txtName>
		</OrgUid>
		<VorUid mid="1527105">
			<codeId>ELB</codeId>
			<geoLat>424348.00N</geoLat>
			<geoLong>0102346.00E</geoLong>
		</VorUid>
		<txtName>ELB</txtName>
		<codeChannel>94X</codeChannel>
		<codeDatum>U</codeDatum>
		<Ttt>
			<codeWorkHr>HO</codeWorkHr>
		</Ttt>
	</Tcn>
```
###### Exemple de structure d'un Système ILS (Ils)
```xml
	<Ils>
		<IlsUid mid="1543400">
			<RdnUid mid="1533637">
				<RwyUid mid="1528035">
					<AhpUid mid="1520972">
						<codeId>LFPO</codeId>
					</AhpUid>
					<txtDesig>07/25</txtDesig>
				</RwyUid>
				<txtDesig>25</txtDesig>
			</RdnUid>
		</IlsUid>
		<DmeUid mid="1527720">
			<codeId>OLW</codeId>
			<geoLat>484327.30N</geoLat>
			<geoLong>0022333.70E</geoLong>
		</DmeUid>
		<codeCat>III</codeCat>
		<Ilz>
			<codeId>OLW</codeId>
			<valFreq>111.75</valFreq>
			<uomFreq>MHZ</uomFreq>
			<geoLat>484306.90N</geoLat>
			<geoLong>0022114.90E</geoLong>
			<codeDatum>WGE</codeDatum>
			<valElev>274</valElev>
			<uomDistVer>FT</uomDistVer>
			<valCrc>93AFB916</valCrc>
			<Ilt>
				<codeWorkHr>H24</codeWorkHr>
			</Ilt>
		</Ilz>
		<Igp>
			<valFreq>333.35</valFreq>
			<uomFreq>MHZ</uomFreq>
			<valSlope>3</valSlope>
			<valRdh>16.5</valRdh>
			<uomRdh>M</uomRdh>
			<geoLat>484327.30N</geoLat>
			<geoLong>0022333.70E</geoLong>
			<codeDatum>WGE</codeDatum>
			<valElev>286</valElev>
			<uomDistVer>FT</uomDistVer>
			<valCrc>85B1C312</valCrc>
			<Igt>
				<codeWorkHr>H24</codeWorkHr>
				<txtRmkWorkHr>H24</txtRmkWorkHr>
			</Igt>
		</Igp>
	</Ils>
```
###### Exemple de structure d'une Balise d'approche (Mkr)
```xml
	<Mkr>
		<MkrUid mid="1543614">
			<codeId>--</codeId>
			<geoLat>415953.73N</geoLat>
			<geoLong>0092412.41E</geoLong>
		</MkrUid>
		<OrgUid mid="1520800">
			<txtName>FRANCE</txtName>
		</OrgUid>
		<IlsUid mid="1543428">
			<RdnUid mid="1536153">
				<RwyUid mid="1528089">
					<AhpUid mid="1521026">
						<codeId>LFKS</codeId>
					</AhpUid>
					<txtDesig>18/36</txtDesig>
				</RwyUid>
				<txtDesig>18</txtDesig>
			</RdnUid>
		</IlsUid>
		<codePsnIls>O</codePsnIls>
		<valFreq>75</valFreq>
		<uomFreq>MHZ</uomFreq>
		<codeDatum>WGE</codeDatum>
		<valElev>40</valElev>
		<uomDistVer>FT</uomDistVer>
		<valCrc>02CB3A7C</valCrc>
		<Mtt>
			<codeWorkHr>H24</codeWorkHr>
		</Mtt>
	</Mkr>
```

#### Routes
Les sections relatives aux routes comprennent :
   - **Route (Rte)**: Informations principales sur les routes aériennes
   - **Segments de Route (Rsg)**: Définition des segments individuels d'une route
   - **Utilisation Segments Route (Rsu)**: Conditions d'utilisation des segments de route
   - **Tableau des Niveaux de Croisière (Plb)**: Informations sur les niveaux de croisière disponibles
   - **Colonne des Niveaux de Croisière (Plc)**: Définitions des colonnes de niveaux de croisière
##### Section relative aux Routes
###### Exemple de structure d'une Route (Rte)
```xml
	<Rte>
		<RteUid mid="1544254">
			<txtDesig>T372</txtDesig>
			<txtLocDesig>LF-LF</txtLocDesig>
		</RteUid>
		<txtRmk>Control services provided by LILLE APP within its opening hours and LILLE SIV limits.</txtRmk>
	</Rte>
```
###### Exemple de structure de Segments de Route (Rsg)
```xml
	<Rsg>
		<RsgUid mid="46894775">
			<RteUid mid="1543774">
				<txtDesig>G36</txtDesig>
				<txtLocDesig>EUR</txtLocDesig>
			</RteUid>
			<DpnUidSta mid="1612823">
				<codeId>GOSAD</codeId>
				<geoLat>440636.00N</geoLat>
				<geoLong>0014534.00E</geoLong>
			</DpnUidSta>
			<DpnUidEnd mid="1536724">
				<codeId>AULON</codeId>
				<geoLat>442243.00N</geoLat>
				<geoLong>0013851.00E</geoLong>
			</DpnUidEnd>
		</RsgUid>
		<codeType>RNAV</codeType>
		<codeRnp>5</codeRnp>
		<codeLvl>L</codeLvl>
		<valDistVerUpper>195</valDistVerUpper>
		<uomDistVerUpper>FL</uomDistVerUpper>
		<codeDistVerUpper>STD</codeDistVerUpper>
		<valDistVerLower>065</valDistVerLower>
		<uomDistVerLower>FL</uomDistVerLower>
		<codeDistVerLower>STD</codeDistVerLower>
		<valWid>10</valWid>
		<uomWid>NM</uomWid>
		<codeRepAtcStart>R</codeRepAtcStart>
		<codeRepAtcEnd>R</codeRepAtcEnd>
		<codeTypePath>GRC</codeTypePath>
		<valMagTrack>342</valMagTrack>
		<valReversMagTrack>162</valReversMagTrack>
		<valLen>17</valLen>
		<uomDist>NM</uomDist>
		<txtRmk>AULON-ORBIL segment: BLW FL 145, TOULOUSE APP.</txtRmk>
	</Rsg>
```
###### Exemple de structure d'Utilisation Segments Route (Rsu)
```xml
	<Rsu>
		<RsuUid mid="1555502">
			<RsgUid mid="1548168">
				<RteUid mid="1544056">
					<txtDesig>UP860</txtDesig>
					<txtLocDesig>EUR</txtLocDesig>
				</RteUid>
				<DpnUidSta mid="1539850">
					<codeId>MOKUP</codeId>
					<geoLat>470547.00N</geoLat>
					<geoLong>0012637.00W</geoLong>
				</DpnUidSta>
				<VorUidEnd mid="1527165">
					<codeId>NTS</codeId>
					<geoLat>470939.08N</geoLat>
					<geoLong>0013646.66W</geoLong>
				</VorUidEnd>
			</RsgUid>
			<codeRteAvbl>CDR1</codeRteAvbl>
			<noSeq>20</noSeq>
			<codeDir>B</codeDir>
		</RsuUid>
		<Rul>
			<PlcUid mid="1549858">
				<PlbUid mid="1549854">
					<codeId>IR</codeId>
				</PlbUid>
				<codeId>E</codeId>
			</PlcUid>
			<valDistVerLower>195</valDistVerLower>
			<uomDistVerLower>FL</uomDistVerLower>
			<codeDistVerLower>STD</codeDistVerLower>
			<valDistVerUpper>500</valDistVerUpper>
			<uomDistVerUpper>FL</uomDistVerUpper>
			<codeDistVerUpper>STD</codeDistVerUpper>
		</Rul>
		<Rst>
			<codeWorkHr>H24</codeWorkHr>
			<txtRmkWorkHr>Only by ATC</txtRmkWorkHr>
		</Rst>
	</Rsu>
```
###### Exemple de structure de Tableau des Niveaux de Croisière (Plb)
```xml
	<Plb>
		<PlbUid mid="1549852">
			<codeId>I</codeId>
		</PlbUid>
		<codeDistVer>STD</codeDistVer>
		<uomDistVer>FL</uomDistVer>
	</Plb>
```
###### Exemple de structure de Colonne des Niveaux de Croisière (Plc)
```xml
	<Plc>
		<PlcUid mid="1549856">
			<PlbUid mid="1549852">
				<codeId>I</codeId>
			</PlbUid>
			<codeId>E</codeId>
		</PlcUid>
		<Pll>
			<valDistVer>10</valDistVer>
		</Pll>
	</Plc>
```

#### Organisations, Obstacles, Dates Spéciales
Les sections relatives aux organisations, obstacles, fréquences, services et unités comprennent :
   - **Organisations (Org)**: Informations sur les organisations aéronautiques
   - **Obstacles (Obs)**: Informations sur les obstacles aéronautiques
   - **Fréquences (Fqy)**: Informations sur les fréquences de communication
   - **Services (Ser)**: Informations sur les services aéronautiques
   - **Unités (Uni)**: Informations sur les unités aéronautiques

##### Section relative aux Organisations
###### Exemple de structure d'une Organisation (Org)
```xml
	<Org>
		<OrgUid mid="1520776">
			<txtName>POLYNESIE FRANCAISE</txtName>
		</OrgUid>
		<codeId>NT</codeId>
		<codeType>O</codeType>
	</Org>
```
###### Exemple de structure d'une Fréquence (Fqy)
```xml
	<Fqy>
		<FqyUid mid="1601596">
			<SerUid mid="1526560">
				<UniUid mid="1524972">
					<txtName>LFTH TOULON HYERES</txtName>
				</UniUid>
				<codeType>PAR</codeType>
				<noSeq>30</noSeq>
			</SerUid>
			<valFreqTrans>125.2</valFreqTrans>
		</FqyUid>
		<uomFreq>MHZ</uomFreq>
		<Ftt>
			<codeWorkHr>OTHER</codeWorkHr>
			<txtRmkWorkHr>HOR ATS</txtRmkWorkHr>
		</Ftt>
		<txtRmk>Auxiliary FREQ (RWY 05)</txtRmk>
		<Cdl>
			<txtCallSign>TOULON - PRECISION</txtCallSign>
			<codeLang>FR</codeLang>
		</Cdl>
		<Cdl>
			<txtCallSign>TOULON - PRECISION</txtCallSign>
			<codeLang>EN</codeLang>
		</Cdl>
	</Fqy>
```
###### Exemple de structure d'un Service d'Information (Ser)
```xml
	<Ser>
		<SerUid mid="15617872">
			<UniUid mid="1525192">
				<txtName>LFXA AMBERIEU</txtName>
			</UniUid>
			<codeType>TWR</codeType>
			<noSeq>20</noSeq>
		</SerUid>
		<geoLat>455847.00N</geoLat>
		<geoLong>0052016.00E</geoLong>
		<codeDatum>WGE</codeDatum>
		<valCrc>8E86B5C8</valCrc>
		<Stt>
			<codeWorkHr>OTHER</codeWorkHr>
			<txtRmkWorkHr>HOR ATS</txtRmkWorkHr>
		</Stt>
		<txtRmk>Activation announced by NOTAM.# 2 multi-frequency sets : VHF TELERAD TRST12A 1040 channels with synthesizer and VHF TELERAD TRST21A 7000 channels.</txtRmk>
	</Ser>
```
###### Exemple de structure d'une Unité (Uni)
```xml
	<Uni>
		<UniUid mid="1524610">
			<txtName>LFSN NANCY</txtName>
		</UniUid>
		<OrgUid mid="1520800">
			<txtName>FRANCE</txtName>
		</OrgUid>
		<AhpUid mid="1520954">
			<codeId>LFSN</codeId>
		</AhpUid>
		<codeType>FIC</codeType>
		<codeClass>OTHER</codeClass>
		<geoLat>484131.51N</geoLat>
		<geoLong>0061333.78E</geoLong>
		<codeDatum>WGE</codeDatum>
	</Uni>
```

##### Section relative aux Obstacles
###### Exemple de structure d'un Obstacle (Obs)
```xml
	<Obs>
		<ObsUid mid="36671255">
			<geoLat>480937.69N</geoLat>
			<geoLong>0043106.32E</geoLong>
		</ObsUid>
		<txtName>E1007-15</txtName>
		<txtDescrType>Eolienne(s)</txtDescrType>
		<codeGroup>N</codeGroup>
		<codeLgt>Y</codeLgt>
		<txtDescrLgt>jour et nuit</txtDescrLgt>
		<codeDatum>WGE</codeDatum>
		<valGeoAccuracy>150</valGeoAccuracy>
		<uomGeoAccuracy>M</uomGeoAccuracy>
		<valElev>1454</valElev>
		<valElevAccuracy>100</valElevAccuracy>
		<valHgt>493</valHgt>
		<uomDistVer>FT</uomDistVer>
		<valCrc>8E1726C2</valCrc>
	</Obs>
```

##### Section relative aux Dates Spéciales
###### Exemple de structure d'une Date Spéciale (Spd)
```xml
	<Spd>
		<SpdUid mid="29259275">
			<OrgUid mid="1520800">
				<txtName>FRANCE</txtName>
			</OrgUid>
			<codeType>HOL</codeType>
			<dateDay>21-04</dateDay>
			<dateYear>2025</dateYear>
		</SpdUid>
		<txtName>EASTER MONDAY DAY</txtName>
		<txtRmk>PUBLIC HOLIDAY</txtRmk>
	</Spd>
```

### Conclusion
**Note:** Structure extraite du fichier `AIXM4.5_all_FR_OM_2025-11-27.xml`. Les balises principales incluent Ahp (aérodromes), Ase (espaces aériens), Dpn (points), Rwy (pistes), Ser (services), etc.