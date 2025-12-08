package com.geografi.geoutm;

import java.util.regex.Matcher;
import java.util.regex.Pattern;

/**
 * Purpose:
 * - MGRS/UTMREF <-> UTM <-> Lon Lat
 * 
 * Description:
 * - Package for converting coordinates between WGS84 Lon Lat, UTM and MGRS/UTMREF.
 * 
 * Releases:
 * - v0.1.0 - 2019/05/09 : initial release
 * - v0.2.0 - 2019/05/10 : coord formatting changed
 * 
 * Author:
 * - Klaus Tockloth
 * 
 * Copyright and license:
 * - Copyright (c) 2019 Klaus Tockloth
 * - MIT license
 * 
 * Remarks:
 * - This library is a partial port from "github.com/proj4js/mgrs" (JavaScript).
 * - Possible coordinate conversions:
 *   UTM -> Lon Lat
 *   UTM -> MGRS
 *   Lon Lat -> UTM
 *   Lon Lat -> MGRS
 *   MGRS -> UTM
 *   MGRS -> Lon Lat
 * 
 * Supported conversions:
 * - utm.toLL()   : converts from UTM to LL
 * - utm.toMGRS() : converts from UTM to MGRS
 * - utm.toUSNG() : converts from UTM to USNG
 * - ll.toUTM()   : converts from LL to UTM
 * - ll.toMGRS()  : converts from LL to MGRS
 * - mgrs.toUTM() : converts from MGRS to UTM
 * - mgrs.toLL()  : converts from MGRS to LL
 * - usng.toLL()  : converts from USNG to LL
 * - usng.toMGRS(): converts from USNG to MGRS
 * - usng.toUTM() : converts from USNG to UTM
 * 
 * Data objects:
 * - UTM  : ZoneNumber ZoneLetter Easting Northing
 * - LL   : Longitude Latitude
 * - MGRS : String
 * - USNG : String
 * 
 * Abbreviations:
 * - Lon    : Longitude
 * - Lat    : Latitude
 * - MGRS   : Military Grid Reference System (same as UTMREF)
 * - USNG   : United States National Grid same as MGRS formatted with spaces
 * - UTM    : Universal Transverse Mercator
 * - UTMREF : UTM Reference System (same as MGRS)
 * - WGS84  : World Geodetic System 1984 (same as EPSG:4326)
 */
public class GeoUTM {

    // Constants for column and row letters
    private static final String SET_ORIGIN_COLUMN_LETTERS = "AJSAJS";
    private static final String SET_ORIGIN_ROW_LETTERS = "AFAFAF";

    // Character constants
    private static final int CHAR_A = 65; // character 'A'
    private static final int CHAR_I = 73; // character 'I'
    private static final int CHAR_O = 79; // character 'O'
    private static final int CHAR_V = 86; // character 'V'
    private static final int CHAR_Z = 90; // character 'Z'

    /**
     * UTM coordinate class
     */
    public static class UTM {
        private int zoneNumber;
        private char zoneLetter;
        private double easting;
        private double northing;

        public UTM() {
        }

        public UTM(int zoneNumber, char zoneLetter, double easting, double northing) {
            this.zoneNumber = zoneNumber;
            this.zoneLetter = zoneLetter;
            this.easting = easting;
            this.northing = northing;
        }

        public int getZoneNumber() {
            return zoneNumber;
        }

        public void setZoneNumber(int zoneNumber) {
            this.zoneNumber = zoneNumber;
        }

        public char getZoneLetter() {
            return zoneLetter;
        }

        public void setZoneLetter(char zoneLetter) {
            this.zoneLetter = zoneLetter;
        }

        public double getEasting() {
            return easting;
        }

        public void setEasting(double easting) {
            this.easting = easting;
        }

        public double getNorthing() {
            return northing;
        }

        public void setNorthing(double northing) {
            this.northing = northing;
        }

        @Override
        public String toString() {
            return String.format("%d%c %.0f %.0f", zoneNumber, zoneLetter, easting, northing);
        }

