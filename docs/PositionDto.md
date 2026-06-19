# PositionDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcquisitionDateTime** | Pointer to **time.Time** |  | [optional] 
**Latitude** | Pointer to **float64** |  | [optional] 
**Longitude** | Pointer to **float64** |  | [optional] 
**Heading** | Pointer to **int32** |  | [optional] 
**Precision** | Pointer to **int32** |  | [optional] 
**Altitude** | Pointer to **int32** |  | [optional] 

## Methods

### NewPositionDto

`func NewPositionDto() *PositionDto`

NewPositionDto instantiates a new PositionDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPositionDtoWithDefaults

`func NewPositionDtoWithDefaults() *PositionDto`

NewPositionDtoWithDefaults instantiates a new PositionDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcquisitionDateTime

`func (o *PositionDto) GetAcquisitionDateTime() time.Time`

GetAcquisitionDateTime returns the AcquisitionDateTime field if non-nil, zero value otherwise.

### GetAcquisitionDateTimeOk

`func (o *PositionDto) GetAcquisitionDateTimeOk() (*time.Time, bool)`

GetAcquisitionDateTimeOk returns a tuple with the AcquisitionDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcquisitionDateTime

`func (o *PositionDto) SetAcquisitionDateTime(v time.Time)`

SetAcquisitionDateTime sets AcquisitionDateTime field to given value.

### HasAcquisitionDateTime

`func (o *PositionDto) HasAcquisitionDateTime() bool`

HasAcquisitionDateTime returns a boolean if a field has been set.

### GetLatitude

`func (o *PositionDto) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *PositionDto) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *PositionDto) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *PositionDto) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### GetLongitude

`func (o *PositionDto) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *PositionDto) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *PositionDto) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *PositionDto) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### GetHeading

`func (o *PositionDto) GetHeading() int32`

GetHeading returns the Heading field if non-nil, zero value otherwise.

### GetHeadingOk

`func (o *PositionDto) GetHeadingOk() (*int32, bool)`

GetHeadingOk returns a tuple with the Heading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeading

`func (o *PositionDto) SetHeading(v int32)`

SetHeading sets Heading field to given value.

### HasHeading

`func (o *PositionDto) HasHeading() bool`

HasHeading returns a boolean if a field has been set.

### GetPrecision

`func (o *PositionDto) GetPrecision() int32`

GetPrecision returns the Precision field if non-nil, zero value otherwise.

### GetPrecisionOk

`func (o *PositionDto) GetPrecisionOk() (*int32, bool)`

GetPrecisionOk returns a tuple with the Precision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecision

`func (o *PositionDto) SetPrecision(v int32)`

SetPrecision sets Precision field to given value.

### HasPrecision

`func (o *PositionDto) HasPrecision() bool`

HasPrecision returns a boolean if a field has been set.

### GetAltitude

`func (o *PositionDto) GetAltitude() int32`

GetAltitude returns the Altitude field if non-nil, zero value otherwise.

### GetAltitudeOk

`func (o *PositionDto) GetAltitudeOk() (*int32, bool)`

GetAltitudeOk returns a tuple with the Altitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltitude

`func (o *PositionDto) SetAltitude(v int32)`

SetAltitude sets Altitude field to given value.

### HasAltitude

`func (o *PositionDto) HasAltitude() bool`

HasAltitude returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


