**************************README**********************************

-----------------------------------------REQUIREMENTS-------------------------------------------
Go Lang version 13+

-----------------------------------------MAIN---------------------------------------------------
Find the main file to start working: hotelElectricsSystem/main/main.go
Set the path for input file with output of (cwd/pwd/dir)

# const INPUT_FILE_PATH = "C:\\Users\\Navroz\\go\\src\\hotelElectricsSystem\\input_files\\"

# to run the code
Command prompt/Terminal> go run main/main.go

-----------------------------------------INPUT--------------------------------------------------
Find input files in hotelElectricsSystem/input_files and modify input as follows

layout.txt:
1st line is number of floors
2nd line is number of main corridors per floor
3rd line is number of sub corridor per floor

e.g.:
4  // 4 floors
2  // 2 main corridors per floor
3  // 3 sub corridors per floor


events.txt:
Floor_Number Corridor_Name Corridor_Number Sensor_Input

e.g.
Floor_Number Corridor_Name Corridor_Number Sensor_Input
1				SUB				2			MOVEMENT


----------------------------------------TESTS----------------------------------------------------
Run go test ./... in the main dir i.e. hotelElectricsSystem to run tests

-------------------------------------------------------------------------------------------------




Assumptions while coding:

Algorithm
1. Main corridors ac and lights should always stay on.
2. Whenever movement happens switch on lights in sub-corridor
3. Whenever sensor calls out that sub-corridor doesn't have movement since a minute siwtch light off and ac on.
4. Given the assumption as per (1) it says that main corridors lights and acs can never be switched off.
5. All manipulations to be done on sub-corridor lights and ac.
6. In case (3) if ac is switched on then another ac might be required to be switched off. The priority for choosing the ac to switch off will be 1. one which was not switched on recently and lights are off too. 2. if 1 exhausts then one where light is on we'll switch off ac. 3. ac that was recently switched on. (case 3 should never arise as the constraint always covers having the ac or light switched on in one place)
7. No cross floor switching is necessary


Input:
Takes number of floor, main and sub corridors per floor.
Floor corridor_type corridor_number sensor_input

Assumption with events is that these can come in adhoc at any point in time.

Design:
Can later add more different appliances
Can remove and replace components: controller, layout, floor, corridor, appliances, algorithms
Can feed input via text files 
Can run tests to have a basics test of things


-------------------------------------------------------------------Updates Version 0.2---------------------------------------------------
* Removes usage of single letter pointer receivers and fixes single letter input at one place with th processor.
* Reduces usage of if -else ladders within algorithms(previously processors) and also tried to shorten up functions a bit.
* Breaks generic methods fromm controller like print floor etc and puts into helpers.
* Breaks processor into algorithm and detaches from controller.
* Algorithms are a interfaced on Controller. To add a new algorithm just satisfy an interface and get going.
* Adding build system: I did do a go.mod init but it didn't make sense for me to populate the file with reference to local all the time.
* Not adding build system as it didn't make sense to me until code is published