        /**
         * Converts UTM to Lon Lat.
         */
        public LL toLL() throws CoordinateException {
            // check the ZoneNumber is valid
            if (zoneNumber < 0 || zoneNumber > 60) {
                throw new CoordinateException("Invalid zone number: " + zoneNumber);
            }

            double k0 = 0.9996;
            double a = 6378137.0; // ellip.radius
            double eccSquared = 0.00669438; // ellip.eccsq
            double e1 = (1 - Math.sqrt(1 - eccSquared)) / (1 + Math.sqrt(1 - eccSquared));

            // remove 500,000 meters offset for longitude
            double x = easting - 500000.0;
            double y = northing;

            // We must know somehow if we are in the Northern or Southern hemisphere
            if (zoneLetter < 'N') {
                y -= 10000000.0; // remove 10,000,000 meters offset used for southern hemisphere
            }

            // there are 60 zones with zone 1 being at West -180 to -174
            int longOrigin = (zoneNumber - 1) * 6 - 180 + 3; // +3 puts origin in middle of zone

            double eccPrimeSquared = eccSquared / (1 - eccSquared);

            double M = y / k0;
            double mu = M / (a * (1 - eccSquared / 4 - 3 * eccSquared * eccSquared / 64 - 5 * Math.pow(eccSquared, 3) / 256));

            double phi1Rad = mu + (3 * e1 / 2 - 27 * Math.pow(e1, 3) / 32) * Math.sin(2 * mu)
                    + (21 * e1 * e1 / 16 - 55 * Math.pow(e1, 4) / 32) * Math.sin(4 * mu)
                    + (151 * Math.pow(e1, 3) / 96) * Math.sin(6 * mu);

            double N1 = a / Math.sqrt(1 - eccSquared * Math.sin(phi1Rad) * Math.sin(phi1Rad));
            double T1 = Math.tan(phi1Rad) * Math.tan(phi1Rad);
            double C1 = eccPrimeSquared * Math.cos(phi1Rad) * Math.cos(phi1Rad);
            double R1 = a * (1 - eccSquared) / Math.pow(1 - eccSquared * Math.sin(phi1Rad) * Math.sin(phi1Rad), 1.5);
            double D = x / (N1 * k0);

            double lat = phi1Rad - (N1 * Math.tan(phi1Rad) / R1) * (D * D / 2 - (5 + 3 * T1 + 10 * C1 - 4 * C1 * C1 - 9 * eccPrimeSquared) * Math.pow(D, 4) / 24
                    + (61 + 90 * T1 + 298 * C1 + 45 * T1 * T1 - 252 * eccPrimeSquared - 3 * C1 * C1) * Math.pow(D, 6) / 720);
            lat = radToDeg(lat);

            double lon = (D - (1 + 2 * T1 + C1) * Math.pow(D, 3) / 6 + (5 - 2 * C1 + 28 * T1 - 3 * C1 * C1 + 8 * eccPrimeSquared + 24 * T1 * T1) * Math.pow(D, 5) / 120) / Math.cos(phi1Rad);
            lon = longOrigin + radToDeg(lon);

            return new LL(lat, lon);
        }

        /**
         * Converts UTM to MGRS/UTMREF.
         * 
         * @param accuracy wanted accuracy in meters. Possible values are 1, 10, 100, 1000 or 10000 meters.
         */
        public MGRS toMGRS(int accuracy) {
            // meters to number of digits
            int digits = switch (accuracy) {
                case 1 -> 5;
                case 10 -> 4;
                case 100 -> 3;
                case 1000 -> 2;
                case 10000 -> 1;
                default -> 5;
            };

            // prepend with leading zeroes
            String seasting = String.format("%05.0f", easting);
            seasting = "00000" + seasting;
            String snorthing = String.format("%05.0f", northing);
            snorthing = "00000" + snorthing;

            String mgrsStr = String.format("%d%c%s%s%s",
                    zoneNumber,
                    zoneLetter,
                    get100kID(easting, northing, zoneNumber),
                    seasting.substring(seasting.length() - 5, seasting.length() - 5 + digits),
                    snorthing.substring(snorthing.length() - 5, snorthing.length() - 5 + digits));

            return new MGRS(mgrsStr);
        }

