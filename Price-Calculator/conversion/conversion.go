package conversion

import "strconv"

func StringToFloat(lines []string) ([]float64, error){
	floatLines := make([]float64, len(lines))
	for index, line := range lines {
		floatLine,err := strconv.ParseFloat(line,64)
		if err != nil {
			return nil, err
		} 
		floatLines[index] = floatLine
	}
	return floatLines,nil
}