        /**
         * Converts UTM to USNG.
         * 
         * @param accuracy wanted accuracy in meters. Possible values are 1, 10, 100, 1000 or 10000 meters.
         */
        public USNG toUSNG(int accuracy) {
            // meters to number of digits
            int digits = switch (accuracy) {
                case 1 -> 5;
                case 10 -> 4;
                case 100 -> 3;
                case 1000 -> 2;
                case 10000 -> 1;
                default -> 5;
            };

            // prepend with leading zeroes
            String seasting = String.format("%05.0f", easting);
            seasting = "00000" + seasting;
            String snorthing = String.format("%05.0f", northing);
            snorthing = "00000" + snorthing;

            String usngStr = String.format("%d%c %s %s %s",
                    zoneNumber,
                    zoneLetter,
                    get100kID(easting, northing, zoneNumber),
                    seasting.substring(seasting.length() - 5, seasting.length() - 5 + digits),
                    snorthing.substring(snorthing.length() - 5, snorthing.length() - 5 + digits));

            return new USNG(usngStr);
        }
    }

    /**
     * Lon/Lat coordinate class
     */
    public static class LL {
        private double lat;
        private double lon;

        public LL() {
        }

        public LL(double lat, double lon) {
            this.lat = lat;
            this.lon = lon;
        }

        public double getLat() {
            return lat;
        }

        public void setLat(double lat) {
            this.lat = lat;
        }

        public double getLon() {
            return lon;
        }

        public void setLon(double lon) {
            this.lon = lon;
        }

        @Override
        public String toString() {
            return String.format("%.6f %.6f", lat, lon);
        }

        /**
         * Converts Lon Lat to UTM.
         */
        public UTM toUTM() {
            double a = 6378137.0; // ellip.radius
            double eccSquared = 0.00669438; // ellip.eccsq
            double k0 = 0.9996;
            double latRad = degToRad(lat);
            double longRad = degToRad(lon);

            int zoneNumber = (int) Math.floor((lon + 180) / 6) + 1;

            // make sure the longitude 180.00 is in Zone 60
            if (lon == 180) {
                zoneNumber = 60;
            }

            // Special zone for Norway
            if (lat >= 56.0 && lat < 64.0 && lon >= 3.0 && lon < 12.0) {
                zoneNumber = 32;
            }

            // special zones for Svalbard
            if (lat >= 72.0 && lat < 84.0) {
                if (lon >= 0.0 && lon < 9.0) {
                    zoneNumber = 31;
                } else if (lon >= 9.0 && lon < 21.0) {
                    zoneNumber = 33;
                } else if (lon >= 21.0 && lon < 33.0) {
                    zoneNumber = 35;
                } else if (lon >= 33.0 && lon < 42.0) {
                    zoneNumber = 37;
                }
            }

            int longOrigin = (zoneNumber - 1) * 6 - 180 + 3; // +3 puts origin in middle of zone
            double longOriginRad = degToRad(longOrigin);

            double eccPrimeSquared = eccSquared / (1 - eccSquared);

            double N = a / Math.sqrt(1 - eccSquared * Math.sin(latRad) * Math.sin(latRad));
            double T = Math.tan(latRad) * Math.tan(latRad);
            double C = eccPrimeSquared * Math.cos(latRad) * Math.cos(latRad);
            double A = Math.cos(latRad) * (longRad - longOriginRad);

            double M = a * ((1 - eccSquared / 4 - 3 * eccSquared * eccSquared / 64 - 5 * Math.pow(eccSquared, 3) / 256) * latRad
                    - (3 * eccSquared / 8 + 3 * eccSquared * eccSquared / 32 + 45 * Math.pow(eccSquared, 3) / 1024) * Math.sin(2 * latRad)
                    + (15 * eccSquared * eccSquared / 256 + 45 * Math.pow(eccSquared, 3) / 1024) * Math.sin(4 * latRad)
                    - (35 * Math.pow(eccSquared, 3) / 3072) * Math.sin(6 * latRad));

            double utmEasting = (k0 * N * (A + (1 - T + C) * Math.pow(A, 3) / 6.0
                    + (5 - 18 * T + T * T + 72 * C - 58 * eccPrimeSquared) * Math.pow(A, 5) / 120.0) + 500000.0);

            double utmNorthing = (k0 * (M + N * Math.tan(latRad) * (A * A / 2 + (5 - T + 9 * C + 4 * C * C) * Math.pow(A, 4) / 24.0
                    + (61 - 58 * T + T * T + 600 * C - 330 * eccPrimeSquared) * Math.pow(A, 6) / 720.0)));

            if (lat < 0.0) {
                utmNorthing += 10000000.0; // 10000000 meters offset for southern hemisphere
            }

            UTM utm = new UTM();
            utm.setZoneNumber(zoneNumber);
            utm.setZoneLetter(getLetterDesignator(lat));
            utm.setEasting(Math.floor(utmEasting));
            utm.setNorthing(Math.floor(utmNorthing));

            return utm;
        }

        /**
         * Converts Lon Lat to MGRS.
         * 
         * @param accuracy wanted accuracy in meters. Possible values are 1, 10, 100, 1000 or 10000 meters.
         */
        public MGRS toMGRS(int accuracy) throws CoordinateException {
            if (lon < -180 || lon > 180) {
                throw new CoordinateException("Invalid longitude: " + lon);
            }
            if (lat < -90 || lat > 90) {
                throw new CoordinateException("Invalid latitude: " + lat);
            }
            if (lat < -80 || lat > 84) {
                throw new CoordinateException("Polar regions below 80°S and above 84°N not supported, lat = " + lat);
            }

            UTM utm = toUTM();
            return utm.toMGRS(accuracy);
        }
    }

    /**
     * MGRS coordinate class
     */
    public static class MGRS {
        private String value;

        public MGRS() {
        }

        public MGRS(String value) {
            this.value = value;
        }

        public String getValue() {
            return value;
        }

        public void setValue(String value) {
            this.value = value;
        }

        @Override
        public String toString() {
            return value;
        }

        /**
         * Converts MGRS/UTMREF to Lon Lat.
         */
        public LLWithAccuracy toLL() throws CoordinateException {
            UTMWithAccuracy utmAcc = toUTM();
            LL ll = utmAcc.utm().toLL();
            return new LLWithAccuracy(ll, utmAcc.accuracy());
        }

        /**
         * Converts MGRS/UTMREF to UTM.
         */
        public UTMWithAccuracy toUTM() throws CoordinateException {
            if (value == null || value.isEmpty()) {
                throw new CoordinateException("Invalid empty mgrs string");
            }

            String mgrsTmp = value.toUpperCase();
            StringBuilder sb = new StringBuilder();
            int i = 0;

            // get Zone number
            Pattern pattern = Pattern.compile("[A-Z]");
            while (i < mgrsTmp.length() && !pattern.matcher(String.valueOf(mgrsTmp.charAt(i))).matches()) {
                if (i >= 2) {
                    throw new CoordinateException("Bad conversion, mgrs = " + value);
                }
                sb.append(mgrsTmp.charAt(i));
                i++;
            }

            int zoneNumber;
            try {
                zoneNumber = Integer.parseInt(sb.toString());
            } catch (NumberFormatException e) {
                throw new CoordinateException("Error parsing zone number: " + sb);
            }

            // A good MGRS string has to be 4-5 digits long, ##AAA/#AAA at least.
            if (i == 0 || i + 3 > mgrsTmp.length()) {
                throw new CoordinateException("Bad conversion, mgrs = " + value);
            }

            char zoneLetter = mgrsTmp.charAt(i);
            i++;

            // Check the zone letter
            if (zoneLetter <= 'A' || zoneLetter == 'B' || zoneLetter == 'Y' || zoneLetter >= 'Z' || zoneLetter == 'I' || zoneLetter == 'O') {
                throw new CoordinateException("Zone letter " + zoneLetter + " not handled, mgrs = " + value);
            }

            String hunK = mgrsTmp.substring(i, i + 2);
            i += 2;

            int set = get100kSetForZone(zoneNumber);

            double east100k = getEastingFromChar(hunK.charAt(0), set);
            double north100k = getNorthingFromChar(hunK.charAt(1), set);

            // We have a bug where the northing may be 2000000 too low. How do we know when to roll over?
            double minNorthing = getMinNorthing(zoneLetter);

            while (north100k < minNorthing) {
                north100k += 2000000;
            }

            // calculate the char index for easting/northing separator
            int remainder = mgrsTmp.length() - i;

            if (remainder % 2 != 0) {
                throw new CoordinateException("Uneven number of digits, mgrs = " + value);
            }

            int sep = remainder / 2;

            double sepEasting = 0.0;
            double sepNorthing = 0.0;
            double accuracy = 0.0;

            if (sep > 0) {
                accuracy = 100000.0 / Math.pow(10, sep);

                String sepEastingString = mgrsTmp.substring(i, i + sep);
                try {
                    sepEasting = Double.parseDouble(sepEastingString) * accuracy;
                } catch (NumberFormatException e) {
                    throw new CoordinateException("Error parsing easting: " + sepEastingString);
                }

                String sepNorthingString = mgrsTmp.substring(i + sep);
                try {
                    sepNorthing = Double.parseDouble(sepNorthingString) * accuracy;
                } catch (NumberFormatException e) {
                    throw new CoordinateException("Error parsing northing: " + sepNorthingString);
                }
            }

            double easting = sepEasting + east100k;
            double northing = sepNorthing + north100k;

            UTM utm = new UTM();
            utm.setZoneNumber(zoneNumber);
            utm.setZoneLetter(zoneLetter);
            utm.setEasting(easting);
            utm.setNorthing(northing);

            return new UTMWithAccuracy(utm, (int) accuracy);
        }
    }

    /**
     * USNG coordinate class
     */
    public static class USNG {
        private String value;

        public USNG() {
        }

        public USNG(String value) {
            this.value = value;
        }

        public String getValue() {
            return value;
        }

        public void setValue(String value) {
            this.value = value;
        }

        @Override
        public String toString() {
            return value;
        }

        /**
         * Converts USNG to MGRS
         */
        public MGRS toMGRS() {
            return new MGRS(value.replaceAll(" ", ""));
        }

        /**
         * Converts USNG to UTM
         */
        public UTMWithAccuracy toUTM() throws CoordinateException {
            MGRS mgrs = toMGRS();
            return mgrs.toUTM();
        }

        /**
         * Converts USNG to Lon Lat.
         */
        public LLWithAccuracy toLL() throws CoordinateException {
            MGRS mgrs = toMGRS();
            UTMWithAccuracy utmAcc = mgrs.toUTM();
            LL ll = utmAcc.utm().toLL();
            return new LLWithAccuracy(ll, utmAcc.accuracy());
        }
    }

    /**
     * Helper records for returning multiple values
     */
    public record UTMWithAccuracy(UTM utm, int accuracy) {
    }

    public record LLWithAccuracy(LL ll, int accuracy) {
    }

    /**
     * Custom exception for coordinate conversion errors
     */
    public static class CoordinateException extends Exception {
        public CoordinateException(String message) {
            super(message);
        }

        public CoordinateException(String message, Throwable cause) {
            super(message, cause);
        }
    }

    /**
     * Converts from degrees to radians.
     */
    private static double degToRad(double deg) {
        return deg * (Math.PI / 180.0);
    }

    /**
     * Converts from radians to degrees.
     */
    private static double radToDeg(double rad) {
        return 180.0 * (rad / Math.PI);
    }

    /**
     * Calculates the MGRS letter designator for the given latitude.
     */
    private static char getLetterDesignator(double lat) {
        char letterDesignator = 'Z'; // Error flag

        if (84 >= lat && lat >= 72) {
            letterDesignator = 'X';
        } else if (72 > lat && lat >= 64) {
            letterDesignator = 'W';
        } else if (64 > lat && lat >= 56) {
            letterDesignator = 'V';
        } else if (56 > lat && lat >= 48) {
            letterDesignator = 'U';
        } else if (48 > lat && lat >= 40) {
            letterDesignator = 'T';
        } else if (40 > lat && lat >= 32) {
            letterDesignator = 'S';
        } else if (32 > lat && lat >= 24) {
            letterDesignator = 'R';
        } else if (24 > lat && lat >= 16) {
            letterDesignator = 'Q';
        } else if (16 > lat && lat >= 8) {
            letterDesignator = 'P';
        } else if (8 > lat && lat >= 0) {
            letterDesignator = 'N';
        } else if (0 > lat && lat >= -8) {
            letterDesignator = 'M';
        } else if (-8 > lat && lat >= -16) {
            letterDesignator = 'L';
        } else if (-16 > lat && lat >= -24) {
            letterDesignator = 'K';
        } else if (-24 > lat && lat >= -32) {
            letterDesignator = 'J';
        } else if (-32 > lat && lat >= -40) {
            letterDesignator = 'H';
        } else if (-40 > lat && lat >= -48) {
            letterDesignator = 'G';
        } else if (-48 > lat && lat >= -56) {
            letterDesignator = 'F';
        } else if (-56 > lat && lat >= -64) {
            letterDesignator = 'E';
        } else if (-64 > lat && lat >= -72) {
            letterDesignator = 'D';
        } else if (-72 > lat && lat >= -80) {
            letterDesignator = 'C';
        }

        return letterDesignator;
    }

    /**
     * Gets the two letter 100k designator for a given UTM easting, northing and zone number value.
     */
    private static String get100kID(double easting, double northing, int zoneNumber) {
        int setParm = get100kSetForZone(zoneNumber);
        int setColumn = (int) Math.floor(easting / 100000);
        int setRow = ((int) Math.floor(northing / 100000)) % 20;

        return getLetter100kID(setColumn, setRow, setParm);
    }

    /**
     * Gets the MGRS 100K set for a given UTM zone number.
     */
    private static int get100kSetForZone(int i) {
        // UTM zones are grouped, and assigned to one of a group of 6 sets.
        int numberOfSets = 6;

        int setParm = i % numberOfSets;
        if (setParm == 0) {
            setParm = numberOfSets;
        }

        return setParm;
    }

    /**
     * Gets the two-letter MGRS 100k designator given information translated from the UTM northing, easting and zone number.
     */
    private static String getLetter100kID(int column, int row, int parm) {
        // colOrigin and rowOrigin are the letters at the origin of the set
        int index = parm - 1;
        char colOrigin = SET_ORIGIN_COLUMN_LETTERS.charAt(index);
        char rowOrigin = SET_ORIGIN_ROW_LETTERS.charAt(index);

        // colInt and rowInt are the letters to build to return
        int colInt = colOrigin + column - 1;
        int rowInt = rowOrigin + row;
        boolean rollover = false;

        if (colInt > CHAR_Z) {
            colInt = colInt - CHAR_Z + CHAR_A - 1;
            rollover = true;
        }

        if (colInt == CHAR_I || (colOrigin < CHAR_I && colInt > CHAR_I) || ((colInt > CHAR_I || colOrigin < CHAR_I) && rollover)) {
            colInt++;
        }

        if (colInt == CHAR_O || (colOrigin < CHAR_O && colInt > CHAR_O) || ((colInt > CHAR_O || colOrigin < CHAR_O) && rollover)) {
            colInt++;
            if (colInt == CHAR_I) {
                colInt++;
            }
        }

        if (colInt > CHAR_Z) {
            colInt = colInt - CHAR_Z + CHAR_A - 1;
        }

        if (rowInt > CHAR_V) {
            rowInt = rowInt - CHAR_V + CHAR_A - 1;
            rollover = true;
        } else {
            rollover = false;
        }

        if ((rowInt == CHAR_I) || ((rowOrigin < CHAR_I) && (rowInt > CHAR_I)) || (((rowInt > CHAR_I) || (rowOrigin < CHAR_I)) && rollover)) {
            rowInt++;
        }

        if ((rowInt == CHAR_O) || ((rowOrigin < CHAR_O) && (rowInt > CHAR_O)) || (((rowInt > CHAR_O) || (rowOrigin < CHAR_O)) && rollover)) {
            rowInt++;
            if (rowInt == CHAR_I) {
                rowInt++;
            }
        }

        if (rowInt > CHAR_V) {
            rowInt = rowInt - CHAR_V + CHAR_A - 1;
        }

        return String.valueOf((char) colInt) + (char) rowInt;
    }

    /**
     * Gets the easting value that should be added to the other, secondary easting value.
     */
    private static double getEastingFromChar(char e, int set) throws CoordinateException {
        // colOrigin is the letter at the origin of the set for the column
        char curCol = SET_ORIGIN_COLUMN_LETTERS.charAt(set - 1);
        double eastingValue = 100000.0;
        boolean rewindMarker = false;

        while (curCol != e) {
            curCol++;
            if (curCol == CHAR_I) {
                curCol++;
            }
            if (curCol == CHAR_O) {
                curCol++;
            }
            if (curCol > CHAR_Z) {
                if (rewindMarker) {
                    throw new CoordinateException("Bad character: " + e);
                }
                curCol = (char) CHAR_A;
                rewindMarker = true;
            }
            eastingValue += 100000.0;
        }

        return eastingValue;
    }

    /**
     * Gets the northing value that should be added to the other, secondary northing value.
     */
    private static double getNorthingFromChar(char n, int set) throws CoordinateException {
        if (n > 'V') {
            throw new CoordinateException("Invalid northing, char = " + n);
        }

        // rowOrigin is the letter at the origin of the set for the column
        char curRow = SET_ORIGIN_ROW_LETTERS.charAt(set - 1);
        double northingValue = 0.0;
        boolean rewindMarker = false;

        while (curRow != n) {
            curRow++;
            if (curRow == CHAR_I) {
                curRow++;
            }
            if (curRow == CHAR_O) {
                curRow++;
            }
            // fixing a bug making whole application hang in this loop when 'n' is a wrong character
            if (curRow > CHAR_V) {
                if (rewindMarker) { // making sure that this loop ends
                    throw new CoordinateException("Bad character, char = " + n);
                }
                curRow = (char) CHAR_A;
                rewindMarker = true;
            }
            northingValue += 100000.0;
        }

        return northingValue;
    }

    /**
     * Gets the minimum northing value of a MGRS zone.
     */
    private static double getMinNorthing(char zoneLetter) throws CoordinateException {
        double northing = switch (zoneLetter) {
            case 'C' -> 1100000.0;
            case 'D' -> 2000000.0;
            case 'E' -> 2800000.0;
            case 'F' -> 3700000.0;
            case 'G' -> 4600000.0;
            case 'H' -> 5500000.0;
            case 'J' -> 6400000.0;
            case 'K' -> 7300000.0;
            case 'L' -> 8200000.0;
            case 'M' -> 9100000.0;
            case 'N' -> 0.0;
            case 'P' -> 800000.0;
            case 'Q' -> 1700000.0;
            case 'R' -> 2600000.0;
            case 'S' -> 3500000.0;
            case 'T' -> 4400000.0;
            case 'U' -> 5300000.0;
            case 'V' -> 6200000.0;
            case 'W' -> 7000000.0;
            case 'X' -> 7900000.0;
            default -> -1.0;
        };

        if (northing < 0.0) {
            throw new CoordinateException("Invalid zone letter: " + zoneLetter);
        }

        return northing;
    }
}
